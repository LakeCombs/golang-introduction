package main

import "fmt"

func main(){
	var snames[3] string
	fmt.Printf("%v\n" , snames)

	//to populate the array   
	snames[0] = "lakes"
	snames[1] = "combs"
	snames[2] = "chi"
	//in an array that contain string you cannot add another variable type
	fmt.Printf("%v \n ", snames[0])
	fmt.Println(len(snames))
}