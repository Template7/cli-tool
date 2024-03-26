package cmd

import (
	"cli-tool/internal/backend"
	"cli-tool/internal/user"
	"context"
	"fmt"
	"github.com/Template7/common/logger"
	v1 "github.com/Template7/protobuf/gen/proto/template7/wallet"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var Demo = cli.Command{
	Name:    "Demo",
	Usage:   "Perform the scenario with all the functions such as deposit, transfer ans withdraw",
	Aliases: []string{"dm"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "adminUsername",
			Aliases:  []string{"au"},
			Usage:    "Admin username",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "adminPassword",
			Aliases:  []string{"ap"},
			Usage:    "Admin password",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		ctx := context.WithValue(c.Context, "traceId", uuid.NewString())
		runDemo(ctx, c.String("adminUsername"), c.String("adminPassword"))
		return nil
	},
}

func runDemo(ctx context.Context, adminUsername string, adminPassword string) {
	log := logger.New().WithContext(ctx)
	log.Info("run demo")

	setupDemo(ctx, adminUsername, adminPassword)

	u1 := user.New(ctx, "demoSender", "demoSender")
	if u1 == nil {
		log.Error("fail to new user")
		return
	}

	u2 := user.New(ctx, "demoReceiver", "demoReceiver")
	if u2 == nil {
		log.Error("fail to new user")
		return
	}

	fmt.Println("sender info: ", u1.String())
	fmt.Println("receiver info: ", u2.String())

	show := func() {
		u1.GetWallet(ctx)
		u2.GetWallet(ctx)
		fmt.Println("sender info: ", u1.String())
		fmt.Println("receiver info: ", u2.String())
	}

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_usd)], 1); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	show()

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_usd)], 23); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	show()

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_usd)], 456); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	show()

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_ntd)], 6); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	show()

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_ntd)], 54); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	show()

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_ntd)], 321); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	show()

	for wId, _ := range u2.GetWallet(ctx) {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_usd)], 6); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	show()

	for wId, _ := range u2.GetWallet(ctx) {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_usd)], 78); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	show()

	for wId, _ := range u2.GetWallet(ctx) {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_ntd)], 91); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	show()

	for wId, _ := range u2.GetWallet(ctx) {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_ntd)], 234); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	show()

	for wId, _ := range u2.GetWallet(ctx) {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_ntd)], 5); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	show()

	// withdraw all the currencies
	for wId, bls := range u1.GetWallet(ctx) {
		for cur, am := range bls {
			if am == 0 {
				continue
			}
			log.With("walletId", wId).With("currency", cur).With("amount", am).Debug("do withdraw")
			if err := u1.Withdraw(ctx, cur, am); err != nil {
				log.WithError(err).Error("fail to withdraw")
				return
			}
			u1.GetWallet(ctx)
			u2.GetWallet(ctx)
			fmt.Println("sender info after withdraw: ", u1.String())
			fmt.Println("receiver info after withdraw: ", u2.String())
		}
	}
	for wId, bls := range u2.GetWallet(ctx) {
		for cur, am := range bls {
			if am == 0 {
				continue
			}
			log.With("walletId", wId).With("currency", cur).With("amount", am).Debug("do withdraw")
			if err := u2.Withdraw(ctx, cur, am); err != nil {
				log.WithError(err).Error("fail to withdraw")
				return
			}
			u1.GetWallet(ctx)
			u2.GetWallet(ctx)
			fmt.Println("sender info after withdraw: ", u1.String())
			fmt.Println("receiver info after withdraw: ", u2.String())
		}
	}

	// show record history
	for wId, bls := range u1.GetWallet(ctx) {
		for cur, _ := range bls {
			u1.GetBalanceRecord(ctx, wId, cur)
		}
	}
	for wId, bls := range u2.GetWallet(ctx) {
		for cur := range bls {
			u2.GetBalanceRecord(ctx, wId, cur)
		}
	}
	fmt.Println("show user info: ", u1.String())
	fmt.Println("show user info: ", u2.String())
}

func setupDemo(ctx context.Context, adminUsername string, adminPassword string) {
	log := logger.New().WithContext(ctx)
	log.Info("setup demo")

	adminToken := backend.New().NativeLogin(ctx, adminUsername, adminPassword)
	if adminToken == "" {
		log.Error("admin login fail")
		return
	}

	if err := backend.New().CreateUser(ctx, "demoSender", "demoSender", "user", "demoSenderFirstName", "demoSender@email.com", adminToken); err != nil {
		log.WithError(err).Panic("fail to create user")
	}
	if err := backend.New().CreateUser(ctx, "demoReceiver", "demoReceiver", "user", "demoReceiverFirstName", "demoReceiver@email.com", adminToken); err != nil {
		log.WithError(err).Panic("fail to create user")
	}

	log.Info("setup demo done")
}
