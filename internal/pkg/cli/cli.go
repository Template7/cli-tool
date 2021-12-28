package cli

import (
	"github.com/urfave/cli/v2"
)

const version = "0.1.0"

func InitCli() (app *cli.App) {

	app = &cli.App{
		Name:    "Template7 CLI Tool",
		Version: version,
	}

	app.Commands = initCommand()
	return
}

func initCommand() []*cli.Command {
	return []*cli.Command{
		initDb,
		initAdmin,
		fakeData,
	}
}
