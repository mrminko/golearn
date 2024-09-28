package main

import "fmt"

func main() {
	a := [...]string{
		1: "A",
		2: "B",
		3: "C",
		4: "D",
		5: "E",
	}
	b := a[1:3]
	fmt.Println(cap(b))
}
