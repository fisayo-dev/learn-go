package main

import "fmt"

// Variadic function

// func sums(nums ...int) int {
// 	num := 0
// 	for i := 0; i < len((nums)); i++ {
// 		num += nums[i]
// 	}

// 	return num
// }

// Create matrix
func createMatrix (rows, cols int) [][]int {
	matrix := make([][]int,0)
	for i := 0; i < rows; i++ {
		row := make([]int,0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j) 
		}
		matrix = append(matrix, row)
	}

	return matrix
}
// Create multiplication table
func createMultiplicationTable (num, limit int) [][]int {
	table := make([][]int,0)
	for i := 1; i <= num; i++ {
		row := make([]int,0)
		for j := 1; j <= limit; j++ {
			row = append(row, i*j) 
		}
		table = append(table, row)
	}

	return table
}

func showUserNames(names []string) ([]int, []string) {
	indexedNames := []int{}
	sortedNames := []string{}
	// Alternative (olden method is doing ) for i: = 0; i< len(names); i ++ 
	for i, name := range names  { 
		indexedNames = append(indexedNames, i+1)
		sortedNames = append(sortedNames, name)
	}
	
	return indexedNames, sortedNames
}

func main() {

	// a := [5]int{6, 3, 5, 6, 7} // An array
	// b := a[0:3] // A slice

	// mySlice := make([]int, 5, 10)
	// mySlice := []int{3, 4, 5, 6}

	// fmt.Println(mySlice)
	// fmt.Println("Max length",len(mySlice))
	// fmt.Println("Max cap",cap(mySlice))

	// Or
	// nums := []int{2, 3, 4, 5, 6}
	// nums = append(nums, 3, 10, 20)

	// fmt.Println("Total sum is:", sums(nums...))

	// fmt.Println(createMatrix(10,10))
	// fmt.Println(createMultiplicationTable(10,12))
	msg := []string{"daniel", "moses", "joseph"}
	fmt.Println(showUserNames(msg))
}
