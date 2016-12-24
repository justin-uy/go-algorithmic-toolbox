package stringutils

import (
	"testing"
)

func TestAllUnique(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{input: "a", expect: true},
		{input: "ab", expect: true},
		{input: "aba", expect: false},
		{input: "cat", expect: true},
		{input: "dad", expect: false},
		{input: "animal", expect: false},
	}

	for i, c := range cases {
		out := HasAllUnique(c.input)
		if out != c.expect {
			t.Errorf("Test %v - Expected: %v; Got: %v", i, c.expect, out)
		}
	}
}
