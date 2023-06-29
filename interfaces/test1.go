package main

import (
	"bufio"
	"fmt"
	"os"
)

type Text string

// Create the `TabFormat()` method below:
func (t Text) TabFormat() string {
	return "\t" + string(t)
}

// Create the `DoubleQuoteFormat()` method below:
func (t Text) DoubleQuoteFormat() string {
	return "\"" + string(t) + "\""
}

// Create the `SingleQuoteFormat()` method below:
func (t Text) SingleQuoteFormat() string {
	return "'" + string(t) + "'"
}

// Create the `Formatter` interface that implements the
// `TabFormat()`, `DoubleQuoteFormat()` and `SingleQuoteFormat()` methods below:
type Formatter interface {
	TabFormat() string
    DoubleQuoteFormat() string
    SingleQuoteFormat() string
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	someText := Text(scanner.Text())

	textFormatter := someText

	formattedText := textFormatter.SingleQuoteFormat()
	fmt.Println(formattedText)

	formattedText = textFormatter.DoubleQuoteFormat()
	fmt.Println(formattedText)

	formattedText = textFormatter.TabFormat()
	fmt.Println(formattedText)
}