package main

import (
	"fmt"
	"strings"
)

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

// Declaring that message is a type of pointer
func removeProfanity(message *string) {
	messageVal := *message // Destructuring message to normal string
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	messageVal = strings.ReplaceAll(messageVal, "fuck", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	// Set that memeory address of message to the modified message value
	*message = messageVal
}



func main() {
	msg1 := "Shut the fuck up"
	msg2 := "Men, shoot"
	msg3 := "What the heck is wrong with you"
	removeProfanity(&msg1)
	fmt.Println("Modified words", msg1)
	removeProfanity(&msg2)
	fmt.Println("Modified words", msg2)
	removeProfanity(&msg3)
	fmt.Println("Modified words", msg3)

	myOwnColor := car{"red"}
	myOwnColor.setColor("blue")
	println("Color after change:" , myOwnColor.color)

}

