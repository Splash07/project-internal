package main

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-internal/web"
)

var cfg = config.GetConfig()

func main() {

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "internal",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				return web.SSO.Start()
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	app.Run(os.Args)
}
