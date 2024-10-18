package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for x := 0; x < 100; x++ {
			ch1 <- x
		}
		close(ch1)
	}()
	go func() {
		for v := range ch1 {
			ch2 <- v * v
		}
		close(ch2)
	}()
	for v := range ch2 {
		fmt.Println(v)
	}

}
