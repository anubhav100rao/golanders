package main

import "fmt"

// getSumOfVectorElements() returns the sum of all elements in the vector
func getSumOfVectorElements(vector []int) int {
	var sum = 0

	for _, val := range vector {
		sum += val
	}
	return sum
}

// getProductOfVectorElements() returns the product of all elements in the vector
func getProductOfVectorElements(vector []int) int {
	var product = 1

	for _, val := range vector {
		product *= val
	}
	return product
}

// DO NOT delete or modify the contents within the main function!
func main() {
	var vc []int = []int{1, 2, 3, 4, 5, 6}

	var sum, prod int

	sum = getSumOfVectorElements(vc)
	prod = getProductOfVectorElements(vc)

	fmt.Println(sum, prod)
}
