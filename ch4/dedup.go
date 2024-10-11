package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dedup()
}

func dedup() {
	seen := make(map[string]bool)
	//fmt.Println(seen)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
}
