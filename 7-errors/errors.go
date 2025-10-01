package main

import (
	"errors"
	"fmt"
)

// func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
// 	cost1, err := sendSMS(msgToCustomer)
// 	if err != nil {
// 		return 0.0, err
// 	}

// 	cost2, err := sendSMS(msgToSpouse)
// 	if err != nil {
// 		return 0.0, err
// 	}

// 	totalMsgCost := cost1 + cost2
// 	return totalMsgCost, err
// }

// func sendSMS(message string) (int, error) {
// 	const maxTextLen = 25
// 	const costPerChar = 2
// 	if len(message) > maxTextLen {
// 		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
// 	}
// 	return costPerChar * len(message), nil
// }

// Error interaces

// Creating printError struct
// type printError struct {
// 	message string
// }

// // Assigning error method to struct
// func (pE printError) Error() string {
// 	// Returning custom error message
// 	return fmt.Sprintf("The text '%v' is too long (more than 20 chars)", pE.message)
// }

func printText(text string) (string, error) {
	if len(text) > 20 {
		// return "", printError{message: text} // using error interface
		return "", errors.New("This text is too long")
	}

	return text, nil

}

func main() {
	// sentMessageCost, err := sendSMSToCouple("How are you doing today?", "How is your spouse?")
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	}
	// 	fmt.Printf("Total message cost is: %d", sentMessageCost)

	text1, err := printText("I am a boy with a good girl")
	if err != nil {
		fmt.Println(err)
	}
	text2, err := printText("I am a good girl")
	if err != nil {
		fmt.Println(err)
	}

	combinedText := text1 + text2
	fmt.Println(combinedText)

	// fmt.Println(printText("I am a good boy"))
}
