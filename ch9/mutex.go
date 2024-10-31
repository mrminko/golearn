package main

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func deposit(amt int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amt
}

func showBalance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {

}
