package main

import (
	"fmt"
	"os"
)

func main() {
	a := 0
	num, str := "", ""
	for i := 0; i < len(os.Args); i++ {
		num = os.Args[1]
		str = os.Args[2]
		fmt.Println(num + ". " + str)
		fmt.Println(a)

	}
}
