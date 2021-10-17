package main

import "fmt"

func main()  {
	//for indicating the number of element in the array
	// marks := [3]int{80, 90, 100}

	//for not indicating the number of element in the array
	marks := [...]int{80, 90, 100, 1000, 445, 3, 68}
	fmt.Printf("%v\n ", marks)
}