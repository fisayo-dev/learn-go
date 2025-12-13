package main

import "fmt"

// func calculateSchoolFees(students int) float64 {
// 	totalCost := 0.0
// 	for i := 0; i < students; i++ {
// 		totalCost += (10.0 * float64(i))
// 	}
// 	return totalCost
// }

// func test(students int) {
// 	fmt.Printf("Caluclating school fees for %v students\n", students)
// 	schoolFees := calculateSchoolFees(students)
// 	fmt.Printf("The total school fees is: %.2f\n", schoolFees)
// 	fmt.Println("=======================")
// }

// func fizzbuzz() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("Fizz buzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}

// }

func printPrimes (max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n % 2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i * i < n+1; i++ {
			if n % i == 0 {
				isPrime = false 
				break
			}
		}
		if (!isPrime) {
			continue
		}

		fmt.Println(n)
	}
}

func main() {
	// test(2)
	// test(3)
	// test(5)

	// plantHeight := 5
	// for plantHeight < 5 {
	// 	fmt.Println("plant height is", plantHeight)
	// 	plantHeight++
	// }
	// fmt.Println("Plant has grown to", plantHeight, "inches")

	printPrimes(10)
}
