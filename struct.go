package main

import "fmt"

type Celebrity struct {
	age int
	CelebrityName string
	movies []string
	coActors []string
}


func main() {
	//struct is the only type in go that allow to put different type of data types in one place
	aCelebrity := Celebrity{		
	//you can define a struct without using a key value pair
		 56, 
		"Lake Combs",
		 []string{
		},
		[]string{
			"gal gadot", 
			"paul walker", 
			"vin diesel",
		},
	}
	fmt.Println(aCelebrity)

	//go get values from the struct
	fmt.Println(aCelebrity.age)
	fmt.Println(aCelebrity.coActors[1])

}