package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Author = "Oleg Lebedev <oolebedev@gmail.com>"
	app.Name = "srlt"
	app.Usage = "save and restore repositories at given path"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Value: ".",
			Usage: "path to scan and restore, will be saved at first time",
		},
		cli.StringFlag{
			Name:  "file, f",
			Value: "srlt.yaml",
			Usage: "filename for read and write",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "verbose mode",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "snapshot",
			Aliases: []string{"s"},
			Usage:   "Save your current state into the file",
			Action:  snapshotAction,
		},
		{
			Name:    "restore",
			Aliases: []string{"r"},
			Usage:   "Restore state from the file",
			Action:  restoreAction,
		},
		{
			Name:    "exec",
			Aliases: []string{"e"},
			Usage:   "Execute give shell programm to each dependency",
			Action:  execAction,
		},
	}

	app.Run(os.Args)
}
