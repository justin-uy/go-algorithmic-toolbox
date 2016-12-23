package bigint

import (
	"fmt"
	"testing"
)

const MaxInt int64 = 1<<63 - 1

func TestNew(t *testing.T) {
	cases := []struct {
		input  interface{}
		expect string
	}{
		{input: "1", expect: "1"},
		{input: "-1", expect: "-1"},
		{input: "01", expect: "1"},
		{input: "10", expect: "10"},
		{input: "1000000000000", expect: "1000000000000"},
		{input: "la", expect: ""},
		{input: "", expect: ""},
		{input: 1, expect: "1"},
		{input: -1, expect: "-1"},
		{input: 10, expect: "10"},
		{input: 1000000000000, expect: "1000000000000"},
		{input: 0, expect: "0"},
		{input: int(MaxInt), expect: "9223372036854775807"},
	}

	for i, c := range cases {
		b, _ := New(c.input)
		if string(b) != c.expect {
			t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, c.expect, b))
		}
	}
}

func TestIsNegative(t *testing.T) {
	cases := []struct {
		input  BigInt
		expect bool
	}{
		{input: BigInt("0"), expect: false},
		{input: BigInt("322939"), expect: false},
		{input: BigInt("-124"), expect: true},
		{input: BigInt("-1"), expect: true},
	}

	for i, c := range cases {
		out := c.input.IsNegative()
		if out != c.expect {
			t.Error(fmt.Sprintf("Test %v - Expected: %v; Got: %v", i, c.expect, out))
		}
	}
}

func TestIsEqual(t *testing.T) {
	cases := []struct {
		b1, b2 BigInt
		expect bool
	}{
		{b1: BigInt("0"), b2: BigInt("0"), expect: true},
		{b1: BigInt("234"), b2: BigInt("432"), expect: false},
		{b1: BigInt("-123"), b2: BigInt("-123"), expect: true},
		{b1: BigInt("-12314"), b2: BigInt("12341"), expect: false},
		{b1: BigInt("-2345"), b2: BigInt("-2345"), expect: true},
	}

	for i, c := range cases {
		out := c.b1.IsEqual(c.b2)
		if out != c.expect {
			t.Error(fmt.Sprintf("Test %v - Expected: %v; Got: %v", i, c.expect, out))
		}
	}
}
