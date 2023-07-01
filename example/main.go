package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	paths := []string{"/example/files/img/goland.svg", "example/files", "..files//img///", "", "///"}

	for _, path := range paths {
		fmt.Println("Base:", filepath.Base(path), "|", "Dir:", filepath.Dir(path))
	}

	fmt.Println(filepath.Join("example", "files", "img"))
	fmt.Println(filepath.Join("example", "", "files/img"))

	fmt.Println(filepath.Join("home/User/example", "../../../../files"))
	fmt.Println(filepath.Join("", "")) // returns an empty string!

	for _, path := range paths {
		abs, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(abs)
	}
}
