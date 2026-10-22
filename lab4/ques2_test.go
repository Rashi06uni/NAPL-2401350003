
package main

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	got, err := divide(10, 2)
	if err != nil || got != 5 {
		t.Errorf("divide(10, 2) = %v, %v; want 5, nil", got, err)
	}

	_, err = divide(10, 0)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestSquareRoot(t *testing.T) {
	got, err := squareRoot(25)
	if err != nil || math.Abs(got-5) > 1e-9 {
		t.Errorf("squareRoot(25) = %v, %v; want 5, nil", got, err)
	}

	_, err = squareRoot(-4)
	if err == nil {
		t.Error("expected negative square root error")
	}
}

func TestValidateInput(t *testing.T) {
	got, err := validateInput("Hello")
	if err != nil || got != "Hello" {
		t.Errorf("validateInput(Hello) = %q, %v; want Hello, nil", got, err)
	}

	_, err = validateInput("")
	if err == nil {
		t.Error("expected empty input error")
	}
}
