package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	sum1 := 0
	//another way to right complete for loop on array or slice

	for _, number := range numbers {
		sum += number
	}

	for i := 0; i < len(numbers); i++ {
		sum1 += numbers[i]
	}
	if sum == sum1 {
		return sum
	} else {
		return 0
	}

}

func main() {
	//With variadic operator you can pass as many variables
	sumofNumbers := sum(1, 2, 3, 4, 5)
	fmt.Println(sumofNumbers)

	slice := []int{1, 2, 3, 4, 5}

	//With spread operator you can pass slice or some values of slices
	sumofNumbersusingSpredOperation := sum(slice...)
	fmt.Println(sumofNumbersusingSpredOperation)

	//With spread operator you can pass slice or some values of slices
	sumofNumbersusingSpredOperation1 := sum(slice[2:4]...)
	fmt.Println(sumofNumbersusingSpredOperation1)
}
