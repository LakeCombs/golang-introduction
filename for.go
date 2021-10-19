package main

import "fmt"

func main (){
	nums := []string{"one", "two", "three",}
	
	for idx, num := range nums {
		fmt.Printf("%d: %s\n", idx, num)
	}

	for idx := range nums {
		fmt.Printf("%d \n", idx)
	}

	for _,num := range nums {
		fmt.Printf("%s \n", num)
		
	}
}