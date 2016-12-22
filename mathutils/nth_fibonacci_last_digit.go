package mathutils

import (
	"errors"
)

// GetNthFibonacciLastDigit returns that last digit in the Nth value in the
// Fibonacci sequence
func GetNthFibonacciLastDigit(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Fibonacci sequence starts at 1")
	} else if n < 3 {
		return 1, nil
	}

	nMinus1 := 1
	nMinus2 := 1

	for i := 3; i <= n; i++ {
		nVal := nMinus1 + nMinus2

		if nVal > 9 {
			nVal = nVal % 10
		}

		if i == n {
			return nVal, nil
		}

		nMinus2 = nMinus1
		nMinus1 = nVal
	}

	return 0, errors.New("Should never reach here")
}
