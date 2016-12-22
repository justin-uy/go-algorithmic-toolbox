package bigint

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type BigInt string

func invalidBigInt() error {
	return errors.New("Invalid value for BigInt")
}

func New(val interface{}) (BigInt, error) {
	switch val.(type) {
	case int:
		return BigInt(strconv.Itoa(val.(int))), nil
	case string:
		valStr := val.(string)
		// make sure there are no leading zeros and preserve the negative
		if len(valStr) > 0 {
			if valStr[0:1] != "-" {
				valStr = strings.TrimLeft(valStr, "0")
			} else {
				valStr = "-" + strings.TrimLeft(valStr[1:], "0")
			}
		} else {
			return "", invalidBigInt()
		}

		// validate and make sure that the valStr is a valid BigInt
		intRegex, _ := regexp.Compile(`^-?\d+`)
		if !intRegex.MatchString(valStr) {
			return "", invalidBigInt()
		}

		return BigInt(valStr), nil
	}

	return "", nil
}

func (b BigInt) IsNegative() bool {
	digits := []rune(string(b))
	if string(digits[0]) == "-" {
		return true
	}
	return false
}

// Add takes two BigInts as arguments and returns the sum of them as BigInt
// Currently BigInts that are less than int64 max are treated the same as actual BigInts
// Further optimizations will be done in the future to ensure that we use regular integer
// arithmetic for values that will not overflow in64, which leverages the processor
// much better
func (b1 BigInt) Add(b2 BigInt) BigInt {
	// start at the first digit (last index) of b1 and b2
	// assume initially that b1 is longer and then update after
	max := len(b1) - 1
	maxVal := b1
	min := len(b2) - 1
	minVal := b2

	if min > max {
		temp := max
		tempVal := maxVal
		max = min
		maxVal = minVal
		min = temp
		minVal = tempVal
	}

	result := ""
	carry := 0
	for min > 0 {
		digit1, _ := strconv.Atoi(string(minVal[min:]))
		digit2, _ := strconv.Atoi(string(maxVal[max:]))

		sum := digit1 + digit2 + carry
		carry = 0

		if sum > 9 {
			carry = 1
		}
		result += strconv.Itoa(sum)

		min--
		max--
	}

	// add one because the upper range is exclusive
	max += 1

	if max > 0 {
		maxValRest := string(maxVal[0:max])
		if carry == 0 {
			result = maxValRest + result
		} else {
			maxRestBigInt, _ := New(maxValRest)
			carryBigInt, _ := New(carry)
			restSum := maxRestBigInt.Add(carryBigInt)
			// concat strings (not addition)
			result = string(restSum) + result
		}
	}

	resultBigInt, _ := New(result)

	return resultBigInt
}
