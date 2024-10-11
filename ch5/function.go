package main

import "fmt"

func main() {
	fmt.Println(add2(3, 4))
}

func add(x, y int) (z int) {
	return x + y
}

func add2(int, int) int {
	return 2
}
