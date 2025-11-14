package main

import (
	"fmt"

)
const (
   first = iota
   second
   third
)

const (
   four = iota+5
   five
   six
)



func main() {

	 fmt.Println(first, second, third)
	 fmt.Println(four, five, six)

	 
	myarrX :=[...] int {2,3,4,5,6,7}
	sliceX :=myarrX[0:6]
	fmt.Printf("SliceX Length %d \n",len(sliceX))
	fmt.Printf("SliceX CAP %d \n",cap(sliceX))
		fmt.Printf("sliceX %v \n",sliceX)
	sliceY := sliceX[2:5]
		fmt.Printf("sliceY Length %d \n",len(sliceY))
	fmt.Printf("sliceY CAP %d \n",cap(sliceY))
	fmt.Printf("sliceY %v \n",sliceY)
	
	sliceZ := sliceX[1:5]
		fmt.Printf("sliceZ Length %d \n",len(sliceZ))
	fmt.Printf("sliceZ CAP %d \n ",cap(sliceZ))
	fmt.Printf("sliceZ %v \n",sliceZ)
	sliceY[1]=22
	sliceZ[1]=33
	fmt.Printf("sliceX %v \n",sliceX)

}
