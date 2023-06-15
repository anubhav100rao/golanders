package main

import "fmt"

func main() {
	var tt int
	fmt.Scan(&tt)

	for ; tt > 0; tt-- {
		var n int
		fmt.Scan(&n)
		if n < 5 {
			fmt.Println("Bob")
		} else {
			fmt.Println("Alice")
		}
	}
}	
