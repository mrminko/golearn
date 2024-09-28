package main

import "fmt"

func main() {
	a := [3]int{USD: 1, EUR: 2}
	b := [30]int{29: -1}
	fmt.Println(a, b)
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)
