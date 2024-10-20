package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Rocket Launched Successfully")
}

func abortLaunch() {
	fmt.Println("Launch Aborted")
}

func main() {
	fmt.Println("Countdown started")
	//tick := time.Tick(time.Second)

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
		fmt.Println("Abort")
	}()
	select {
	case <-time.After(time.Second * 3):

	case <-abort:
		abortLaunch()
		return
	}
	//for countdown := 10; countdown > 0; countdown-- {
	//	fmt.Println(countdown)
	//	<-tick
	//}
	launch()
}
