package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func aggregate(a, b, c int, f func(x, y int) int) int {
	result := f(f(a, b), c)
	return result
}

// /Currying
func addMaker(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	fmt.Println(aggregate(1, 2, 3, add))
	addTwo := addMaker(2)
	addThree := addMaker(3)
	fmt.Println(addTwo(10))
	fmt.Println(addThree(10))
}
