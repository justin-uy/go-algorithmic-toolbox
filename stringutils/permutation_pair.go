package stringutils

import (
	"bytes"
	"strconv"
	"strings"
)

const asciiOffset = 97

// HasPermutationPair assumes all characters are alphabetical and returns true
// if there exists one or more pairs of strings in the string array that are
// permutations of each other
func HasPermutationPair(input []string) bool {
	sortedWordDictionary := make(map[string]bool)

	for _, word := range input {
		sortedWord := sortString(strings.ToLower(word))
		if !sortedWordDictionary[sortedWord] {
			sortedWordDictionary[sortedWord] = true
		} else {
			return true
		}
	}

	return false
}

func getStringFromLetterCountArr(letterCountArr [26]int) string {
	var wordBuffer bytes.Buffer
	for i, v := range letterCountArr {
		char := strconv.QuoteRune(rune(i + asciiOffset))
		for j := 0; j < v; j++ {
			wordBuffer.WriteString(char)
		}
	}

	return wordBuffer.String()
}

func sortString(word string) string {
	letterCountArr := [26]int{}
	for _, v := range []rune(word) {
		// minux 97 ascii offset
		letterCountArr[v-asciiOffset]++
	}

	return getStringFromLetterCountArr(letterCountArr)
}

// IsPermutationPair returns true if the two string arguments are permutations
// of each other. Permutation is defined as a word has the same characters
// but may not be in the same order
func IsPermutationPair(s1, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)

	if s1Len != s2Len {
		return false
	}

	charMap := make(map[rune]int)

	for _, v := range s1 {
		// initially set each character key to true (not yet unseen) in s2
		charMap[v]++
	}

	for _, v := range s2 {
		// there are more characters of value 'v' in s2 than s1
		if charMap[v] == 0 {
			return false
		}

		charMap[v]--

	}

	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}

	return true
}
