package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Directory is not specified")
		return
	}
	root := os.Args[1]
	fmt.Println(root)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			fmt.Printf("%q\n", path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("walk is not completed")
		return
	}

}
