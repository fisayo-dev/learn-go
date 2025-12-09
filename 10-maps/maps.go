package main

import (
	"fmt"
)

type members struct {
	name string
	age int
	role string
}

func (m members) printStmt() string{
	return fmt.Sprintf("%v is %v years old and is a %v",m.name, m.age, m.role)
}

func getNamesCounts(names []string) map[string]map[string]int {
	counts :=  make(map[string]map[string]int)
	for _, name := range names {
		// Safety checking for empty names
		if name == "" {
			continue
		}
		firstChar := string(name[0])
		_,ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		} 
		counts[firstChar][name]++
	}
	return counts
}

func main() {

	// Create an empty map:
	// userMap := make(map[string]int)
	
	// declaring a map,
	// fscMembers := map[string]members{
	// 	"Olema":  {"Olema", 19, "Product Designer"},
	// 	"Fisayo": {"Fisayo", 20, "Frontend Developer"},
	// 	"Divine": {"Divine", 20, "Backend Developer"},
	// 	"Nifemi": {"Nifemi", 20, "Project Manager"},
	// }


	// fmt.Println(fscMembers["Olema"].age)
	// fmt.Println(fscMembers["Olema"].printStmt())

	countNames := getNamesCounts([]string{"Funmit", "Funke", "Funke", "Fisayo", "Mary", "Mary", "Peter", "Parker"})
	fmt.Println(countNames["F"])
}