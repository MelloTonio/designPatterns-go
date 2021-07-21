package factory

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
)

type (
	User struct {
		Name   string `xml:"name" json:"name"`
		Social Social `xml:"social" json:"social"`
	}

	Social struct {
		Facebook string `xml:"facebook" json:"facebook"`
	}

	XML struct {
		All []User `xml:"user" json:"user"`
	}

	JSON struct {
		All []User `xml:"user" json:"users"`
	}

	Tag struct {
		Key   string `xml:"name,attr"`
		Value string `xml:",chardata"`
	}

	// Esse padrão permite que a partir de uma só interface possamos criar
	// objetos com implementações diferentes
	File interface {
		Insert(User)
		GetSocialMedia(string) (string, error)
	}
)

func (x XML) GetSocialMedia(key string) (string, error) {
	parsedUsers, err := x.getAllInfo()
	if err != nil {
		return "", err
	}

	for _, user := range parsedUsers.All {
		if user.Name == key {
			return user.Social.Facebook, nil
		}
	}

	return "", errors.New("User not found")
}

func (jsonFile JSON) GetSocialMedia(key string) (string, error) {
	users, err := jsonFile.getAllInfo()
	if err != nil {
		return "", err
	}

	for _, user := range users.All {
		if user.Name == key {
			return user.Social.Facebook, nil
		}
	}

	return "", errors.New("User not found")
}

func (jsonFile JSON) Insert(newUser User) {
	users, err := jsonFile.getAllInfo()
	if err != nil {
		return
	}

	users.All = append(users.All, newUser)

	marshaledUser, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./factory/input.json", marshaledUser, 0644)
	if err != nil {
		panic(err)
	}
}

func (j JSON) getAllInfo() (*JSON, error) {
	jsonFile, err := os.Open("./factory/input.json")

	if err != nil {
		return nil, err
	}

	parsedUsers := new(JSON)

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &parsedUsers)
	if err != nil {
		return nil, err
	}

	return parsedUsers, nil
}

func (xmlFile XML) Insert(newUser User) {
	users, err := xmlFile.getAllInfo()
	if err != nil {
		panic(err)
	}

	users.All = append(users.All, newUser)

	out, err := xml.MarshalIndent(users, "", "   ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./factory/input.xml", out, 0644)
	if err != nil {
		panic(err)
	}

}

func (x XML) getAllInfo() (*XML, error) {
	xmlFile, err := os.Open("./factory/input.xml")
	if err != nil {
		return nil, err
	}

	parsedUsers := new(XML)

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(byteValue, &parsedUsers)
	if err != nil {
		return nil, err
	}

	return parsedUsers, nil
}

func ReadFileFactory(fileType string) File {
	switch fileType {
	case "XML":
		return XML{}

	case "JSON":
		return JSON{}

	default:
		return nil
	}

}
