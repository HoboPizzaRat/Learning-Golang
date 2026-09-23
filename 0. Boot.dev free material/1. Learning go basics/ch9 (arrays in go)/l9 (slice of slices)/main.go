package main

import "fmt"

// Slices can hold other slices,
// effectively creating a matrix, or a 2D slice

// Complete the createMatrix function
// its a 2d slice of integers
// each cell value is calculated by i * j
// i is for row
// j is for column

func createMatrix(rows, cols int) [][]int {

	var matrix [][]int = [][]int{}

	for i := 0; i < rows; i++ {
		var row []int = make([]int, cols)

		for j := 0; j < cols; j++ {
			row[j] = j * i
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func main() {
	var rows int = 5
	var cols int = 10

	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Printf("\n")

	}
}
