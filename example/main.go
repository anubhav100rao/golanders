package main

import (
	"fmt"
	"log"
	"os"
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

	fmt.Println(filepath.IsAbs("/home/User/GolandProjects/example")) // true
	fmt.Println(filepath.IsAbs("/.bashrc"))                          // true
	fmt.Println(filepath.IsAbs(".bashrc"))                           // false
	fmt.Println(filepath.IsAbs("/"))                                 // true Linux/macOS | false Windows
	fmt.Println(filepath.IsAbs(""))                                  // false

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			fmt.Println("Directory:", path, "size:", info.Size(), "bytes")
		} else {
			fmt.Println("File:", path, "size:", info.Size(), "bytes")
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
