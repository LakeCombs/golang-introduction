package main

import "fmt"

func main() {
	arr1 := [...]int{1, 1, 1}
	//& is a pointer and i pass the location of array1 to array2
	arr2 := &arr1
	arr2[0] = 0
	fmt.Println(arr1)
	fmt.Println(arr2)
}