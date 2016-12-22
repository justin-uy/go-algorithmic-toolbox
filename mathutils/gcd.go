package mathutils

import (
	"errors"
)

// GCD returns the greatest common divisor of two numbers
func GCD(a, b int) (int, error) {
	if a < 1 || b < 1 {
		return 0, errors.New("Both arguments must be greater than 0")
	}

	// identifying the min and max simplifies the mental math
	big := a
	small := b
	if a < b {
		big = b
		small = a
	}

	// if the two number are divisible by each other then the gcd is the min
	// of the two numbers
	bigPrime := big % small
	if bigPrime == 0 {
		return small, nil
	}
	// Euclidean algorithm due to the fact that
	// GCD(A, B == GCD(A', B), where A' = A % B
	return GCD(bigPrime, small)
}
