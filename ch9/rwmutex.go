package main

import "sync"

var (
	rwmu     sync.RWMutex
	balance1 int
)

func deposit1(amt int) {
	rwmu.Lock()
	defer rwmu.Unlock()
	balance1 += amt
}

func showBalance1() int {
	rwmu.RLock()
	defer rwmu.RUnlock()
	return balance1
}

func main() {

}
