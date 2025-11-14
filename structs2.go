package main

import (
	"fmt"
)

type simpleCar struct {
	make  string
	model int
}

var simplecar = simpleCar{"Audi", 2020}

func changemake(carparam *simpleCar){
	carparam.make="BMW"
}

func changemodel(carparam simpleCar){
	carparam.make="BMW"
}

func main() {
	

	fmt.Printf("Simple Car Model %d\n", simplecar.model)
	fmt.Printf("Simple Car Make %s \n", simplecar.make)
	changemake(&simplecar)
		fmt.Printf("Simple Car Model %d\n", simplecar.model)
	fmt.Printf("Simple Car Make %s \n", simplecar.make)
	changemodel(simplecar)
	fmt.Printf("Simple Car Model %d\n", simplecar.model)
	fmt.Printf("Simple Car Make %s \n", simplecar.make)
	

}
