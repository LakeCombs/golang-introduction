package main

import "fmt"

func main()   {
	//go only have one loop i.e for loop
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	//to createa for loop like a while loop 
	// sum := 1
	// for  sum < 100{ 
	// 	sum += sum
	// }
	
	// fmt.Println(sum)


	//to make an infinite loop
	// sum := 1
	// for {
	// 	fmt.Println(sum)
	// }

	//to break a finite loop
	// sum := 1
	// for{
	// 	sum += sum
	// 	fmt.Println(sum) 
	// 	if sum == 8 {
	// 		break
	// 	}
	// }

	//to use the continue statement
	sum := 1
	for{
		sum += sum
		// fmt.Println(sum)
		if(sum == 8 ){
			continue
		}
		fmt.Println(sum)
		if(sum == 32){
			break
		}
	}
}