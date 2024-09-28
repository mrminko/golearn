package main

func reverse(a []int) {
	for i, j := 0, len(a); i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [...]int{1, 2, 3}
	reverse(a[:])

}
