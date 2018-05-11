package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/zaru/rep/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "rep"
	app.Usage = "Initial setting of GitHub repository"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "rep init ./config.json / initialize to GitHub repository",
			Action: func(c *cli.Context) error {
				commands.Init()
				return nil
			},
		},
	}
	app.Run(os.Args)
}
