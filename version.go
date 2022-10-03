package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var appVersion = "v1"

func getVersion() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the version of zetools",
		Action: func(c *cli.Context) error {
			fmt.Println(appVersion)
			return nil
		},
	}
}

func setVersion(version string) {
	appVersion = version
}
