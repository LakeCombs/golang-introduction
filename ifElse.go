package main

import "fmt"

func main() {
	// if true {
	// 	fmt.Println("it is true")
	// }

	// if false {
	// 	fmt.Println("it is false")
	// }

	a := 30
	b := 38
	//Relational Operator ( < > <= >= ==)
	if a > b {
		fmt.Println("A is bigger then B")
	}
	if b < a {
		fmt.Println("B is bigger then A")
	}
	if a ==b {
		fmt.Println("Both are equal")
	}

	//logical operator example are (&& || !)

	p := 30
	q := 40
	r := 50

	if p > q && p > r {
		fmt.Println("p is biggest")
	}else if q > p && q > r {
		fmt.Println("q is the biggest")
	}else{
		fmt.Println(" r is the b iggest  ")
	}

}