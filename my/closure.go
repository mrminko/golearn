package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d"}
	ch := make(chan struct{})
	for _, v := range list {
		go func(v string) {
			fmt.Println(v)
			ch <- struct{}{}
		}(v)
	}
	for range list {
		<-ch
	}
	return
}
