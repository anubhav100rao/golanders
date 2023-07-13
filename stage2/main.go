package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type FileGroup map[int64][]string

func (fileGroup *FileGroup) showListBySizes(sizes *Sizes) {
	for _, size := range *sizes {
		fmt.Printf("%d bytes\n", size)
		for _, path := range (*fileGroup)[size] {
			fmt.Println(path)
		}
		fmt.Println()
	}
}

type Sizes []int64

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
}
