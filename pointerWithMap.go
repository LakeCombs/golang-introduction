package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3 }

	for k := range m {
	fmt.Printf("keys: %q \n " , k)
	fmt.Printf("keys: %s\n " , k)
	//%q use double quote while %S uses the single quote
	}

	for k , v := range m {
		fmt.Printf("keys: %q, value : %v \n", k, v)
	}
}