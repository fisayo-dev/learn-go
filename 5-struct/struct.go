package main

import "fmt"

// Structs can hold any types and be nested

// Named structs
// Anonymous structs

// type Res struct {
// 	Data
// 	message string
// 	sucess  bool
// }

// type Data struct {
// 	users      string
// 	totalUsers int
// }

// func sendResponse(r Res) {
// 	fmt.Println("This is the response sent")
// 	fmt.Printf("Data: \n Users Name: %v \n Total Users Length: %d \nMessage: %s\nSuccess: %v ", r.users, r.totalUsers, r.message, r.sucess)
// }

// Creating methods in struct

type rect struct {
	width  float64
	height float64
}

// Assign area func as a method to rect struct.
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width * r.height
}

func returnAreaDetails(r rect) {
	fmt.Println("The width of a rectangle is:", r.width)
	fmt.Println("The height of a rectangle is:", r.height)
	fmt.Println("The area of a rectangle is:", r.area())
	fmt.Println("The perimeter of a rectangle is:", r.perimeter())
	fmt.Println("============================")
}

func main() {
	// sendResponse(Res{
	// 	Data: Data{users: "Fisayo, Obadina, Anna, Taye, Oluche Oche",
	// 		totalUsers: 10},
	// 	message: "Some dummy text message",
	// 	sucess:  false,
	// })

	r1 := rect{
		height: 4,
		width:  5,
	}
	r2 := rect{
		height: 10,
		width:  15,
	}
	r3 := rect{
		height: 3,
		width:  7,
	}
	r4 := rect{
		height: 6,
		width:  12,
	}
	r5 := rect{
		height: 45,
		width:  12,
	}

	returnAreaDetails(r1)
	returnAreaDetails(r2)
	returnAreaDetails(r3)
	returnAreaDetails(r4)
	returnAreaDetails(r5)
}
