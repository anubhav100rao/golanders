package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var fileName string
	fmt.Scanln(&fileName)

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	mode := fileInfo.Mode()
	fmt.Printf("File permission bits: %#o\n", mode.Perm())
	fmt.Println("Is directory:", mode.IsDir())
}
