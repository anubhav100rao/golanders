package main

import "fmt"

/*
Since arrays always have a fixed number of elements, the Go compiler must know the amount of memory to allocate for an array at compile-time; therefore, constant values, which are created at compile-time, can be used as the array size. However, the value of a variable var is unknown until run time; therefore, you're not allowed to use a variable as the array size.
*/
func main() {
	var arr [5]int
	fmt.Println(arr)

	var a = [4]int{10, 2}
	fmt.Println(a) // Will print [10 2 0 0]

	var b = [4]int{0: 10, 3: 2}
	fmt.Println(b) // Will print [10 0 0 2]

	var c = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(c) // [[1 2 3 4] [5 6 7 8] [9 10 11 12]]

	var d = [3][4]int{
		1: {
			2: 5,
			3: 6,
		},
	}
	fmt.Println(d) // [[0 0 0 0] [0 0 5 6] [0 0 0 0]]

	// aa := [4]int{} this is ok
	// aa := [4]int this is not ok

	array := [4][4][4]float32{
		1: {
			0: {
				2: 88.6,
			},
		},
	}

	// DO NOT delete, this line prints the array!
	fmt.Println(array)
}
