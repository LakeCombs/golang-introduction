package main

import "fmt"

type Vehicle struct {
	Name string
	Type string
}

type Car struct {
	Vehicle 
	Maxspeed float32
	Fueltype string
}

func main() {
	// c := Car {}
	// c.Name = "Ferrari"
	// c.Type = "Convertible"
	// c.Maxspeed  = 250
	// c.Fueltype = "Diesel"

	//if you want to use the literal syantax
	c := Car{
		Vehicle: Vehicle{Name: "Ferrari", Type:"Convetible" },
		Maxspeed: 250,
		Fueltype: "Diesel",

	}

	fmt.Println(c)
	fmt.Println(c.Name)
}