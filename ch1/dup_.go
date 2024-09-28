package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; {
		input.Scan()
		line := input.Text()
		//count[line] = count[line] + 1
		count[line] += 1
		i += 1
	}
	for l, i := range count {
		fmt.Printf("%s\t%d\n", l, i)
	}
}
