package main

import "fmt"

func main() {
	// map is a data structure that map key to there value
	// countryPopulation := map[string]int{
	// 	"brazil": 22345, 
	// 	"USA":4455,
	// 	"france":334499,
		
	// }

	//use cannot use a slice as data type for a map 

	//to  make a map with a make function

		countryPopulation := make(map[string]int, 10)

		countryPopulation = map[string]int{
		"brazil": 22345, 
		"USA":4455,
		"france":334499,
		}
	
	//to add value to a map
	countryPopulation["Russia"] = 34534


	//to remove value from a map
	delete(countryPopulation, "USA")

	//if an element do not exist in a map the output is 0
	fmt.Println(countryPopulation["USA"])

	// temp, ok := countryPopulation["brazil"]
	// fmt.Println(temp, ok)

	_, ok := countryPopulation["france"] 
	fmt.Println(ok)
	// fmt.Println(countryPopulation)
	fmt.Println(len(countryPopulation))

	cp := countryPopulation
	delete(countryPopulation, "france")
	fmt.Println(cp)
	fmt.Println(countryPopulation)

	//maps are reference type so any change in it child type will lead to a change in the base type

}