package cmd

import (
	"cli-tool/internal/user"
	"context"
	"fmt"
	"github.com/Template7/common/logger"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var Simulation = cli.Command{
	Name:  "Simulation",
	Usage: "Simulate random user transaction process",
	Description: "For each user behavior: \n\t - If the user have no money in his wallet, " +
		"it will deposit random [1, 1000) with random currency(one of usd, ntd, jpy, cny). \n\t" +
		" - Once the user have more than 700 of any currency in its wallet, it will withdraw random [1, 700] from the wallet of the currency. \n\t" +
		" - If the user have less or equal than 700 of any currency, it will transfer random [1, 700] to any one of other user of the currency.",
	Aliases: []string{"sm"},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "amount",
			Aliases: []string{"a"},
			Usage:   "Amount of users",
			Value:   2,
		},
		&cli.IntFlag{
			Name:    "rest",
			Aliases: []string{"r"},
			Usage:   "Rest duration for each user after transaction in millisecond",
			Value:   100,
		},
		&cli.IntFlag{
			Name:    "threshold",
			Aliases: []string{"t"},
			Usage:   "Threshold to withdraw",
			Value:   700,
		},
		&cli.IntFlag{
			Name:    "duration",
			Aliases: []string{"d"},
			Usage:   "Simulation duration in second",
			Value:   10,
		},
	},

	Action: func(c *cli.Context) error {
		ctx, cancel := context.WithCancel(context.WithValue(c.Context, "traceId", uuid.NewString()))

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-quit
			cancel()
		}()

		runSimulation(ctx, c.Int("amount"), c.Int("rest"), c.Int("threshold"), c.Int("duration"))

		return nil
	},
}

func runSimulation(ctx context.Context, amount int, rest int, threshold int, duration int) {
	log := logger.New().WithContext(ctx).With("amount", amount)
	log.Info("run simulation")

	wallets := make([]string, amount)
	users := make([]*user.User, amount)
	for i := 0; i < amount; i++ {
		username := fmt.Sprintf("fakeUser%03d", i+1)
		u := user.New(ctx, username, username)
		if u == nil {
			log.With("username", username).Error("fail to new user")
			return
		}
		wl := u.GetWallet(ctx)
		for wId, _ := range wl {
			wallets[i] = wId
			break
		}
		users[i] = u
	}

	var wg sync.WaitGroup
	for i, u := range users {
		wg.Add(1)
		go func(u *user.User, i int) {
			defer wg.Done()
			u.Do(ctx, threshold, append(wallets[:i], wallets[i+1:]...), rest)
		}(u, i)
	}

	wg.Wait()
	log.Info("finish simulation")
}
