package main

import "fmt"

type Animal struct {
	Name        string
	Class       string
	Emoji       string
	avgLifespan int
	Domestic    bool
}

func main() {
	crocodile := Animal{Name: "Crocodile", Class: "Reptile", Emoji: "🐊", avgLifespan: 55, Domestic: false}
	fmt.Printf("%#v\n", crocodile)

	elephant := new(Animal) // this will return pointer
	elephant.avgLifespan = 100
	fmt.Printf("%#v\n", elephant)

	var tiger Animal

	tiger.avgLifespan = 500
	tiger.Name = "bagh"
	tiger.Class = "cat"
	fmt.Printf("%#v\n", tiger)
}
