package main

import "fmt"

func main() {
	a := "First variable"
	b := "Second variable"
	fmt.Printf("%[1]s | %[1]s \n", a) // First variable | First variable
	fmt.Printf("%[2]s | %[1]s", a, b)
}
