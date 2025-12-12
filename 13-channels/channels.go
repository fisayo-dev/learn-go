package main

import (
	"fmt"
	"time"
)

// func sendEmail(message string) {
// 	time_length := time.Millisecond * 250
// 	go func(){
// 		time.Sleep(time_length)
// 		fmt.Printf("Email received: '%s' in %v\n", message, time_length)
// 	}()
// 	fmt.Printf("Email sent: %s\n",message)
// }

func countReports(numSentCh chan int) int {
	total := 0

	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}

	return total
	
}

func test(numBatches int){
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	// It doesn't prit the numReports until the channel has been closed
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("===================")
}


func main() {
	// sendEmail("How are you?")
	// sendEmail("Is your prayer altar still burning?")
	// sendEmail("How many soldier are at the battle field?")

	test(3)
	test(4)
	test(5)
	test(6)

}

func sendReports(numBatches int, ch chan int){
	for i := range numBatches {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Send batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond*100)
	}
	close(ch)
}
