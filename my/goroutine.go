package main

import (
	"fmt"
	"time"
)

func concurrent() {
	fmt.Println("Function Start")
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Println("I have slept enough")
	}()
	fmt.Println("After sleep")
}

func main() {
	concurrent()
	time.Sleep(time.Millisecond * 300)
}
