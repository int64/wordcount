package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	src, err := readInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	counter := 0
	for _, word := range strings.Fields(src) {
		if isWord(word) {
			counter++
		}
	}
	fmt.Println(counter)
}

func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), " ")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}

func isWord(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
