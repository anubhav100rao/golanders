package main

import "fmt"

func main() {
	var ranks = make(map[string]string)
	ranks["bronze"] = "third"
	ranks["silver"] = "second"
	ranks["gold"] = "first"

	cars := make(map[string]string)
	cars["bronze"] = "Ford"
	cars["silver"] = "Mercedes"
	cars["gold"] = "Ferrari"

	if val, ok := ranks["bronze"]; ok {
		println(val)
	}

	if val, ok := cars["bronze"]; ok {
		println(val)
	}

	for key, val := range ranks {
		println(key, val)
	}
	for key, val := range cars {
		println(key, val)
	}

	fmt.Println(ranks)
	fmt.Println(len(cars))
}
