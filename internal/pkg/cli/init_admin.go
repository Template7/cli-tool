package cli

import (
	"cli-tool/internal/pkg/db"
	"cli-tool/internal/pkg/db/collection"
	"cli-tool/internal/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	initAdmin = &cli.Command{
		Name:    "InitAdmin",
		Usage:   "Initialize admin data in DB",
		Aliases: []string{"ia"},
		Flags:   initAdminFlag,
		Action: func(c *cli.Context) error {
			username := c.String("username")
			password := c.String("password")
			createAdmin(username, password)
			return nil
		},
	}

	initAdminFlag = []cli.Flag{
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "Admin username",
			Value:   "admin",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "Admin password",
			Value:   "password",
		},
	}
)

func createAdmin(username string, password string) {
	log.Debug("create admin")

	hashedPassword, err := util.HashedPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	admin := collection.Admin{
		Username: username,
		Password: hashedPassword,
	}
	if err := db.New().CreateAdmin(admin); err != nil {
		log.Fatal(err)
	}

	log.Info("admin crated")
	return
}
