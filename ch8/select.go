package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func ints(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func alpha(ch chan<- string) {
	list := []string{"a", "b", "c", "d", "e"}
	defer wg.Done()
	for _, v := range list {
		ch <- v
	}
	close(ch)
}

func main() {
	intch := make(chan int)
	alphach := make(chan string)
	wg.Add(2)
	var close1, close2 bool
	go ints(intch)
	go alpha(alphach)
	go func() {
		for {
			if close1 && close2 {
				break
			}
			select {
			case i, ok := <-intch:
				if !ok {
					fmt.Println("Closed IntCh")
					close1 = true
				}
				fmt.Println("From IntCh", i)
			case a, ok := <-alphach:
				if !ok {
					fmt.Println("Closed AlphaCh")
					close2 = true
				}
				fmt.Println("From AlphaCh", a)
			}
		}

	}()
	wg.Wait()
}
