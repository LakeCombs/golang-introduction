package main

import "fmt"

func main (){
	// s := make([]int, 3, 100)
	// fmt.Println(s)
	// fmt.Printf("length   %v\n", len(s))
	// fmt.Printf("Capacity %v\n", cap(s))

	// s := []int{}
	s := make([]int, 5, 100)
	fmt.Println(s)
	fmt.Printf("length   %v\n", len(s))
	fmt.Printf("Capacity %v\n", cap(s))

	s = append(s, 1)
	fmt.Println(s)

	fmt.Printf("length   %v\n", len(s))
	fmt.Printf("Capacity %v\n", cap(s))

	// s = append(s, 1 ,2,3,4,6,2,4,5)
	//using the spread operator
	s = append(s, []int{2,3,4,5}...)
	fmt.Println(s)

	 fmt.Printf("length   %v\n", len(s))
	fmt.Printf("Capacity %v\n", cap(s))
} 