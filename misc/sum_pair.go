package misc

// HasPairMatchingSum returns true if there are two ints in the array that
// equal the sum argument
func HasPairMatchingSum(numList []int, sum int) bool {
	head := 0
	tail := len(numList) - 1
	for head < tail {
		pairSum := numList[head] + numList[tail]
		if pairSum > sum {
			tail--
		} else if pairSum < sum {
			head++
		} else {
			return true
		}
	}

	return false
}
