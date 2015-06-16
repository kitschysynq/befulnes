// Package befulnes provides tools for generating and working with lists of nonsense words
package befulnes

import (
	"fmt"
	"math/rand"
	"os/exec"
)

// WordSource is the interface that wraps the Load method for
// getting a list of words. It should probably be called Loader.
type WordSource interface {
	Load() *WordList
}

// WordList is a list of words from which a random word can be
// consumed. There are not many things here that aren't pretty
// basic list operations.
type WordList []string

// FromSlice creates a new WordList from the given slice
func FromSlice(w []string) *WordList {
	wl := WordList(w)
	return &wl
}

// Consume removes a word from the WordList and returns it
func (w *WordList) Consume() string {
	index := rand.Intn(len(*w))
	word := (*w)[index]
	*w = append((*w)[:index], (*w)[index+1:]...)
	return word
}

// Shuffle randomizes the order of the words in the WordList
func (w *WordList) Shuffle() {
	for i := len(*w) - 1; i > 0; i-- {
		swap := rand.Intn(i)
		(*w)[swap], (*w)[i] = (*w)[i], (*w)[swap]
	}
}

// AddList extends the WordList with the given words
func (w *WordList) AddList(l *WordList) {
	*w = append((*w), (*l)...)
}

// Length returns the number of words currently ready to be consumed.
func (w *WordList) Length() int {
	return len(*w)
}

// GetWord is a utility function that handles a default case of
// filling a local cache from a soybomb source and using words
// from the cache first
func GetWord(cache string, verbose, project bool) {
	words := FromFile(cache)
	for words.Length() < 5 {
		sb := &Soybomb{}
		w := sb.Load()
		words.AddList(w)
	}

	word := words.Consume()
	words.ToFile(cache)

	if verbose || !project {
		fmt.Printf("%s\n", word)
	}

	if project {
		cmd := exec.Command("git", "init", word)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error initializing git repo: %q\n", err.Error())
			return
		}
	}

	return
}
