package mathutils

import (
	"testing"
)

func TestGCD(t *testing.T) {
	cases := []struct {
		a, b, expect int
		hasErr       bool
	}{
		{a: 0, b: 1, expect: 0, hasErr: true},
		{a: 4, b: 2, expect: 2, hasErr: false},
		{a: 10, b: 15, expect: 5, hasErr: false},
		{a: 22, b: 18, expect: 2, hasErr: false},
		{a: 24, b: 32, expect: 8, hasErr: false},
		{a: 50, b: 10, expect: 10, hasErr: false},
		{a: 357, b: 234, expect: 3, hasErr: false},
		{a: 3918848, b: 1653264, expect: 61232, hasErr: false},
	}

	for i, c := range cases {
		// number less than 0 will err
		out, err := GCD(c.a, c.b)

		if c.expect != out || (err != nil && !c.hasErr) {
			if err != nil {
				t.Errorf("Test %v - Error %v", i, err)
			} else {
				t.Errorf("Test %v - Expect %v; Got %v", i, c.expect, out)
			}
		}
	}
}
