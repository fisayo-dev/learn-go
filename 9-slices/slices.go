package main

import "fmt"

func main() {

	a := [5]int{6, 3, 5, 6, 7} // An array
	b := a[0:3] // A slice 

	fmt.Println(b)

}
