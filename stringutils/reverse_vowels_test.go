package stringutils

import "testing"

func TestReverseVowels(t *testing.T) {
	cases := []struct{ input, expect string }{
		{input: "", expect: ""},
		{input: "a", expect: "a"},
		{input: "b", expect: "b"},
		{input: "ab", expect: "ab"},
		{input: "abe", expect: "eba"},
		{input: "apple", expect: "eppla"},
		{input: "input", expect: "unpit"},
	}

	for i, c := range cases {
		out := ReverseVowels(c.input)
		if out != c.expect {
			t.Errorf("Test %v - Expected %v; Got %v", i, c.expect, out)
		}
	}
}
