package misc

import (
  "testing"
  "fmt"
)

type isPermutationPairTestCase struct {
  s1, s2 string
  expect bool
}

var isPermutationPairTestCaseProvider = []isPermutationPairTestCase{
  isPermutationPairTestCase{
    s1: "cat",
    s2: "tac",
    expect: true,
  },
  isPermutationPairTestCase{
    s1: "cat",
    s2: "bac",
    expect: false,
  },
  isPermutationPairTestCase{
    s1: "ccat",
    s2: "catt",
    expect: false,
  },
  isPermutationPairTestCase{
    s1: "cat",
    s2: "catt",
    expect: false,
  },
  isPermutationPairTestCase{
    s1: "tcat",
    s2: "catt",
    expect: true,
  },
}

func TestIsPermutationPair(t *testing.T) {
  for i, v := range isPermutationPairTestCaseProvider {
    output := IsPermutationPair(v.s1, v.s2)

    if output != v.expect {
      t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, v.expect, output))
    }
  }
}

type hasPermutationPairTestCase struct {
  input []string
  expect bool
}

var hasPermutationPairTestCaseProvider = []hasPermutationPairTestCase {
  hasPermutationPairTestCase{
    input: []string{"cat", "bat", "cart", "tart", "tac"},
    expect: true,
  },
  hasPermutationPairTestCase{
    input: []string{"god", "dog", "flu", "blue", "du"},
    expect: true,
  },
  hasPermutationPairTestCase{
    input: []string{"dart", "starter", "tekken", "doop", "dork"},
    expect: false,
  },
  hasPermutationPairTestCase{
    input: []string{"before", "erofeb", "cart", "tarc", "tac"},
    expect: true,
  },
  hasPermutationPairTestCase{
    input: []string{"tart", "tac"},
    expect: false,
  },
  hasPermutationPairTestCase{
    input: []string{},
    expect: false,
  },
}

func TestHasPermutationPair(t *testing.T) {
  for i, v := range hasPermutationPairTestCaseProvider {
    output := HasPermutationPair(v.input)

    if output != v.expect {
      t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, v.expect, output))
    }
  }
}
