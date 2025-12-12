package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	time_length := time.Millisecond * 250
	go func(){
		time.Sleep(time_length)
		fmt.Printf("Email received: '%s' in %v\n", message, time_length)
	}()
	fmt.Printf("Email sent: %s\n",message)
}

func main() {

	sendEmail("How are you?")
	sendEmail("Is your prayer altar still burning?")
	sendEmail("How many soldier are at the battle field?")

}
