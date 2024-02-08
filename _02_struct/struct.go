package main

import (
	"fmt"
)

type wheel struct {
	numberOfWheels int
	buildMaterial  string
}

//nested struct

type car struct {
	model      string
	wheelModel wheel
}

//embeding struct

type bike struct {
	model string
	wheel
}

func displayBike(vehicle bike) {
	fmt.Println(vehicle)
}

func displayCar(vehicle car) {
	fmt.Println(vehicle)
}

func main() {
	// ceat := wheel{
	// 	numberOfWheels: 2,
	// 	buildMaterial:  "Rubber",
	// }
	honda := bike{
		model: "bullet",
		wheel: wheel{
			numberOfWheels: 2,
			buildMaterial:  "Rubber",
		},
	}
	displayBike(honda)
	fmt.Println(honda.numberOfWheels)
}
