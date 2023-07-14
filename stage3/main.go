package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type FileGroup map[int64][]string
type Sizes []int64

func (fileGroup *FileGroup) showListBySizes(sizes *Sizes) {
	for _, size := range *sizes {
		fmt.Printf("%d bytes\n", size)
		for _, path := range (*fileGroup)[size] {
			fmt.Println(path)
		}
		fmt.Println()
	}
}

func (sizes *Sizes) sort(order int) {
	sort.Slice(*sizes, func(i, j int) bool {
		if order == 1 {
			return (*sizes)[i] > (*sizes)[j]
		}

		return (*sizes)[i] < (*sizes)[j]
	})
}

func walkDir(path string, extension string) (*FileGroup, *Sizes) {
	var fileGroup FileGroup = make(map[int64][]string)
	var sizes Sizes = make([]int64, 0)

	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() && (extension == "" || (extension != "" && filepath.Ext(path) == extension)) {
			_, ok := fileGroup[info.Size()]

			if !ok {
				sizes = append(sizes, info.Size())
				fileGroup[info.Size()] = make([]string, 0)
			}

			fileGroup[info.Size()] = append(fileGroup[info.Size()], path)
		}

		return nil
	})

	return &fileGroup, &sizes
}

func setFileFormatValue(fileFormat *string) {
	fmt.Scanln(fileFormat)

	if *fileFormat != "" {
		*fileFormat = "." + *fileFormat
	}
}

func setSortingValue(sorting *int) {
	fmt.Println("\nEnter a sorting option")
	fmt.Scanln(sorting)

	if *sorting != 1 && *sorting != 2 {
		fmt.Println("\nWrong option")
		setSortingValue(sorting)
	}
}

func getMd5Hash(path string) string {

	hashFunc := md5.New()
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer file.Close()

	if _, err := io.Copy(hashFunc, file); err != nil {
		log.Fatal(err)
		return ""
	}

	return fmt.Sprintf("%x", hashFunc.Sum(nil))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
		return
	}

	var fileFormat string
	var sorting int

	fmt.Println("Enter file format")
	setFileFormatValue(&fileFormat)

	fmt.Println("\nSize sorting option")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")

	setSortingValue(&sorting)
	fmt.Println()

	fileGroup, sizes := walkDir(os.Args[1], fileFormat)
	sizes.sort(sorting)
	fileGroup.showListBySizes(sizes)

	fmt.Println("Check duplicates?")
	var response string
	fmt.Scanln(&response)

	for response != "yes" && response != "no" {
		fmt.Println("Wrong option")
		fmt.Println("Check duplicates?")
		fmt.Scanln(&response)
	}

	if response == "no" {
		return
	}

	var counter = 1

	for _, size := range *sizes {
		hashes := make(map[string][]string)
		found := false
		for _, path := range (*fileGroup)[size] {
			hash := getMd5Hash(path)
			_, ok := hashes[hash]

			if !ok {
				hashes[hash] = make([]string, 0)
			} else {
				found = true
			}

			hashes[hash] = append(hashes[hash], path)
		}

		if found {
			fmt.Printf("%d bytes\n", size)
			for _, paths := range hashes {
				if len(paths) > 1 {
					fmt.Printf("Hash: %s\n", getMd5Hash(paths[0]))
					for _, path := range paths {
						fmt.Printf("%d. %s\n", counter, path)
						counter++
					}
					fmt.Println()
				}
			}
		}
	}
}
