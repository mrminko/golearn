package main

import "fmt"

func main() {
	//var runes []rune
	//str := "Hello, 世界"
	//for _, v := range str {
	//	runes = append(runes, v)
	//}
	//fmt.Printf("%c\n", runes)
	////fmt.Println(runes)
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap <= len(x)*2 {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
