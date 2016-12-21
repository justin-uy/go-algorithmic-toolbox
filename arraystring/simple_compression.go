package arraystring

import (
	"bytes"
	"strconv"
)

// Input: ""; Expect ""
// Input: "a"; Expect "a"
// Input: "aabb"; Expect "aabb"
// Input: "aaa"; Expect "a3"
// Input: "aabcccccaaa"; Expect: "a2b1c5a3"
func SimpleCompression(s string) string {
	sSize := len(s)
	// Any string less than 3 characters will always return the original string
	if sSize < 3 {
		return s
	}

	compressedSize := 0
	var compressedBuffer bytes.Buffer
	for i := 0; i < sSize; {
		current := s[i]
		currentCount := 1
		for j := i + 1; j < sSize && s[j] == current; j, currentCount = j+1, currentCount+1 {
		}

		i += currentCount
		compressedBuffer.WriteString(string(current) + strconv.Itoa(currentCount))

		// if compressed string is greater than or equal to the original string just return the original string
		if compressedSize += 2; compressedSize >= sSize {
			return s
		}
	}

	return compressedBuffer.String()
}
