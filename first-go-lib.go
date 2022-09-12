/**
 * Author: Alex Aloshchenko
 * File: first-go-lib.go
 * Package description: The package with simple addition, subtraction, division and multiplication functions.
 *				If a math operation causes integer overflow, the function returns: 0, error.
 *				The package is designed to learn the process of creating third-party Golang packages.
 */

package lib

import (
	"errors"
	"math"
)

var ErrOverflow = errors.New("integer overflow")

// Add Returns the sum of two integer numbers.
// In case of integer overflow, the function returns: 0, error.
func Add(a int, b int) (int, error) {
	if b > 0 {
		if a > math.MaxInt-b {
			return 0, ErrOverflow
		}
	} else {
		if a < math.MinInt-b {
			return 0, ErrOverflow
		}
	}
	return a + b, nil
}

// Subtract Returns the difference between two integer numbers.
// In case of integer overflow, the function returns: 0, error.
func Subtract(a int, b int) (int, error) {
	if b > 0 {
		if a < math.MinInt+b {
			return 0, ErrOverflow
		}
	} else {
		if a > math.MaxInt+b {
			return 0, ErrOverflow
		}
	}
	return a - b, nil
}

// Multiply Returns the multiplication of two integer numbers.
// In case of integer overflow, the function returns: 0, error.
func Multiply(a int, b int) (int, error) {
	if b > math.MaxInt/a || a > math.MaxInt/b {
		return 0, ErrOverflow
	}
	return a * b, nil
}

// Division  Returns the division of two integer numbers.
// Only quotient is returned, remainder is ignored.
func Division(a int, b int) int {
	return a / b
}
