package main

import "fmt"

func sum(num ...int) int {
	var result int
	for _, val := range num {
		result += val
	}
	return result
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)
	nums := []int{1, 2, 3, 4, 5}
	result2 := sum(nums...)
	fmt.Println(result2)
}
