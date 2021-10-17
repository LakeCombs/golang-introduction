package main

import "fmt"

// const (
// 	_ = iota
// 	//_ is write only variable
// 	mbaHOD
// 	cscHOD
// 	iseHOD
// //iota is scoped to constant block
// )

const (
	_ = iota + 3
	mbaHOD
	cscHOD
	iseHOD 
)


func main()  {
// 	var hodType int = mbaHOD
// 	fmt.Printf("%v\n", hodType == mbaHOD)
// 	fmt.Printf("%v\n", hodType == cscHOD )
	fmt.Printf("%v\n", mbaHOD)
}