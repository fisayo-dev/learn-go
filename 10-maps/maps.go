package main

import "fmt"

type members struct {
	age int
	role string
}

func (m members) printStmt() string{
	return fmt.Sprintf("is %v years old and is a %v", m.age, m.role)
}

func main() {
	// declaring a map,
	fscMembers := map[string]members{
		"Olema":  {19, "Product Designer"},
		"Fisayo": {20, "Frontend Developer"},
		"Divine": {20, "Backend Developer"},
		"Nifemi": {20, "Project Manager"},
	}

	// fmt.Println(fscMembers["Olema"].age)
	fmt.Println(fscMembers["Olema"].printStmt())
}