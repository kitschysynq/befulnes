// Package befulnes provides tools for generating and working with lists of nonsense words
package befulnes

import (
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"

	"golang.org/x/net/html"
)

func GetWord(verbose, project bool) {
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

	if verbose || !project {
		fmt.Printf("%s\n", word)
	}

	if project {
		cmd := exec.Command("git", "init", word)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("error initializing git repo: %q\n", err.Error())
			return
		}
	}

	return
}
