package main

import "fmt"

// func equilTriangle(base, height float64) (float64, float64){
// 	var area  = 0.5 * base * height
// 	var perimter =  3 * base
// 	return area, perimter
//  }

func equilTriangle(base, height float64) ( area ,  perimeter  float64){
	area  = 0.5 * base * height
	perimeter =  3 * base 
	return 
 }



func main()  {
	base, height := 10.5, 4.6
	area, perimeter := equilTriangle(base, height)
	fmt.Printf("Area is %f\n  Perimeter is %f\n", area, perimeter)
	fmt.Println("The area and perimeter of the triangle are : ", area, perimeter)
}