package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5}
	//to remove the first element
	r := s[1:]
	fmt.Println(s)	
	fmt.Println(r)	
	//to remove element from the back
	p := s[:len(s)-1]
	fmt.Println(p)

	//to remove element from the middle
	a := append(s[:2], s[3:]...)
	fmt.Println(a)
	
}