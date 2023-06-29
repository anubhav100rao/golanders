package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{"/example/files/img/goland.svg", "example/files", "..files//img///", "", "///"}

	for _, path := range paths {
		fmt.Println("Base:", filepath.Base(path), "|", "Dir:", filepath.Dir(path))
	}
}
