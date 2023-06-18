package main

import "fmt"

type Quotes []string

// 'DisplayQuotes()' method will print all the quotes within the 'Quotes' slice
func (q Quotes) DisplayQuotes() {
	for _, quote := range q {
		fmt.Println(quote)
	}
}

func main() {
	spQuotes := Quotes{
		"I'm ready! I'm ready! 🧽",
		"I'm ugly and I'm proud! 👺",
		"Ravioli, ravioli, give me the formuoli! 🍝",
		"Krusty Krab pizza, it's the pizza for you and me! 🍕",
		"I heard in times of hardship, the pioneers would eat coral. 🪨",
	}

	// the selected option to call the 'DisplayQuotes()' method goes here!
	spQuotes.DisplayQuotes()
	Quotes.DisplayQuotes(spQuotes)
}
