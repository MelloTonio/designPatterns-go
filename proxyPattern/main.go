package main

import "fmt"

type Handler interface {
	request()
}

type ConnectionHandler struct {
}

func (rs ConnectionHandler) request() {
	fmt.Println("Real subject: handling request")
}

type Proxy struct {
	connection ConnectionHandler
}

func (p Proxy) request() {
	if p.accessGranted() {
		p.log()
		p.connection.request()
	}

}

func (p Proxy) accessGranted() bool {
	fmt.Println("Proxy: Access Granted")
	return true
}

func (rs Proxy) log() {
	fmt.Println("Proxy: Accessing content")
}

func executeRequest(h Handler) {
	h.request()
}

func main() {

	ch := ConnectionHandler{}
	p := Proxy{connection: ch}
	executeRequest(p)
}
