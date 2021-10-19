package main

import "fmt"
 
func accountDetails(userID *int, branchCode *int){

	if userID == nil {
		panic("runtime error: user id can be nil ")
	}

	if branchCode == nil {
		panic("runtime error : branchCode can't be nil")
	}
	fmt.Printf("%d %d\n", *userID, *branchCode)
	fmt.Println("returned normally from acountDetails function")
}

func main (){
	defer fmt.Println("deffered called in the main function")
	userID := 4567
	// branchCode := 12
	accountDetails(&userID, nil)
	fmt.Println("returned normally from main function")

	//if i replace branchCode with nil the panic error will return

} 