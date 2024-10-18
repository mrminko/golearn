package main

import "fmt"

/*
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately,
but the function call is not executed until the surrounding function returns.
*/

func printDefer() string {
	a := 10
	defer fmt.Println("defer print", a)
	a++
	return fmt.Sprintln("final print", a)

}

func main() {
	result := printDefer()
	fmt.Println(result)
}
