package main

import (
	"fmt"
)

func deferfunct1() {
	fmt.Println("Defer Function 1")
}
func deferfunct2() {
	fmt.Println("Defer Function 2")
}
func deferfunct3() {
	fmt.Println("Defer Function 3")
}
func deferfunct4() {
	fmt.Println("Defer Function 4")
}

func callFunctionFirstandThenDefer() {
	fmt.Println("callFunctionFirstandThenDefer Statment 1")
	fmt.Println("callFunctionFirstandThenDefer Statment2")
	fmt.Println("callFunctionFirstandThenDefer Statment 3")
	defer deferfunct1()
	defer deferfunct2()
	defer deferfunct3()
	defer deferfunct4()

	fmt.Println("callFunctionFirstandThenDefer Statment 4")
}

func handlePanic() {
	if a := recover(); a != nil {
		fmt.Println("Recovered from PAnic")
	}
}

func checkEmployeeDetails(name, empid string) string {

	//if you comment below line progam will go into panic condition
	//handPanic will recover from panic its like catching exception in java
	defer handlePanic()

	if name == "" {
		panic("Name is Empty")
	}
	if empid == "" {
		panic("empid is Empty")
	}
	return "Validated"
}

func main() {

	checkEmployeeDetails("Ramesh", "")
	checkEmployeeDetails("", "56789")
	checkEmployeeDetails("Neer", "56789")
	callFunctionFirstandThenDefer()

	fmt.Println("Contol Returned to main")

	//HERE DEFER FUNCTIONS ARE NOT CALLED BECAUSE PANIC IS RAISED BEFORE DEFER

	//control does not go to a defer statement after panic is called if the defer statement
	// is placed after the panic statement. The panic statement halts the normal execution
	// flow of the function immediately. If a defer statement is located after the panic statement,
	// the control will never reach it because the panic will prevent it from being registered.

}
