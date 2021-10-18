package main

import (
	"fmt"
)

func main() {
	// countryPopulation := map[string]int{
	// 	"India":1234, 
	// 	"USA": 757584, 
	// 	"france":377840,
	// 	"canada":122334,
	// }

	countryPopulation := make(map[string]int, 10)

	countryPopulation = map[string]int{
		"India":1234, 
		"USA": 757584, 
		"france":377840,
		"canada":122334,
	}

	// fmt.Println(countryPopulation)
	fmt.Println(countryPopulation["India"])
	countryPopulation["England"] = 77485
	fmt.Println(countryPopulation)
	delete(countryPopulation, "canada")

	//if a kay is not present in a map the result wil be zero

	fmt.Println(countryPopulation["russaia"])

	//we can check if a value is in a map with the temp, ok function
	// temp, ok := countryPopulation["USA"]
	// fmt.Println(temp, ok)

	// if we wawnt to only check if it is available we use _
	_ , ok := countryPopulation["USA"]
	fmt.Println(ok)

	//to print the lengnt of the map
	fmt.Println(len(countryPopulation))

	cp := countryPopulation
	delete(countryPopulation, "India")
	fmt.Println(cp)
	//the resutl is the same as printing the countryPopulation

	fmt.Println(countryPopulation)

}