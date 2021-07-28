package main

import "fmt"

type Calculator interface {
	Calculate() int
}

type Number struct {
	Value int
}

func (p Number) Calculate() int {
	return p.Value
}

type Add struct {
	calculator Calculator
}

func (p Add) Calculate() int {
	newNumber := p.calculator.Calculate()

	return newNumber + 10
}

type Multiply struct {
	calculator Calculator
}

func (p Multiply) Calculate() int {
	newNumber := p.calculator.Calculate()

	return newNumber * 10
}

func main() {

	newNumber := Number{Value: 5}

	// Com o uso do padrão decorator conseguimos "empilhar" objetos que implementam a mesma interface
	sumNumber := Add{
		calculator: newNumber,
	}

	multiplyNumber := Multiply{
		calculator: sumNumber,
	}

	sumNumber = Add{
		calculator: multiplyNumber,
	}

	sumNumber = Add{
		calculator: sumNumber,
	}

	fmt.Println(sumNumber.Calculate())
}
