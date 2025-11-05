package main

import "fmt"

func main() {
	myint := 42
	pointervar := &myint
	fmt.Printf("myint is of type %T and value is %v and address is  %v \n", myint, myint, &myint)
	fmt.Printf("pointervar is of type %T and value is %v and address is %v \n", pointervar, *pointervar, &pointervar)
	*pointervar = 10
	fmt.Println("changed the value of myint from 5 to 10 by using pointer variable pointervar")
	fmt.Printf("myint is of type %T and value is %v and address is  %v \n", myint, myint, &myint)
	fmt.Printf("pointervar is of type %T and value is %v and address is %v \n", pointervar, pointervar, &pointervar)

}
