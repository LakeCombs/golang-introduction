package main

import "fmt"

func main() {
	nums := [...]string{"one", "two", "three"}

	for _, num := range nums {
		fmt.Printf("%s \n ",  num)
	}

	for range nums {
		fmt.Printf("lakes")
	}
}