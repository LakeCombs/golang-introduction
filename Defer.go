package main

import "fmt"

func main (){
	//defer make a func to perform later until a surrounding code runs first
	defer fmt.Println("Hello")
	fmt.Println("combs")
	fmt.Println("lakes")
}

//if you have more then one defer statement 
//it behave like a last in first out s