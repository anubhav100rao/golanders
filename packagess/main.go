package main

import (
	"fmt"
	"github.com/fatih/color"
)

//import (
//	_ "os"
//)
//
//import (
//	_ "database/sql"      // imports the Go database/sql package
//	_ "github.com/lib/pq" // imports the PostgreSQL sql drivers that the database/sql package requires
//)
//
//
//import (
//	. "fmt"
//	. "math"
//)

// executable packages
//The name of the package must be main.
//It should contain a function called main(), which will be the starting point of your program.

// installing external package
// go get github.com/fatih/color
// go mod tidy // command to add missing and remove unused modules in your project.
func main() {
	color.Red("This text should appear in red color")
	var n rune = '$'
	fmt.Print(n)
}
