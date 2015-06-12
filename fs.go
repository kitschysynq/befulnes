package befulnes

import (
	"bufio"
	"os"
	"unicode/utf8"
)

const (
	unitSeparator   byte = 0x1f
	recordSeparator byte = 0x1e
	groupSeparator  byte = 0x1d
	fileSeparator   byte = 0x1c
)

func (w *WordList) ToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	for _, v := range *w {
		_, err = file.Write([]byte(v))
		if err != nil {
			return err
		}
		_, err = file.Write([]byte{unitSeparator})
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *WordList) FromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	// Create a custom split function to split at unitSeparator
	split := func(data []byte, atEOF bool) (int, []byte, error) {
		// Scan until unitSeparator, marking end of word.
		for width, i := 0, 0; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if r == rune(unitSeparator) {
				return i + width, data[:i], nil
			}
		}
		// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
		if atEOF && len(data) > 0 {
			return len(data), data[:], nil
		}
		return 0, nil, nil
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		word := scanner.Text()
		if err := scanner.Err(); err != nil {
			return err
		}
		(*w) = append((*w), word)
	}
	return nil
}
