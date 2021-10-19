package main

import "fmt"

func main() {
	a := 150
	// b := 40
	// c := 50

	if a > 50 || a < 100 {
		fmt.Println("A lies in Between 50 and 100")
	}

	if !false {
		fmt.Println("True")
	}

	if !true {
		fmt.Println("True")
		//this code will not run because !true is false  
	}

	// if err != nil {
	// 	// handle the error
	// }

	if num := 100; num > 0{
		fmt.Println("A is a positive ")
	}

	
	if num := 100; num % 2 == 0{
		fmt.Println("num is Even")
	}
}