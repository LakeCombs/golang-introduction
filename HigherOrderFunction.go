package main

import "fmt"

// func simple (a func(a, b int) int){
// 	fmt.Println(a(50, 5))
// }

// func main(){
// 	f := func(a , b int) int{
// 		return a + b
// 	}

// 	simple(f)
// }



func simple()  func(a, b int) int {
	f := func(a, b int) int{
		return a + b
	}
	return f
}

func main()  {
	s := simple()
	fmt.Println(s(50, 5))
}