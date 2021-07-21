package main

import (
	"fmt"

	"github.com/mellotonio/factory/factory"
)

func main() {
	jsonFile := factory.ReadFileFactory("JSON")
	xmlFile := factory.ReadFileFactory("XML")

	fmt.Println(jsonFile.Get("Elliot"))
	fmt.Println(xmlFile.Get("Elliot"))

	user1 := factory.User{
		Name: "XYZ",
		Social: factory.Social{
			Facebook: "facebook.com/909090",
		},
	}

	xmlFile.Insert(user1)

	jsonFile.Insert(user1)
}
