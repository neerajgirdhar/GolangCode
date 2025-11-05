package main

import "fmt"

func multiply(num1, num2 int) int {
	return num1 * num2
}
func add(num1, num2 int) int {
	return num1 + num2
}

func aggregator(num1, num2, num3 int, funcasparameter func(a, b int) int) int {
	return funcasparameter(funcasparameter(num1, num2), num3)
}
func main() {
	fmt.Printf("Aggreator function(high order function) is called by passing first class function which is passed as data in this case first class function is multiply and highorder function is aggregator %d\n ", aggregator(2, 3, 4, multiply))
	fmt.Printf("Aggreator function(high order function) is called by passing first class function which is passed as data in this case first class function is add and highorder function is aggregator %d\n ", aggregator(2, 3, 4, add))
}
