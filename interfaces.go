package main

import "fmt"

type generic interface{
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{ 
	width , height float64
}

func (c circle) area () float64 {
	return 3.143 * c.radius * c.radius
}

func (r rectangle) area () float64 {
	return r.height * r.width
}

func calculate( g generic) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main (){
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	calculate(r)
	calculate(c)
}