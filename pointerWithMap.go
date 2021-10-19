package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3 }

	for k := range m {
fmt.Printf("keys: %q \n " , k)
}
}