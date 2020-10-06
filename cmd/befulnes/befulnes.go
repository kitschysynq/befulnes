// Provides a tool to generate a project name
package main

import (
	"os"
	"path"

	"github.com/urfave/cli"

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
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "project",
			Usage: "create a git project",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "always display selection",
		},
	}
	app.Run(os.Args)
}

func getWord(c *cli.Context) {
	cacheDir := os.Getenv("HOME")
	befulnes.GetWord(path.Join(cacheDir, ".befulnes.cache"), c.Bool("verbose"), c.Bool("project"))
}
