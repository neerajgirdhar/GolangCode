package main

import ( 
"fmt"
"time"
)

func main() {

	coffeeOrder := make(chan string,3)
	go func(){
		for i:=1;i<=5;i++ {
			order := fmt.Sprintf("Coffee Order %d",i)
			coffeeOrder  <- order
			fmt.Printf("%s Placed By Customer \n",order)
		}
		close(coffeeOrder)
	}()

	for order := range(coffeeOrder) {
		fmt.Printf("Preparing %s \n", order)
		time.Sleep(2*time.Second)
		fmt.Printf("Served %s \n", order)
	}
}
