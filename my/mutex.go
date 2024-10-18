package main

import (
	"fmt"
	"math"
	"sync"
)

type raceMe struct {
	value int
	mux   *sync.Mutex
}

func (rm *raceMe) double() {
	rm.mux.Lock()
	defer rm.mux.Unlock()
	rm.value *= 2
}

func (rm *raceMe) triple() {
	rm.mux.Lock()
	defer rm.mux.Unlock()
	rm.value *= 3
}

func main() {
	var wg sync.WaitGroup
	mux := sync.Mutex{}
	race := raceMe{
		value: 1,
		mux:   &mux,
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			race.double()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			race.triple()
		}
	}()
	wg.Wait()
	fmt.Printf("Result should be %.2f.\nReal result: %d ", math.Pow(2, 5)*math.Pow(3, 5), race.value)

}
