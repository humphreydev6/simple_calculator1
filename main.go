package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the user provided enough arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <operation> <num1> <num2>")
		return
	}

	// Parse the operation and operands
	operation := os.Args[1]
	num1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid operand:", os.Args[2])
		return
	}
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid operand:", os.Args[3])
		return
	}

	// Perform the requested operation
	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("Invalid operation:", operation)
		return
	}

	// Print the result
	fmt.Println(result)
}
