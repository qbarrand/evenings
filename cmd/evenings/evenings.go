package main

import (
	"os"

	"github.com/kataras/golog"

	"github.com/qbarrand/evenings/internal/db"

	"github.com/qbarrand/evenings/internal/session"
	"github.com/qbarrand/evenings/internal/topic"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "evenings"

	app.Commands = []cli.Command{
		{
			Name:        "session",
			Description: "Manage sessions",
			Subcommands: session.CliArgs,
		},
		{
			Name:        "show-summary",
			Description: "Show a summary",
		},
		{
			Name:        "topic",
			Description: "Manage topics",
			Subcommands: topic.CliArgs,
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db",
			Value: "./evenings.db",
			Usage: "Path to the SQLite database",
		},
		cli.StringFlag{
			Name:  "log",
			Value: "info",
			Usage: "The log level",
		},
	}

	app.Before = func(c *cli.Context) error {
		logLevel := c.String("log")
		golog.Debugf("Setting the log level to %s", logLevel)
		golog.SetLevel(logLevel)

		path := c.String("db")
		golog.Debugf("Using database %s", path)
		db.SetDbPath(path)

		return nil
	}

	app.Run(os.Args)
}
