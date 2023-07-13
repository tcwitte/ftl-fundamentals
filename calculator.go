// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying the second
// with the first.
func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, errors.New("division by zero")
	}
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative numbers not supported")
	}
	return math.Sqrt(a), nil
}
