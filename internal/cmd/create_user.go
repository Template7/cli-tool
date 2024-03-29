package cmd

import (
	"cli-tool/internal/backend"
	"context"
	"fmt"
	"github.com/Template7/common/logger"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var CreateUser = cli.Command{
	Name:    "Create user",
	Usage:   "Create users for test",
	Aliases: []string{"cu"},
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
		&cli.IntFlag{
			Name:    "amount",
			Aliases: []string{"a"},
			Usage:   "Amount of users",
			Value:   1,
		},
	},
	Action: func(c *cli.Context) error {
		ctx := context.WithValue(c.Context, "traceId", uuid.NewString())
		runCreateUser(ctx, c.String("adminUsername"), c.String("adminPassword"), c.Int("amount"))
		return nil
	},
	Subcommands: cli.Commands{
		{
			Name:    "Specify username",
			Usage:   "Specify username",
			Aliases: []string{"su"},
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
				log := logger.New().WithContext(ctx).With("username", c.String("username"))
				log.Info("create user with specify username and password")

				adminToken := backend.New().NativeLogin(ctx, c.String("adminUsername"), c.String("adminPassword"))
				if adminToken == "" {
					log.Error("admin login fail")
					return nil
				}

				if err := backend.New().CreateUser(ctx, c.String("username"), c.String("password"), "user", gofakeit.FirstName(), gofakeit.Email(), adminToken); err != nil {
					log.WithError(err).Error("fail to create user")
				}
				return nil
			},
		},
	},
}

func runCreateUser(ctx context.Context, adminUsername string, adminPassword string, amount int) {
	log := logger.New().WithContext(ctx).With("amount", amount)
	log.Info("run create user")

	adminToken := backend.New().NativeLogin(ctx, adminUsername, adminPassword)
	if adminToken == "" {
		log.Error("admin login fail")
		return
	}

	for i := 0; i < amount; i++ {
		username := fmt.Sprintf("fakeUser%03d", i+1)
		if err := backend.New().CreateUser(ctx, username, username, "user", gofakeit.FirstName(), gofakeit.Email(), adminToken); err != nil {
			log.WithError(err).Error("fail to create user")
		}
	}
}
