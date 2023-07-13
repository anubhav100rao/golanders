package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

	fileInfo, err := os.Stat("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())        // File name: info.txt
	fmt.Println("Size:", fileInfo.Size(), "bytes")    // Size: 35 bytes
	fmt.Println("Permission mode:", fileInfo.Mode())  // Permission mode: -rw-r--r--
	fmt.Println("Last modified:", fileInfo.ModTime()) // Last modified: 2022-03-10 11:15:22 -0500 EST
	fmt.Println("Is directory:", fileInfo.IsDir())
}

























