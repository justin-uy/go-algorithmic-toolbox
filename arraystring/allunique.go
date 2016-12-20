package arraystring

import (
  "strings"
)

func HasAllUniqueChars(s string) bool {
  chars := strings.Split(s, "")
  seenChars := make(map[string]bool)

  for _, v := range chars {
    if !seenChars[v] {
      seenChars[v] = true
    } else {
      return false
    }
  }

  return true
}
