package main

import "fmt"

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Toyota",
		carModel: "Camry",
		engine: gasEngine{
			gallons: 15.0,
			mpg:     30.0,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   75.0,
			mpkwh: 4.0,
		},
	}
	fmt.Println(gasCar)
	fmt.Println(electricCar)
}

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}
