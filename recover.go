package main

import (
	"fmt"
	"runtime/debug"
)

func recoverAccount(){
	if r := recover(); r !=  nil{
		fmt.Println("recovered from \n  " ,  r )
		debug.PrintStack()
	}
}

func accountDetails(userID *int, branchCode *int) {
	defer recoverAccount()
	defer fmt.Println("deffering call in account details")

	if userID == nil{
		panic("runtime error: user id can't be nil")
	}

	if branchCode == nil {
		panic("runtime error : branch code can't be nil")
	}
	fmt.Printf("%d %d\n", *userID, *branchCode)
	fmt.Println("returned normally for accountDetails funcion")
}

func main() {
	defer fmt.Println("defered call in main function")
	userID := 4567

	accountDetails(&userID, nil)
	fmt.Println("returned normally from main function  ")
}