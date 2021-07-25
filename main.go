package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	fmt.Println(len(strings.Fields(strings.Join(flag.Args(), " "))))
}
