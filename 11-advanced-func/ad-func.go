package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func aggregator(a, b, c int, aggregate func(int, int) int) int {
	return aggregate(aggregate(a, b), c)
}

func main() {
	fmt.Println(aggregator(2,3,4, add))
	fmt.Println(aggregator(2,3,4, multiply))

}