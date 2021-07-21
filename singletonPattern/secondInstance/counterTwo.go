package secondinstance

import "github.com/singletonpattern/counter"

func Increment() {
	c := counter.GetOrCreateInstance()
	c.IncrementCounter()
}
