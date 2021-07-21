package main

import (
	"github.com/mellotonio/patternsgo/vehicles"
)

func buildCar() *vehicles.CarBuilder {
	newVehicle := vehicles.New()

	return &newVehicle
}

func main() {
	car := buildCar().
		SetBrand("HONDA").
		SetSpeed(50).
		SetColor("RED").
		SetLicensePlate("ADPT-2324").
		Get()

	vehicles.SeeCarDetails(*car)
}
