package main

func appendInt(a []int, b int) []int {
	var z []int
	zlen := len(a) + 1
	if zlen <= cap(a) {
		z = a[:zlen]
	} else {
		zcap := len(a)
		zcap *= 2
		z = make([]int, zlen, zcap)
		copy(z, a)
	}
	z[zlen] = b
	return z
}

//hi
