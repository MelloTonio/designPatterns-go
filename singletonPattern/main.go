package main

import (
	"fmt"

	"github.com/singletonpattern/counter"
	first_instance "github.com/singletonpattern/firstInstance"
	second_instance "github.com/singletonpattern/secondInstance"
)

func main() {
	// Acessamos a mesma instância apesar das chamadas serem de pacotes diferentes
	first_instance.Increment()
	first_instance.Increment()
	first_instance.Increment()
	
	second_instance.Increment()
	second_instance.Increment()
	second_instance.Increment()

	c := counter.GetOrCreateInstance()

	fmt.Println(c.GetCurrentNumber())
}
