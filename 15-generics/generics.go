package main

import "fmt"

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zeroValue T
		return zeroValue
	}
	lastEl := s[len(s)-1]
	return lastEl
}

func main() {

	lastEL1 := getLast([]int{1, 2, 3, 4, 5})
	fmt.Println("Last element:", lastEL1)
	lastEL2 := getLast([]string{})
	fmt.Println("Last element:", lastEL2)

}