package main

import "fmt"

func main()  {
	aCelebrity := struct{age int}{age: 55}
	bCelebrity := &aCelebrity
	bCelebrity.age = 95
	fmt.Println(aCelebrity)
	fmt.Println(bCelebrity)

} 