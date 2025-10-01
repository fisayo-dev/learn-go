package main

import "fmt"

const PI = 3.142

// func areaOfCircle (r int) float64 {
// 	return float64(r)*float64(r)*PI
// }

// func perimeterOfCircle ( r int) float64 {
// 	return 2 * float64(r) * PI 	
// }


// func incrementAge (age, increment int) int {
// 	incrementedAge := age + increment
// 	return  incrementedAge
// }


func getNames () (string, string) {
	return "Fisayo", "Obadina"
}

func main() {
	// age := 4
	// increment := 2
	// age = incrementAge(age,increment)
	fName,_ := getNames() // Ignored variable lastname
	fmt.Println(fName)


	// fmt.Println("incremented age is", age)
	// fmt.Println("The area of a circle is",areaOfCircle(4))
	// fmt.Println("The perimeter of a circle is",perimeterOfCircle(4))
}