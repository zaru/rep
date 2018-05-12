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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Value: "config.toml",
					Usage: "Specify the configuration file path",
				},
			},
			Action: func(c *cli.Context) error {
				commands.Init(c.String("config"))
				return nil
			},
		},
	}
	app.Run(os.Args)
}
