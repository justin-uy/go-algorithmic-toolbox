package arraystring

// HasAllUnique returns true if all runes in the string are unique
// lower case and upper case letters are considered different.
func HasAllUnique(s string) bool {
	runes := []rune(s)
	seenChars := make(map[rune]bool)

	for _, v := range runes {
		if !seenChars[v] {
			seenChars[v] = true
		} else {
			return false
		}
	}

	return true
}
