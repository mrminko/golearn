package main

import (
	"errors"
	"fmt"
)

type zeroDivision struct {
	dividend float64
}

func (zd zeroDivision) Error() string {
	return fmt.Sprintf("Cannot divide %v with zero", zd.dividend)
}

func divide(dividend float64, divisor float64) (quotient float64, err error) {
	if divisor == 0 {
		return 0, zeroDivision{dividend}
	}
	return dividend / divisor, nil
}

func divide2(dividend float64, divisor float64) (quotient float64, err error) {
	if divisor == 0 {
		return 0, errors.New(fmt.Sprintf("Cannot divide %v with zero", dividend))
	}
	return dividend / divisor, nil
}

func main() {
	q, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred=>", err)
		return
	}
	fmt.Println("Quotient=", q)

}
