package main

import "fmt"

func main() {
	nums := &[...]string{"one", "two", "three"}


	//& is a pointer while * is a reference to the pointer
	for idx, num := range *nums {
		fmt.Printf("%d :  %s\n ", idx , num)
	}

	for idx, num := range nums {
		fmt.Printf("%d : %s \n", idx, num)
	}

}