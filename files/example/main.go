package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	var info = func() {
		fileInfo, err := os.Stat("info.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File name:", fileInfo.Name())        // File name: info.txt
		fmt.Println("Size:", fileInfo.Size(), "bytes")    // Size: 35 bytes
		fmt.Println("Permission mode:", fileInfo.Mode())  // Permission mode: -rw-r--r--
		fmt.Println("Last modified:", fileInfo.ModTime()) // Last modified: 2022-03-10 11:15:22 -0500 EST
		fmt.Println("Is directory:", fileInfo.IsDir())    // Is directory: false

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if ok {
			fmt.Println("User identifier:", stat.Uid)  //  User identifier: 1000
			fmt.Println("Group identifier:", stat.Gid) // Group identifier: 1000
		}
	}
	info()

	var files = func() {
		fileInfo, err := os.Stat("files")
		if err != nil {
			log.Fatal(err)
		}
		mode := fileInfo.Mode()                           // assign the FileMode value returned by the Mode() method to `mode`.
		fmt.Printf("File perm. bits: %#o\n", mode.Perm()) // File perm. bits: 0775
		fmt.Println("File type bits:", mode.Type())       // File type bits: d---------
		fmt.Println("Is regular:", mode.IsRegular())      // Is regular: false
	}
	files()

	var fileExists = func() {
		fileName := "impostor.png"
		fileInfo, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			log.Fatal("The file ", fileName, " does not exist!")
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fileInfo.Name(), "exists!")
	}
	fileExists()
}

// Output:
// Linux/macOS -> ~/GolandProjects/example
// Windows -> C:\GolandProjects\example
