package main

import "fmt"

// func main()  {
// 	first := func ()  {
// 		fmt.Println("inside the anonymout function")
// 	}
// 	first()
// 	fmt.Printf("%T", first )
// }

// func main()  {
// 	func( str string){
// 		fmt.Println("Inside the Anonymouns function", str)
// 	}("Hello There")
// }


type add func(a int, b int) int

func main()  {
	var a add = func (a, b int) int {
		return a + b 
	}

	s := a (50, 5)
	fmt.Println("sum ", s)
}