package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			ch <- i
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		<-ch
	}
	return
}
