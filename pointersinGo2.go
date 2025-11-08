package main

import "fmt"


func doubleValue(initValue  *int) {
	*initValue = (*initValue)*(*initValue)
	
}

func changeArray(initarray  *[5]string) {
	*initarray =[5]string{"Ram1","Sham1","Hari1","Sunita1","Sangita1"}
	
}

func changeSlice(initslice  []string) {
	fmt.Println("Slice is by default passed by reference")
	initslice[2]="2"
	
}

func changeMap(initmap map[int]string ) {
	fmt.Println("Map is by default passed by reference")
	initmap[676033]="Nzee"
	
}

func main() {
	myint := 4
	doubleValue(&myint)
	fmt.Println(myint)

	myarray:= [5]string{"Ram","Sham","Hari","Sunita","Sangita"}
	fmt.Println(myarray)
	changeArray(&myarray)
	fmt.Println(myarray)

	myslice := make([]string,3,5)
	myslice[0]="0"
	myslice[1]="1"
	fmt.Println(myslice)
	changeSlice(myslice)
	fmt.Println(myslice)

	mymap := make(map[int]string)
	mymap[676001] = "Doe"
	mymap[676002] = "David"
	
	fmt.Println(mymap)
	
	changeMap(mymap)
	fmt.Println(mymap)

}
