package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Define the flags
	setFlag := flag.String("set", "", "Set the symbols for percentage ranges 0-100%")
	overFlag := flag.String("over", "", "Set the symbols for percentages over 100%")
	flag.Parse()

	// Validate the usage
	if *setFlag == "" || flag.NArg() != 1 {
		fmt.Println("Usage: prettzel --set \"<symbols>\" <--over \"<over_symbols>\"> <percentage>")
		return
	}

	// Parse the percentage argument
	percentageStr := flag.Arg(0)
	percentage, err := strconv.ParseFloat(strings.TrimSuffix(percentageStr, "%"), 64)
	if err != nil {
		fmt.Println("Invalid percentage value")
		return
	}

	// Split the symbols into a slice
	symbolList := strings.Split(strings.Trim(*setFlag, "\""), " ")
	numSymbols := len(symbolList)

	// Determine the appropriate symbol list
	var selectedSymbol string
	if percentage <= 100 {
		index := int((percentage / 100) * float64(numSymbols-1))
		selectedSymbol = symbolList[index]
	} else if *overFlag != "" {
		overSymbolList := strings.Split(strings.Trim(*overFlag, "\""), " ")
		numOverSymbols := len(overSymbolList)
		index := int((percentage / 100) * float64(numOverSymbols-1))
		selectedSymbol = overSymbolList[index]
	} else {
		selectedSymbol = symbolList[numSymbols-1] // Default to the last symbol in the regular list
	}

	// Print the corresponding symbol
	fmt.Println(selectedSymbol)
}
