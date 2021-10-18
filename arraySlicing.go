package main

import "fmt"

func main() {
	// arr1 := []int{1, 2, 3}

	// fmt.Println(arr1)
	// fmt.Println(len(arr1))
	// fmt.Println("Capacity %v ", cap(arr1))

	a := [...]int{1,2,3,4,5,6,7,8,9,10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	a[3] = 99
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)


}