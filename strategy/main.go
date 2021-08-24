package main

import "fmt"

type billingStrategy interface {
	getPrice(b float32) float32
}

func main() {
	normalStrategy := &normalStrategy{}
	happyHourStrategy := &happyHourStrategy{}

	costumer := &Costumer{}
	costumer.setStrategy(happyHourStrategy)
	costumer.add(20.0, 1)

	costumer.setStrategy(normalStrategy)
	costumer.add(3, 1)

	fmt.Println(costumer.total)
}

type happyHourStrategy struct {
}

func (h *happyHourStrategy) getPrice(price float32) float32 {
	return price * 0.7
}

type normalStrategy struct {
}

func (h *normalStrategy) getPrice(price float32) float32 {
	return price
}

type Costumer struct {
	billingStrategy billingStrategy
	total           float32
}

func (c *Costumer) add(price float32, quantity int) {
	c.total += c.billingStrategy.getPrice(price * float32(quantity))
}

func (c *Costumer) setStrategy(b billingStrategy) {
	c.billingStrategy = b
}
