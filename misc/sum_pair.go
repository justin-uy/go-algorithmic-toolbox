package misc

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
