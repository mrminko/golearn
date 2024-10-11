package main

import "fmt"

func main() {
	//a := map[string]int{"A": 10, "B": 20}
	b := map[string]int{"A": 10, "B": 20}
	c := map[string]int{"A": 10}
	fmt.Println(equal(b, c))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, yv := range y {
		if xv, ok := x[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}
