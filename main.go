package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: prettzel --set \"<symbols>\" <percentage>")
		return
	}

	// Parse the arguments
	symbols := os.Args[2]
	percentageStr := os.Args[3]

	// Clean the symbols input
	symbols = strings.Trim(symbols, "\"")

	// Convert percentage to a float
	percentage, err := strconv.ParseFloat(strings.TrimSuffix(percentageStr, "%"), 64)
	if err != nil {
		fmt.Println("Invalid percentage value")
		return
	}

	// Split symbols into a slice
	symbolList := strings.Split(symbols, " ")

	// Determine the index of the symbol based on the percentage
	numSymbols := len(symbolList)
	index := int((percentage / 100) * float64(numSymbols-1))

	// Print the corresponding symbol
	fmt.Println(symbolList[index])
}
