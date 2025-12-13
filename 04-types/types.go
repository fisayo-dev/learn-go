package main

import "fmt"

func main () {

	const ceo = "Mark zuckerberg"
	const weight = 51.654
	
	msg := fmt.Sprintf("Hey!, Did you know that %s weighs %.2fkg", ceo, weight)
	
	fmt.Println(msg)
}