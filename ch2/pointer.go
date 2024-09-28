package main

import "fmt"

//func main() {
//	x := 1
//	p := &x
//	fmt.Println("Address, %v", p)
//}

func main() {
	v := 1
	inc(&v)
	fmt.Println("Value of v,", v)
}

func inc(p *int) int {
	*p++
	return *p
}
