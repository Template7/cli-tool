package cmd

import (
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
			Name:     "username",
			Aliases:  []string{"u"},
			Usage:    "Username",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Aliases:  []string{"p"},
			Usage:    "Password",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		ctx := context.WithValue(c.Context, "traceId", uuid.NewString())
		runDemo(ctx, c.String("username"), c.String("password"))
		return nil
	},
}

func runDemo(ctx context.Context, username string, password string) {
	log := logger.New().WithContext(ctx)
	log.Info("run demo")

	u1 := user.New(ctx, username, password)
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

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_usd)], 1); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after deposit: ", u1.String())
	fmt.Println("receiver info after deposit: ", u2.String())

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_usd)], 23); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after deposit: ", u1.String())
	fmt.Println("receiver info after deposit: ", u2.String())

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_usd)], 456); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after deposit: ", u1.String())
	fmt.Println("receiver info after deposit: ", u2.String())

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_ntd)], 6); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after deposit: ", u1.String())
	fmt.Println("receiver info after deposit: ", u2.String())

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_ntd)], 54); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after deposit: ", u1.String())
	fmt.Println("receiver info after deposit: ", u2.String())

	if err := u1.Deposit(ctx, v1.Currency_name[int32(v1.Currency_ntd)], 321); err != nil {
		log.WithError(err).Error("fail to deposit")
		return
	}
	u1.GetWallet(ctx)
	u2w := u2.GetWallet(ctx)
	fmt.Println("sender info after deposit: ", u1.String())
	fmt.Println("receiver info after deposit: ", u2.String())

	for wId, _ := range u2w {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_usd)], 6); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after transfer: ", u1.String())
	fmt.Println("receiver info after transfer: ", u2.String())

	for wId, _ := range u2w {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_usd)], 78); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	u1w := u1.GetWallet(ctx)
	u2w = u2.GetWallet(ctx)
	fmt.Println("sender info after transfer: ", u1.String())
	fmt.Println("receiver info after transfer: ", u2.String())

	for wId, _ := range u2w {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_ntd)], 91); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	u1.GetWallet(ctx)
	u2.GetWallet(ctx)
	fmt.Println("sender info after transfer: ", u1.String())
	fmt.Println("receiver info after transfer: ", u2.String())

	for wId, _ := range u2w {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_ntd)], 234); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	u1w = u1.GetWallet(ctx)
	u2w = u2.GetWallet(ctx)
	fmt.Println("sender info after transfer: ", u1.String())
	fmt.Println("receiver info after transfer: ", u2.String())

	for wId, _ := range u2w {
		if err := u1.Transfer(ctx, wId, v1.Currency_name[int32(v1.Currency_ntd)], 5); err != nil {
			log.WithError(err).Error("fail to deposit")
		}
		break
	}
	u1w = u1.GetWallet(ctx)
	u2w = u2.GetWallet(ctx)
	fmt.Println("sender info after transfer: ", u1.String())
	fmt.Println("receiver info after transfer: ", u2.String())

	// withdraw all the currencies
	for wId, bls := range u1w {
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
	for wId, bls := range u2w {
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
	for wId, bls := range u1w {
		for cur, _ := range bls {
			u1.GetBalanceRecord(ctx, wId, cur)
		}
	}
	for wId, bls := range u2w {
		for cur, _ := range bls {
			u2.GetBalanceRecord(ctx, wId, cur)
		}
	}
	fmt.Println("show user info: ", u1.String())
	fmt.Println("show user info: ", u2.String())
}
