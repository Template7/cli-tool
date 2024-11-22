package cmd

import (
	"context"
	"github.com/Template7/common/config"
	"github.com/Template7/common/db"
	"github.com/Template7/common/logger"
	"github.com/Template7/common/models"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
)

var MigrateDb = cli.Command{
	Name:        "MigrateDB",
	Usage:       "Migrate DB, init table if not exist.",
	Description: "DB info config at config.yaml",
	Aliases:     []string{"md"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"s"},
			Usage:   "DB host",
			Value:   config.New().Db.Sql.Host,
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "DB host port",
			Value:   config.New().Db.Sql.Port,
		},
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "DB username",
			Value:   config.New().Db.Sql.Username,
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"pw"},
			Usage:   "DB password",
			Value:   config.New().Db.Sql.Password,
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

		if c.String("host") != "" {
			config.New().Db.Sql.Host = c.String("host")
		}
		if c.Int("port") != 0 {
			config.New().Db.Sql.Port = c.Int("port")
		}
		if c.String("username") != "" {
			config.New().Db.Sql.Username = c.String("username")
		}
		if c.String("password") != "" {
			config.New().Db.Sql.Password = c.String("password")
		}

		runDbMigration(ctx)

		return nil
	},
}

func runDbMigration(ctx context.Context) {
	log := logger.New().WithContext(ctx)
	log.Info("run db migration")

	log.With("db_info", config.New().Db.Sql).Info("db info")

	if err := db.NewSql().Debug().AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Balance{},
		&models.DepositHistory{},
		&models.WithdrawHistory{},
		&models.TransferHistory{}); err != nil {
		log.WithError(err).Error("db migration fail")
		return
	}
	log.Info("finish db migration")
}
