package main

import "fmt"

func multiply1(num1, num2 int) int {
	return num1 * num2
}
func add1(num1, num2 int) int {
	return num1 + num2
}

func printValuesSepratedByCommon(str1, str2 string) string {
	return str1 + "," + str2
}

func printValuesSepratedByColon(str1, str2 string) string {
	return str1 + ":" + str2
}

func highOrderFunction(funcAsParameter func(a, b int) int) (custFunc func(aa, bb, cc int) int) {
	return func(aa, bb, cc int) int {
		return funcAsParameter(aa, bb) - cc
	}
}

func highOrderFunction1(funcAsParameter func(a, b string) string) (custFunc func(aa, bb, cc string) string) {
	return func(aa, bb, cc string) string {
		return funcAsParameter(aa, bb) + " :: Category " + cc
	}
}
func main() {

	squareandmultiplyby10 := highOrderFunction(multiply1)

	fmt.Printf("%d\n", squareandmultiplyby10(7, 10, 5))
	//fmt.Printf("Aggreator function(high order function) is called by passing first class function which is passed as data in this case first class function is multiply and highorder function is aggregator %d\n ", aggregator(2, 3, 4, multiply))
	//fmt.Printf("Aggreator function(high order function) is called by passing first class function which is passed as data in this case first class function is add and highorder function is aggregator %d\n ", aggregator(2, 3, 4, add))
	printExceptionLogSeperatedByColon := highOrderFunction1(printValuesSepratedByColon)
	fmt.Println(printExceptionLogSeperatedByColon("IndexOutOFBoundException", "This is index Out of BoundException", "Exception"))

	printErrorLogSeperatedByComma := highOrderFunction1(printValuesSepratedByCommon)
	fmt.Println(printErrorLogSeperatedByComma("ClassNotFoundError", "This is Class Not Found Error", "ERROR"))

}
