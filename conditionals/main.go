package main

import "fmt"

func main() {

	var number int = 100 // Your number goes here.

	switch { // Equivalent to `switch true`
	case number%2 == 0:
		fmt.Println("The number", number, "is even")
	case number%2 != 0:
		fmt.Println("The number", number, "is odd")
	}

	letter := "a"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("The letter", letter, "is a vowel")
	default:
		fmt.Println("The letter", letter, "is a consonant")
	}

	fmt.Scanf("%d", &number)

	// Write your code here.
	switch {
	case number > 0:
		fmt.Println("Positive!")
	case number < 0:
		fmt.Println("Negative!")
	default:
		fmt.Println("Zero!")
	}

	
}
