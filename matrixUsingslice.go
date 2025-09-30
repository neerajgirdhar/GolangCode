package main

import (
	"fmt"
)

func createMatrix(row, cols int) [][]int {
	fmt.Println(row, cols)
	matrix := make([][]int, 0)
	for i := 0; i < row; i++ {
		slicerow := make([]int, 0)
		for j := 0; j < cols; j++ {

			slicerow = append(slicerow, i*j)
			//fmt.Printf("%d ", i*j)
		}
		fmt.Printf("\n")
		matrix = append(matrix, slicerow)
	}
	return matrix
}

func main() {

	matrix := createMatrix(3, 3)

	matrix1 := createMatrix(3, 4)
	fmt.Println(matrix)

	fmt.Println(matrix1)

}
