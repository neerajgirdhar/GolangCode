package main

import (
	"errors"
	"fmt"
)

type divideErrorCustom struct {
	dividend float64
	divisor  int
}

func (de divideErrorCustom) Error() string {
	return fmt.Sprintf("Custom Error which implements the Error Interface : Can't divide by zero: Dividend : (%f) Divisor %d", de.dividend, de.divisor)
}

func divide(dividend float64, divisor int) (quotient float64, error error) {
	if divisor == 0 {
		return 0.0, divideErrorCustom{dividend, divisor}
	}
	return dividend / float64(divisor), nil
}

func divide1(dividend float64, divisor int) (quotient float64, error error) {
	fmt.Println("This is another way to define error using error package")
	if divisor == 0 {

		return 0.0, errors.New("Can't divide by zero")
	}
	return dividend / float64(divisor), nil
}

func main() {
	fmt.Println()
	fmt.Println("Custom Error Interface Demo")

	quotient, err := divide(100.000, 20)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("outcome Quotient %.2f \n", quotient)
	}

	quotient1, err := divide(100.000, 0)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("outcome Quotient %.2f /n", quotient1)
	}

	quotient2, err := divide1(100.000, 20)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("outcome Quotient %.2f /n", quotient2)
	}

	quotient3, err := divide1(100.000, 0)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("outcome Quotient %.2f /n", quotient3)
	}
}
