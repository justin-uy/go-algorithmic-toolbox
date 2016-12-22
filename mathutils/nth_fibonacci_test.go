package mathutils

import (
	"fmt"
	"testing"
)

func TestNthFibonacci(t *testing.T) {
	cases := []struct{ input, expect int }{
		{input: 0, expect: 0},
		{input: 1, expect: 1},
		{input: 2, expect: 1},
		{input: 3, expect: 2},
		{input: 4, expect: 3},
		{input: 10, expect: 55},
		{input: 20, expect: 6765},
		{input: 26, expect: 121393},
	}

	for i, c := range cases {
		out := GetNthFibonacci(c.input)
		if out != c.expect {
			t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, c.expect, out))
		}
	}
}
