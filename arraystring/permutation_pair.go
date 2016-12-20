package arraystring

import (
  "strings"
  "strconv"
  "bytes"
)

const asciiOffset = 97;

// assumes all characters are alphabetical
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
    letterCountArr[v - asciiOffset]++
  }

  return getStringFromLetterCountArr(letterCountArr)
}

func IsPermutationPair(s1, s2 string) bool {
  s1Len := len(s1)
  s2Len := len(s2)

  if s1Len != s2Len {
    return false;
  }

  charMap := make(map[rune]int)

  for _, v := range s1 {
    // initially set each character key to true (not yet unseen) in s2
    charMap[v] += 1
  }

  for _, v := range s2 {
    // there are more characters of value 'v' in s2 than s1
    if charMap[v] ==  0 {
      return false
    } else {
      charMap[v]--
    }
  }

  for _, v := range charMap {
    if v != 0 {
      return false
    }
  }

  return true
}
