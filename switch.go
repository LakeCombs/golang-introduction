package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 2
	// switch i {
	// 	case 1:fmt.Println("one")
	// case 2: 
	// 	fmt.Println("Two")
	// default:
	// 	fmt.Println("Default")
	// }

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a a weekend")
	default: 
	fmt.Println("It is not a weekend")
	// break 
	fmt.Println("Hello welome ")
	}

	//go already do the break statement for use so we don't need to add it

	t := time.Now()
	switch  {
		//go automatically refer to because we use it n the switch statement 
	case t.Hour() < 12:
		fmt.Println("Its before morning")
	case t.Hour() == 12:
		fallthrough
		//fallthrough is used to transfer control to the first statement of the case 
		//is present immediately after the case which has been executed
	default:
		fmt.Println("its after noon")
	}
	// fmt.Println(t)
}