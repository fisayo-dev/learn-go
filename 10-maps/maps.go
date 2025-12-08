package main

import "fmt"

type members struct {
	name string
	age int
	role string
}

func (m members) printStmt() string{
	return fmt.Sprintf("%v is %v years old and is a %v",m.name, m.age, m.role)
}

func main() {

	// Create an empty map:
	// userMap := make(map[string]int)
	
	// declaring a map,
	fscMembers := map[string]members{
		"Olema":  {"Olema", 19, "Product Designer"},
		"Fisayo": {"Fisayo", 20, "Frontend Developer"},
		"Divine": {"Divine", 20, "Backend Developer"},
		"Nifemi": {"Nifemi", 20, "Project Manager"},
	}

	// fmt.Println(fscMembers["Olema"].age)
	fmt.Println(fscMembers["Olema"].printStmt())
}