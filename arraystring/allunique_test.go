package arraystring

import (
  "testing"
  "fmt"
)

func TestAllUnique(t *testing.T) {
  cases := []struct{input string; expect bool} {
    {input: "a", expect: true},
    {input: "ab", expect: true},
    {input: "aba", expect: false},
    {input: "cat", expect: true},
    {input: "dad", expect: false},
    {input: "animal", expect: false},
  }

  for i, c := range cases {
    output := HasAllUniqueChars(c.input)
    if output != c.expect {
      t.Error(fmt.Sprintf("Test %v - Expected: %v; Got: %v", i, c.expect, output))
    }
  }
}
