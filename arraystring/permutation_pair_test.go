package arraystring

import (
  "testing"
  "fmt"
)

func TestIsPermutationPair(t *testing.T) {
  cases := []struct{s1, s2 string; expect bool} {
    {
      s1: "cat",
      s2: "tac",
      expect: true,
    },
    {
      s1: "cat",
      s2: "bac",
      expect: false,
    },
    {
      s1: "ccat",
      s2: "catt",
      expect: false,
    },
    {
      s1: "cat",
      s2: "catt",
      expect: false,
    },
    {
      s1: "tcat",
      s2: "catt",
      expect: true,
    },
  }

  for i, c := range cases {
    output := IsPermutationPair(c.s1, c.s2)

    if output != c.expect {
      t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, c.expect, output))
    }
  }
}

func TestHasPermutationPair(t *testing.T) {
  cases := []struct{input []string; expect bool} {
    {
      input: []string{"cat", "bat", "cart", "tart", "tac"},
      expect: true,
    },
    {
      input: []string{"god", "dog", "flu", "blue", "du"},
      expect: true,
    },
    {
      input: []string{"dart", "starter", "tekken", "doop", "dork"},
      expect: false,
    },
    {
      input: []string{"before", "erofeb", "cart", "tarc", "tac"},
      expect: true,
    },
    {
      input: []string{"tart", "tac"},
      expect: false,
    },
    {
      input: []string{},
      expect: false,
    },
  }

  for i, c := range cases {
    output := HasPermutationPair(c.input)

    if output != c.expect {
      t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, c.expect, output))
    }
  }
}
