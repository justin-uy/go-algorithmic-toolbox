package week1

import (
	"errors"
	"strconv"
	"strings"
)

func getFirstTwoValuesFromStringArray(arr []string) (int, int, error) {
	x, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, nil
	}

	y, err := strconv.Atoi(arr[1])

	if err != nil {
		return 0, 0, nil
	}

	return x, y, nil
}

// MaxPairProduct takes a string where there should be numbers separated by
// spaces. If the input is valid, then we find the product of the two largest
// numbers
func MaxPairProduct(input string) (int, error) {
	values := strings.Split(input, " ")

	numValues := len(values)

	// invalid input; return an error
	if numValues < 2 {
		return 0, errors.New("Invalid input; too few arguments")
	}

	a, b, err := getFirstTwoValuesFromStringArray(values)
	if err != nil {
		return 0, err
	}

	// If there are only two values we can just return the product immediately
	if numValues == 2 {
		return a * b, err
	}

	// we can guarantee that the array has more than 2 values after this point
	var max1, max2 int

	if a < b {
		max1 = b
		max2 = a
	} else {
		max1 = a
		max2 = b
	}

	for i := 2; i < numValues; i++ {
		a, err := strconv.Atoi(values[i])

		if err != nil {
			return 0, err
		}

		if a > max2 {
			if a > max1 {
				max2 = max1
				max1 = a
			} else {
				max2 = a
			}
		}
	}

	return max1 * max2, nil
}
