package main

const (
	hello  string = "Hello"
	hubble int    = 77
)

func main() {
	var (
		isEnable bool   = false
		name     string = "hello"
		id       int    = 100
	)

	const (
		Monday    = iota + 1 // Start from 1 instead of 0
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)

	const (
		Read = 1 << iota  // << A bit operation
		Write
		Execute
	)

	a, b := 1, 2
	println(a, b)

	println(Read, Write, Execute)
	println(isEnable, name, id)
}

