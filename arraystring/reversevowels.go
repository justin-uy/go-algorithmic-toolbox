package arraystring

import (
	"strings"
)

func isVowel(c rune) bool {
	switch strings.ToLower(string(c)) {
	case "a", "e", "i", "o", "u":
		return true
	}
	return false
}

func ReverseVowels(s string) string {
	runes := []rune(s)
	l := 0
	r := len(s) - 1
	for l < r {
		isLVowel := isVowel(runes[l])
		isRVowel := isVowel(runes[r])

		if isLVowel && isRVowel {
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
		if !isLVowel {
			l++
		}
		if !isRVowel {
			r--
		}
	}

	return string(runes)
}
