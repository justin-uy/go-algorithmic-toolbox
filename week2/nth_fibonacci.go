package week2

// GetNthFibonacci will attempt to return the Nth value in the Fibonacci sequence
// There will be an integer overflow if given too large a value
func GetNthFibonacci(n int) int {
	if n <= 1 {
		return n
	} else if n == 2 {
		return 1
	}

	nMinus1 := 1
	nMinus2 := 1
	var nVal int

	for i := 3; i <= n; i++ {
		nVal = nMinus1 + nMinus2

		if i == n {
			break
		}

		nMinus2 = nMinus1
		nMinus1 = nVal
	}

	return nVal
}
