package vehicles

import "fmt"

type Car struct {
	Brand        string
	LicensePlate string
	Color        string
	Speed        int
	Type         string
}

func (c *Car) Start() {
	fmt.Printf("Starting car.... max speed: %d", c.Speed)
}

func (c *Car) Stop() {
	fmt.Println("Stopping car....")
}

func SeeCarDetails(c Car) {
	fmt.Printf("Brand: %s\n\nColor: %s\n\nSpeed: %d\n\nLicense Plate: %s", c.Brand, c.Color, c.Speed, c.LicensePlate)
}
