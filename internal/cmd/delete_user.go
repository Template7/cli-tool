package cmd

import (
	"cli-tool/internal/backend"
	"cli-tool/internal/db"
	"context"
	"github.com/Template7/common/logger"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var DeleteUser = cli.Command{
	Name:    "Delete user",
	Usage:   "Delete test users",
	Aliases: []string{"du"},
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
		runDeleteUser(ctx, c.String("adminUsername"), c.String("adminPassword"))
		return nil
	},
}

func runDeleteUser(ctx context.Context, adminUsername string, adminPassword string) {
	log := logger.New().WithContext(ctx)
	log.Info("run delete user")

	adminToken := backend.New().NativeLogin(ctx, adminUsername, adminPassword)
	if adminToken == "" {
		log.Error("admin login fail")
		return
	}

	for _, userId := range db.New().ListFakeUsers(ctx) {
		if err := backend.New().DeleteUser(ctx, userId, adminToken); err != nil {
			log.WithError(err).Error("fail to delete fake user")
		}
	}
}
