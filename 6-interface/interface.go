package main

import "fmt"

// Rectangle types
type rect struct {
	width  float64
	height float64
}

type shape interface {
	area() float64
	perimeter() float64
}

const PI = 3.142

// Area and perimter methods to rectangle struct
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// Circle types
type circle struct {
	radius float64
}

// Area and perimter methods to circle struct
func (c circle) area() float64 {
	return PI * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * PI * c.radius
}

func getShapeDetails(s shape) {
	// Assertion doene with if else

	// // Check is shape is a cricle - Type asserion
	// // If shape is cicle -  Print radius of circle
	// c, ok := s.(circle)
	// if !ok {
	// 	fmt.Println("This radius of circle is", c.radius)
	// 	fmt.Println("============================")
	// }
	// r, ok := s.(rect)
	// // Else - print width and height of rectangle
	// if ok {
	// 	fmt.Println("This width of the reactangle is", r.width)
	// 	fmt.Println("This height of the reactangle is", r.height)
	// 	fmt.Println("============================")
	// }

	// Assertion by Switch case
	switch v := s.(type) {
	case rect:
		fmt.Println("This width of the reactangle is", v.width)
		fmt.Println("This height of the reactangle is", v.height)
		fmt.Printf("This area of the reactangle is %.2f\n", v.area())
		fmt.Println("============================")
	case circle:
		fmt.Println("This radius of circle is", v.radius)
		fmt.Printf("This area of the cricle is %.2f\n", v.area())
		fmt.Println("============================")
	default:
		fmt.Println("This shape is neither a circle nor a rectangle")
	}

}

// Because getShapeDetails func is an interface we can pass in any type as long has it implements that interface.

func main() {

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
	c1 := circle{
		radius: 15,
	}
	c2 := circle{
		radius: 7,
	}

	getShapeDetails(r1)
	getShapeDetails(r2)
	getShapeDetails(r3)
	getShapeDetails(c1)
	getShapeDetails(c2)
}
