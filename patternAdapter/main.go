package main

import (
	"fmt"
	"math"
)

type circularForm interface {
	getRadius() float64
}

type circle struct {
	Radius float64
}

type square struct {
	Width float64
}

type circleHole struct {
	Radius float64
}

type squareAdapter struct {
	Square square
}

func (c circle) getRadius() float64 {
	return c.Radius
}

func (sa squareAdapter) getRadius() float64 {
	result := math.Sqrt(math.Pow((sa.Square.Width/2), 2) * 2)

	return result
}

func (c circleHole) fits(h circularForm) bool {
	result := c.Radius >= h.getRadius()

	return result
}

func main() {
	var CircularHole circleHole = circleHole{Radius: 5}
	var Circle circle = circle{5}

	if CircularHole.fits(Circle) {
		fmt.Println("Yes, it fits")
	}

	// O padrão adapter permite "encaixar" objetos com "formatos" diferentes
	// fazemos uma tradução de um objeto para o outro por meio da utilização de interfaces
	var Square square = square{Width: 3}
	var SquareAdapter squareAdapter = squareAdapter{Square: Square}

	if CircularHole.fits(SquareAdapter) {
		fmt.Println("Yes, it fits")
	}
}
