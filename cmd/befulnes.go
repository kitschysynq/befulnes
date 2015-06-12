// Provides a tool to generate a project name
package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/kitschysynq/befulnes"
)

var (
	version string = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "befulnes"
	app.Usage = "save time naming your project"
	app.Action = getWord
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "project",
			Usage: "create a git project",
		},
	}
	app.Run(os.Args)
}

func getWord(c *cli.Context) {
	befulnes.GetWord(c.Bool("verbose"), c.Bool("project"))
}
