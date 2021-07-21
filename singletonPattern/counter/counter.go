package counter

// Esse padrão visa a criação de uma instância única de uma classe
// fazendo com que o estado dela seja refletido (e alterado) igualmente para todos
var c *counter

func GetOrCreateInstance() *counter {
	if c == nil {
		c = &counter{}
	}

	return c
}

type counter struct {
	currentNumber int
}

func (c *counter) IncrementCounter() {
	c.currentNumber++
}

func (c *counter) GetCurrentNumber() int {
	return c.currentNumber
}
