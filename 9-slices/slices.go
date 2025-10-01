package main

import "fmt"

// Variadic function

func sums(nums ...int) int {
	num := 0
	for i := 0; i < len((nums)); i ++ {
		num += (nums)[i]
	}

	return num
}

func main() {

	// a := [5]int{6, 3, 5, 6, 7} // An array
	// b := a[0:3] // A slice

	// mySlice := make([]int, 5, 10)
	// mySlice := []int{3, 4, 5, 6}

	// fmt.Println(mySlice)
	// fmt.Println("Max length",len(mySlice))
	// fmt.Println("Max cap",cap(mySlice))

	fmt.Println("Total sum is:",sums(2,4,4,5))

}
