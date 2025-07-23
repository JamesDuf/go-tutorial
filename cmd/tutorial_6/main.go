package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8 // miles per gallon
	gallons uint8 // gallons of gas
	owner         // embedded struct for owner information
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string // owner's name
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons // calculate miles left based on mpg and gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh // calculate miles left based on mpg and gallons
}

func canMakeIt(e engine, distance uint8) {
	if distance <= e.milesLeft() {
		fmt.Println("The engine can make it.")
	} else {
		fmt.Println("The engine cannot make it.")
	}
}

type engine interface {
	milesLeft() uint8 // method to calculate miles left signature
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}} // create an instance of gasEngine with initial values
	//myEngine.mpg = 25     // set miles per gallon
	//myEngine.gallons = 10 // set gallons of gas
	fmt.Println("This engine is owned by:", myEngine.owner.name) // access embedded struct field
	fmt.Println("Miles per gallon:", myEngine.mpg)               // access field of gasEngine
	fmt.Println("Gallons of gas:", myEngine.gallons)             // access field of
	// calculate distance that can be traveled
	distance := myEngine.mpg * myEngine.gallons
	fmt.Println("Distance that can be traveled:", distance, "miles")

	// Anonymous struct example
	// Anonymous structs are useful for quick, one-off data structures without defining a type
	// They can be used for temporary data storage or when you don't need a full struct definition
	// They are not reusable like named structs, but they can be convenient for simple tasks
	var myEngine2 = struct {
		mpg     uint8 // miles per gallon
		gallons uint8 // gallons of gas
	}{25, 15} // create an anonymous struct with initial values
	fmt.Println("Anonymous struct - Miles per gallon:", myEngine2.mpg) // access field
	fmt.Println("Anonymous struct - Gallons of gas:", myEngine2.gallons)
	fmt.Println("Miles left in the gas engine: ", myEngine.milesLeft()) // call method to calculate miles left

	canMakeIt(myEngine, 50) // check if the engine can make it 300 miles
	//canMakeIt(myEngine2, 50) // will not compile, as myEngine2 does not implement the engine interface
}
