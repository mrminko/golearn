package main

import "fmt"

/*
	type error interface {
		Error() string
	}
*/
func main() {
	//msg1 := "Hello, how are you? I hope you are happy and healthy. As for me, I'm fine."
	msg2 := "Mingalarbar"
	cost, err := sendSms(msg2)
	if err != nil {
		fmt.Printf("Error occurred when sending SMS, %s", err)
		return
	}
	fmt.Printf("Sucessfully sent a message. Cost is %.2f", cost)

}

/*
returning the error should be accompanied by zero values of other return values
*/
func sendSms(text string) (float64, error) {
	const maxLength = 25
	const costPerChar = 10.00
	if len(text) > maxLength {
		return 0.0, fmt.Errorf("excedding max length")
	}
	return float64(len(text)) * costPerChar, nil
}
