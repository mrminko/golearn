package main

import "fmt"

func main() {
	d := []int{0, 1, 2, 3, 4, 5}
	s := make([]int, 10)
	n := copy(s, d)
	fmt.Println(s, n, d)
	fmt.Println(remove(d, 3))
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
