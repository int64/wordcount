package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	count := len(strings.Fields(src))
	fmt.Println(count)
}

func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), " ")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}
