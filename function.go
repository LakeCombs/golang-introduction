package main

import "fmt"


func calcAmount(price , qty int) int  {
	var totalAmount = price * qty
	return totalAmount
}

func main()  {
	price , qty := 20, 5
	totalAmount := calcAmount(price, qty)
	fmt.Println("the total amount is: ", totalAmount)
}

