package main

import "fmt"

func main() {
	a := [3]int{1, 1, 1}
	changeArray(a)
	fmt.Println(a) // [1,1,1]
	changeArrayPtr(&a)
	fmt.Println(a) // [1,3,1]
}

func changeArray(array [3]int) {
	array[1] = 3
}
func changeArrayPtr(ptr *[3]int) {
	ptr[1] = 3
}
