package main

import "fmt"

func main () {
	totalAnalysisScore := 65
	goodAnlysisScoreMark := 70
	
	// If - else stmt

	var msg string // General
	if totalAnalysisScore >= goodAnlysisScoreMark {
		msg = "Analysis score is good"
	} else if totalAnalysisScore < goodAnlysisScoreMark {
		msg = "Analysis score is not good!"
	}

	fmt.Println(msg)
}