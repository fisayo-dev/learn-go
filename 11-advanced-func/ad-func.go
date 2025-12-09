package main

import "fmt"

// High order and First-class fund

// func add(x, y int) int {
// 	return x + y
// }

// func multiply(x, y int) int {
// 	return x * y
// }

// func aggregator(a, b, c int, aggregate func(int, int) int) int {
// 	return aggregate(aggregate(a, b), c)
// }

// Closures

func adder() func (int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	fmt.Println("====================")
	defer fmt.Println("====================")
	nums := []int{2,3,4,5,5,6,7,7}
	countAdder, numAdder := adder(), adder()
	for _, num := range nums {
		fmt.Printf("The count of the %v numbers is %v \n", countAdder(1), numAdder(num))
	}

	// fmt.Println(aggregator(2,3,4, add))
	// fmt.Println(aggregator(2,3,4, multiply))

}