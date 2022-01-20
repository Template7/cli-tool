package cli

import (
	"cli-tool/internal/pkg/config"
	"github.com/Template7/common/logger"
	"github.com/urfave/cli/v2"
)

const version = "1.0.0"

var (
	log = logger.GetLogger()
)

func InitCli() (app *cli.App) {
	config.New()
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
		stressTest,
	}
}
