package main

import "fmt"

func printAllPrimesBelowTheGivenNumber(maxNumber int) {
	for n := 2; n <= maxNumber; n++ {
		if n == 2 {
			fmt.Printf("%d is Prime \n", n)

		}
		if n%2 == 0 {

			continue
		}
		isprime := true
		for i := 3; i*i <= n; i++ {
			fmt.Printf("I %d N %d \n", i*i, n)
			if n%i == 0 {
				isprime = false
				break
			}
		}
		if !isprime {
			continue
		}
		fmt.Printf("%d is prime \n", n)
	}
}

func main() {

	printAllPrimesBelowTheGivenNumber(50)

	//	fmt.Println("Simple for Loop")

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("The Value of I is %d \n", i)
	//}

	//fmt.Println("Without  Condition Loop Unlimited")

	//for k := 0; ; k++ {
	//	fmt.Printf("The Value of K is %d \n", k)
	//}

	//	fmt.Println("Without  Condition Loop Unlimited althoug it will print till 99 only but loop will continue")

	//	maxValue := 100
	//	for l := 0; ; l++ {
	//		if l < maxValue {
	//			fmt.Printf("The Value of L is %d \n", l)
	//		}

	//fmt.Println("Without  Condition Loop Unlimited althoug it will print till 99 but after that it will break")

	//maxValue := 100
	//for m := 0; ; m++ {
	//	if m <= maxValue {
	//		fmt.Printf("The Value of M is %d \n", m)
	//		continue
	//	}
	//	if m > maxValue {
	//
	//		break
	//	}
	//
	//}

}
