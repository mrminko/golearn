package main

import "fmt"

type rect struct {
	width  int
	height int
}

//(r rect)-> receiver, only associate with the given type

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	r := rect{
		2, 5,
	}
	area := r.area()
	fmt.Println(area)
}
