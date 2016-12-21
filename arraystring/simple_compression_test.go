package arraystring

import (
	"fmt"
	"testing"
)

func TestSimpleCompression(t *testing.T) {
	cases := []struct{ in, expect string }{
		{in: "", expect: ""},
		{in: "a", expect: "a"},
		{in: "ab", expect: "ab"},
		{in: "aabb", expect: "aabb"},
		{in: "aaa", expect: "a3"},
		{in: "aabcccccaaa", expect: "a2b1c5a3"},
	}

	for i, c := range cases {
		out := SimpleCompression(c.in)
		if out != c.expect {
			t.Error(fmt.Sprintf("Test %v - Expect %v; Got %v", i, c.expect, out))
		}
	}
}
