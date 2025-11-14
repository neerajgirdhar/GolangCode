package main

import (
	"fmt"
	"time"
)


func printTogether(){
go func (){
		fmt.Println("Printing Even")
	for i:=1 ;i<10;i++{
		if(i%2 ==0) {
			fmt.Println(i)
			time.Sleep(50*time.Millisecond)
		}
		
	}
}()

go func (){
	fmt.Println("Printing odd")
	for i:=1 ;i<10;i++{
		if(i%2 !=0) {
			fmt.Println(i)
			time.Sleep(50*time.Millisecond)
		}
		
	}
}()
}


func main() {

//Look These functions run concurrently and its not the traditional way as these 2 functions are not dependent
 printTogether()
time.Sleep(1*time.Second)
	
}
