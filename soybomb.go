package befulnes

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type Soybomb struct{}

func (sb *Soybomb) Load() *WordList {
	resp, err := http.Get("http://www.soybomb.com/tricks/words")
	if err != nil {
		fmt.Printf("Error grabbing wordlist: %s\n", err)
		return nil
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil
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

	return FromSlice(words)
}
