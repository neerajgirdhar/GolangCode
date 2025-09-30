package main

import (
	"fmt"
)

var arr1 [3]int

//declartion outside methods or at package level
//var primes = [4]int {2,3,5,7}

func main() {
	//declartion inside methods
	primes := [4]int{2, 3, 5, 7}

	fmt.Printf("arr1 Length: %v \n", len(arr1))
	fmt.Printf("arr1 Cap: %v \n", cap(arr1))
	fmt.Printf("0th Index value %v \n", arr1[0])

	fmt.Printf("primes Length: %v \n", len(primes))
	fmt.Printf("primes Cap: %v \n", cap(primes))
	fmt.Printf("First Prime Number %v \n", primes[0])

	// making slice from array with all values
	slice1 := primes[0:4]
	fmt.Printf("slice1 Length: %v \n", len(slice1))
	fmt.Printf("slice1 Cap: %v \n", cap(slice1))
	fmt.Printf("0th Index value slice1 %v \n", slice1[0])

	// making slice from array with all values 2nd way
	slice2 := primes[:]
	fmt.Printf("slice2 Length: %v \n", len(slice2))
	fmt.Printf("slice2 Cap: %v \n", cap(slice2))
	fmt.Printf("Ist Index value slice2 %v \n", slice2[1])

	// making slice from array with only first 2 value
	slice3 := slice2[0:2]
	fmt.Printf("slice3 Length: %v \n", len(slice3))
	fmt.Printf("slice3 Cap: %v \n", cap(slice3))
	fmt.Printf("Ist Index value slice3 %v \n", slice3[1])

	//indiviually creating slice with length and cap this will allocate 5 space to
	//underlying array which is referenced by slice 4
	slice4 := make([]float64, 3, 5)
	fmt.Printf("slice4 Length: %v \n", len(slice4))
	fmt.Printf("slice4 Cap: %v \n", cap(slice4))
	fmt.Printf("Ist Index value slice4 %v \n", slice4[1])

	//indiviually creating slice with length and no cap this will allocate default space to
	//underlying array which is referenced by slice 4
	slice5 := make([]float64, 3)
	fmt.Printf("slice5 Length: %v \n", len(slice5))
	fmt.Printf("slice5 Cap: %v \n", cap(slice5))
	fmt.Printf("Ist Index value slice5 %v \n", slice5[1])

	//Now append values method will rebase slice5 to the new underlying array
	for i := 0; i <= 20; i++ {
		slice5 = append(slice5, float64(i))
	}

	for i := 1; i <= len(slice5)-1; i++ {
		fmt.Println(slice5[i])
	}

	fmt.Printf("slice5 Length: %v \n", len(slice5))
	fmt.Printf("slice5 Cap: %v \n", cap(slice5))
	fmt.Printf("5th Index value slice5 %v \n", slice5[5])

	slice11 := []int{1, 2, 3, 4, 5}
	slice12 := []int{6, 7, 8, 9, 10}
	slice13 := []int{11, 12, 13, 14, 15}

	fmt.Println("address of slice11 [0] is", &slice11[0])
	fmt.Println("address of slice11 [1] is", &slice11[1])
	fmt.Println("address of slice11 [2] is", &slice11[2])
	fmt.Println("address of slice11 [3] is", &slice11[3])
	fmt.Println("address of slice11 [4] is", &slice11[4])

	//We should/must append in Same Slice refer 3.30-3.35 video
	slice11 = append(slice11, slice12...)

	fmt.Println("address of slice11 [0] after append is done ", &slice11[0])
	fmt.Println("address of slice11 [1] after append is done ", &slice11[1])
	fmt.Println("address of slice11 [2] after append is done ", &slice11[2])
	fmt.Println("address of slice11 [3] after append is done ", &slice11[3])
	fmt.Println("address of slice11 [4] after append is done ", &slice11[4])

	slice11 = append(slice11, slice13...)
	for i := 0; i <= len(slice11)-1; i++ {
		fmt.Println(slice11[i])
	}

	slice15 := make([]int, 3, 8)
	slice15 = append(slice15, 1)
	slice15 = append(slice15, 2)
	slice15 = append(slice15, 3)

	slice16 := append(slice15, 4)
	fmt.Println("Value of slice16  is", slice16)
	slice17 := append(slice15, 5)
	fmt.Println("Value of slice16  is:HERE VALUE OF SLICE 15 ALSO GOT CHANGED BECAUSE slice 15 had cap of 8 ", slice16)
	fmt.Println("Value of slice17 is", slice17)

	slice18 := make([]int, 3)
	slice18 = append(slice18, 1)
	slice18 = append(slice18, 2)
	slice18 = append(slice18, 3)

	slice19 := append(slice18, 4)
	fmt.Println("Value of slice19  is", slice19)
	slice20 := append(slice18, 5)
	fmt.Println("Value of slice19  is:HERE VALUE OF SLICE 18 did not changed as cap was not defined ", slice19)
	fmt.Println("Value of slice20 is", slice20)

	fmt.Println("Usage of Range Method")
	for i, value := range slice20 {
		fmt.Printf("Index :: %d, Value :: %d, Memoryaddress :: %d \n", i, value, &value)
	}

	initializedSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("initializedSlice:", initializedSlice)
	sliceusingmakeMethod := make([]int, 4, 5)
	fmt.Println("sliceusingmakeMethod:", sliceusingmakeMethod)
}
