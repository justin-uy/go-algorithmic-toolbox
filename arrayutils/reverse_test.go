package arrayutils

import (
	"fmt"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	cases := []struct{ input, expect SliceInterface }{
		{
			input:  StrSlice([]string{}),
			expect: StrSlice([]string{}),
		},
		{
			input:  StrSlice([]string{"cat"}),
			expect: StrSlice([]string{"cat"}),
		},
		{
			input:  StrSlice([]string{"cat", "dog"}),
			expect: StrSlice([]string{"dog", "cat"}),
		},
		{
			input:  StrSlice([]string{"cat", "dog", "bat"}),
			expect: StrSlice([]string{"bat", "dog", "cat"}),
		},
		{
			input:  IntSlice([]int{}),
			expect: IntSlice([]int{}),
		},
		{
			input:  IntSlice([]int{1}),
			expect: IntSlice([]int{1}),
		},
		{
			input:  IntSlice([]int{1, 2}),
			expect: IntSlice([]int{2, 1}),
		},
		{
			input:  IntSlice([]int{1, 2, 3}),
			expect: IntSlice([]int{3, 2, 1}),
		},
	}

	for i, c := range cases {
		out := Reverse(c.input)
		if !SlicesEqual(out.ToGenericSlice(), c.expect.ToGenericSlice()) {
			t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, c.expect, out))
		}
	}
}
