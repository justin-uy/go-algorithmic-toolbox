package arrayutils

// Reverse will reverse the elements in the slice and return it
func Reverse(input SliceInterface) SliceInterface {
	i := 0
	j := input.Len() - 1

	for i < j {
		input.Swap(i, j)
		i++
		j--
	}
	return input
}
