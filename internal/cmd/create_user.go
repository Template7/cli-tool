package cmd

import (
	"context"
	"github.com/Template7/common/logger"
	"github.com/urfave/cli/v2"
)

var CreateUser = cli.Command{
	Name:    "Create user",
	Usage:   "Create users for test",
	Aliases: []string{"cu"},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "amount",
			Aliases: []string{"a"},
			Usage:   "Amount of users",
			Value:   1,
		},
	},
	Action: func(c *cli.Context) error {
		runCreateUser(c.Context, c.Int("amount"))
		return nil
	},
}

func runCreateUser(ctx context.Context, amount int) {
	log := logger.New().WithContext(ctx).With("amount", amount)
	log.Debug("run create user")

	// TODO: implementation
}
