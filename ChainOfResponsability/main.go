package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type password struct {
	plainPassword       string
	alreadyReversed     bool
	alreadyDuplicated   bool
	alreadyAddedSymbols bool
}

type PasswordGenerator interface {
	increasePasswordStrength(password)
	setNext(PasswordGenerator)
}

func main() {
	text := "batatinha"

	capitalize := &capitalize{}

	reverse := &reverse{}
	reverse.setNext(capitalize)

	duplicate := &duplicate{}
	duplicate.setNext(reverse)

	addSymbols := &addSymbols{}
	addSymbols.setNext(duplicate)

	password := &password{
		plainPassword: text,
	}

	addSymbols.increasePasswordStrength(*password)
}

type reverse struct {
	next PasswordGenerator
}

func (r *reverse) increasePasswordStrength(p password) {
	if p.alreadyReversed {
		r.next.increasePasswordStrength(p)
		return
	}

	p.plainPassword = reverseString(p.plainPassword)
	p.alreadyReversed = true
	fmt.Println(p.plainPassword)
	r.next.increasePasswordStrength(p)
}

func (r *reverse) setNext(next PasswordGenerator) {
	r.next = next
}

type duplicate struct {
	next PasswordGenerator
}

func (r *duplicate) increasePasswordStrength(p password) {
	if p.alreadyDuplicated {
		r.next.increasePasswordStrength(p)
		return
	}

	p.plainPassword = duplicateString(p.plainPassword)
	p.alreadyDuplicated = true
	fmt.Println(p.plainPassword)
	r.next.increasePasswordStrength(p)
}

func (r *duplicate) setNext(next PasswordGenerator) {
	r.next = next
}

type addSymbols struct {
	next PasswordGenerator
}

func (r *addSymbols) increasePasswordStrength(p password) {
	if p.alreadyAddedSymbols {
		r.next.increasePasswordStrength(p)
		return
	}

	p.plainPassword = addSymbol(p.plainPassword)
	p.alreadyAddedSymbols = true
	fmt.Println(p.plainPassword)
	r.next.increasePasswordStrength(p)
}

func (r *addSymbols) setNext(next PasswordGenerator) {
	r.next = next
}

type capitalize struct {
	next PasswordGenerator
}

func (r *capitalize) increasePasswordStrength(p password) {
	p.plainPassword = strings.Title(p.plainPassword)
	fmt.Println(p.plainPassword)
	fmt.Println("done")
}

func (r *capitalize) setNext(next PasswordGenerator) {
	r.next = next
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func duplicateString(s string) string {
	duplicated := s

	s = s + duplicated

	return s
}

func addSymbol(s string) string {
	rand.Seed(time.Now().UnixNano())
	byteString := []byte(s)
	symbols := []byte{'#', '@', '$', '%'}

	for i := 0; i < 3; i++ {
		position := rand.Intn(len(s)-1-0) + 0

		byteString[position] = symbols[position%4]
	}

	return string(byteString)
}
