package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
