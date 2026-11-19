
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Function to accept and validate integer input
func readInt(scanner *bufio.Scanner, message string) int64 {
	for {
		fmt.Print(message)
		if !scanner.Scan() {
			fmt.Println("Input ended.")
			os.Exit(1)
		}

		num, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)

		if err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			continue
		}

		if num < -1000000 || num > 1000000 {
			fmt.Println("Out of range! Enter a value between -1000000 and 1000000.")
			continue
		}

		return num
	}
}

// Function to accept and validate floating-point input
func readFloat(scanner *bufio.Scanner, message string) float64 {
	for {
		fmt.Print(message)
		if !scanner.Scan() {
			fmt.Println("Input ended.")
			os.Exit(1)
		}

		num, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

		if err != nil {
			fmt.Println("Invalid input! Please enter a numeric value.")
			continue
		}

		if math.IsNaN(num) || math.IsInf(num, 0) ||
			num < -1000000 || num > 1000000 {
			fmt.Println("Out of range! Enter a valid number between -1000000 and 1000000.")
			continue
		}

		return num
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("===== SIMPLE CALCULATOR =====")

	// Integer operations
	fmt.Println("\n--- INTEGER OPERATIONS ---")

	a := readInt(scanner, "Enter first integer: ")
	b := readInt(scanner, "Enter second integer: ")

	fmt.Println("\nInteger Results:")
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)

	// Floating-point operations
	fmt.Println("\n--- FLOATING-POINT OPERATIONS ---")

	x := readFloat(scanner, "Enter first float: ")
	y := readFloat(scanner, "Enter second float: ")

	fmt.Println("\nFloating-Point Results:")
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)

	fmt.Println("\nCalculator completed successfully!")
}
