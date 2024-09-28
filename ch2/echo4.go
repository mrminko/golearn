package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit new line")

var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Print(args)
	fmt.Print(strings.Join(args, *sep))
	if !*n {
		fmt.Println()
	}
}
