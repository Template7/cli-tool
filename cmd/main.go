package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

const (
	version = "v1.0.0"
)

func main() {
	app := &cli.App{
		Name:     "Template7 Internal CLI Tool",
		Version:  version,
		Commands: []*cli.Command{},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(-1)
	}
}
