package main

import "fmt"

func recoverSlicePanic()  {
	if r := recover(); r != nil {
		fmt.Println("recover from ", r)
	}
}

func slicePanic()  {
	defer recoverSlicePanic()
	n := []int{5,7,4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a ")
	
}

func main(){
	slicePanic()
	fmt.Println("normally returned from main") 
}