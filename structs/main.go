package main

import "fmt"

type Animal struct {
	Name        string
	Class       string
	Emoji       string
	avgLifespan int
	Domestic    bool
}

type Person struct {
	Name string
	Age  int
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

	rick := Person{
		Name: "Rick",
		Age:  71,
	}
	fmt.Printf("%#v\n", rick)

	morty := new(Person)
	morty = &Person{
		Name: "Morty",
		Age:  15,
	}
	fmt.Printf("%#v\n", morty)

	var summer = Person{
		Name: "Summer",
		Age:  18,
	}

	fmt.Printf("%#v\n", summer)
}
