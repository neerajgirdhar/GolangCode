package main

import (
	"fmt"
)

func increment(number int) {
	number++
	fmt.Println("The incremented value within methods is not returned and a new variable is created for number in memory and it wont reference the original number x ")
	fmt.Printf("Value of Number in increment method is %d \n ", number)
}

func increment1(number int) int {
	number++
	return number
}

func main() {

	x := 20
	fmt.Printf("Before calling increment value of x is %d  and address of x is %d  \n", x, &x)
	increment(x)
	fmt.Printf("After calling increment value of x is %d and address of x is %d  \n", x, &x)

	fmt.Printf("Before calling increment1 value of x is %d and address of x is %d  \n", x, &x)
	x = increment1(x)
	fmt.Printf("After calling increment1 value of x is %d and address of x is %d  \n", x, &x)

}
