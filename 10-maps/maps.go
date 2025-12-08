package main

import "fmt"

type members struct {
	age int
	role string
}

func main() {
	// declaring a map,
	fscMembers := map[string]members{
		"Olema":  {19, "Product Designer"},
		"Fisayo": {20, "Frontend Developer"},
		"Divine": {20, "Backend Developer"},
		"Nifemi": {20, "Project Manager"},
	}

	fmt.Println(fscMembers["Olema"].age)
}