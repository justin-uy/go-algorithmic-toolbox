package arraystring

import (
  "testing"
  "fmt"
)

type testCase struct {
  input string
  expect bool
}

var testCaseProvider = []testCase {
  testCase {
    input: "a",
    expect: true,
  },
  testCase {
    input: "ab",
    expect: true,
  },
  testCase {
    input: "aba",
    expect: false,
  },
  testCase {
    input: "cat",
    expect: true,
  },
  testCase {
    input: "dad",
    expect: false,
  },
  testCase {
    input: "animal",
    expect: false,
  },
}

func TestAllUnique(t *testing.T) {
  for i, v := range testCaseProvider {
    output := HasAllUniqueChars(v.input)
    if output != v.expect {
      t.Error(fmt.Sprintf("Test %v - Expected: %v; Got: %v", i, v.expect, output))
    }
  }
}
