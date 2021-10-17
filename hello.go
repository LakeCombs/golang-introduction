package main

import "fmt"

func main()  {
	s := "hello world"
	b := []byte(s)

	fmt.Printf("%v, %T\n", b,b)
}