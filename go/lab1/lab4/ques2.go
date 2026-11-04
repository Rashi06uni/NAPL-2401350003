
package main

import (
	"fmt"
	"math"
	"strings"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("square root of a negative number is not allowed")
	}
	return math.Sqrt(n), nil
}

func validateInput(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return input, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = squareRoot(-4)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = validateInput("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
