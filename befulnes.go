// Provides a tool to generate a project name
package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
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
	fmt.Printf("%s\n", words[rand.Intn(len(words))])
	return
}
