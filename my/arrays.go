package main

import "fmt"

func arrayAppend() {
	originalArray := [3]int{0, 1, 2}
	fmt.Println("length of original array", len(originalArray), "capacity:", cap(originalArray))
	slice := originalArray[:]
	slice = append(slice, 5)
	fmt.Printf("%T", slice)
	fmt.Println("\nSlice Value:", slice)
	fmt.Println("length of original array", len(originalArray), "capacity:", cap(originalArray))

}

func arrayGrowth() {
	var originalArray = []int{0, 1, 2}
	originalArray = append(originalArray, 3, 5, 6, 7, 8, 9)
	//originalArray = append(originalArray, 3, 5, 6, 7, 8, 9)
	fmt.Println("Original Array:", originalArray)
	anotherArray := originalArray[:20]
	fmt.Println("Another Array:", anotherArray)
}

func main() {
	arrayGrowth()
}
