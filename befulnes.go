// Package befulnes provides tools for generating and working with lists of nonsense words
package befulnes

import (
	"fmt"
	"math/rand"
	"os/exec"
)

type WordSource interface {
	Load() *WordList
}

type WordList []string

func FromSlice(w []string) *WordList {
	wl := WordList(w)
	return &wl
}

func (w *WordList) Consume() string {
	index := rand.Intn(len(*w))
	word := (*w)[index]
	*w = append((*w)[:index], (*w)[index+1:]...)
	return word
}

func (w *WordList) Shuffle() {
	for i := len(*w) - 1; i > 0; i-- {
		swap := rand.Intn(i)
		(*w)[swap], (*w)[i] = (*w)[i], (*w)[swap]
	}
}

func (w *WordList) Length() int {
	return len(*w)
}

func GetWord(verbose, project bool) {
	sb := &Soybomb{}
	words := sb.Load()
	word := words.Consume()

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
