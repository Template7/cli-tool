package cli

import (
	"cli-tool/internal/pkg/config"
	"cli-tool/internal/pkg/db"
	"github.com/urfave/cli/v2"
)

var (
	initDb = &cli.Command{
		Name:    "InitDB",
		Usage:   "Initialize DB",
		Aliases: []string{"id"},
		Flags:   initDbFlag,
		Action: func(c *cli.Context) error {
			config.SetDbConnectionString()
			db.New().InitDb(c.Bool("force"))
			return nil
		},
	}

	initDbFlag = []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Force init (drop DB first).",
			Value:   false,
		},
		&cli.StringFlag{
			Name:        "mongo_db",
			Aliases:     []string{"md"},
			Usage:       "Mongo DB name",
			Value:       config.New().Mongo.Db,
			Destination: &config.New().Mongo.Db,
		},
		&cli.StringFlag{
			Name:        "mongo_host",
			Aliases:     []string{"mh"},
			Usage:       "Mongo DB host address",
			Value:       config.New().Mongo.Host,
			Destination: &config.New().Mongo.Host,
		},
		&cli.IntFlag{
			Name:        "mongo_port",
			Aliases:     []string{"mp"},
			Usage:       "Mongo DB port",
			Value:       config.New().Mongo.Port,
			Destination: &config.New().Mongo.Port,
		},
		&cli.StringFlag{
			Name:        "mongo_user",
			Aliases:     []string{"mu"},
			Usage:       "Mongo DB user name",
			Value:       config.New().Mongo.Username,
			Destination: &config.New().Mongo.Username,
		},
		&cli.StringFlag{
			Name:        "mongo_password",
			Aliases:     []string{"mpw"},
			Usage:       "Mongo DB user password",
			Value:       config.New().Mongo.Password,
			Destination: &config.New().Mongo.Password,
		},
		&cli.StringFlag{
			Name:        "sql_db",
			Aliases:     []string{"sd"},
			Usage:       "SQL DB name",
			Value:       config.New().MySql.Db,
			Destination: &config.New().MySql.Db,
		},
		&cli.StringFlag{
			Name:        "sql_host",
			Aliases:     []string{"sh"},
			Usage:       "SQL DB host address",
			Value:       config.New().MySql.Host,
			Destination: &config.New().MySql.Host,
		},
		&cli.IntFlag{
			Name:        "sql_port",
			Aliases:     []string{"sp"},
			Usage:       "SQL DB port",
			Value:       config.New().MySql.Port,
			Destination: &config.New().MySql.Port,
		},
		&cli.StringFlag{
			Name:        "sql_user",
			Aliases:     []string{"su"},
			Usage:       "SQL DB user name",
			Value:       config.New().MySql.Username,
			Destination: &config.New().MySql.Username,
		},
		&cli.StringFlag{
			Name:        "sql_password",
			Aliases:     []string{"spw"},
			Usage:       "SQL DB user password",
			Value:       config.New().MySql.Password,
			Destination: &config.New().MySql.Password,
		},
		&cli.StringFlag{
			Name:        "sql_root",
			Aliases:     []string{"sr"},
			Usage:       "SQL DB root",
			Value:       config.New().MySql.Root,
			Destination: &config.New().MySql.Root,
		},
		&cli.StringFlag{
			Name:        "sql_root_password",
			Aliases:     []string{"srp"},
			Usage:       "SQL DB root password",
			Value:       config.New().MySql.RootPassword,
			Destination: &config.New().MySql.RootPassword,
		},
	}
)
