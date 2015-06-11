// Provides a tool to generate a project name
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"golang.org/x/net/html"
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
	resp, err := http.Get("http://www.soybomb.com/tricks/words")
	if err != nil {
		fmt.Printf("Error grabbing wordlist: %s\n", err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	var words []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "td" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "b" && c.FirstChild.Type == html.TextNode {
					words = append(words, c.FirstChild.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	word := words[rand.Intn(len(words))]

	if c.Bool("verbose") || !c.Bool("project") {
		fmt.Printf("%s\n", word)
	}

	if c.Bool("project") {
		cmd := exec.Command("git", "init", word)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("error initializing git repo: %q\n", err.Error())
			return
		}
	}

	return
}
