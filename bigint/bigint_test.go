package bigint

import (
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
			t.Errorf("Test %v - Expected %v; Got %v", i, c.expect, b)
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
			t.Errorf("Test %v - Expected: %v; Got: %v", i, c.expect, out)
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
			t.Errorf("Test %v - Expected: %v; Got: %v", i, c.expect, out)
		}
	}
}

func TestAbs(t *testing.T) {
	cases := []struct {
		b, expect BigInt
	}{
		{b: BigInt("0"), expect: BigInt("0")},
		{b: BigInt("-1"), expect: BigInt("1")},
		{b: BigInt("12"), expect: BigInt("12")},
		{b: BigInt("-12340"), expect: BigInt("12340")},
	}

	for i, c := range cases {
		out := c.b.Abs()
		if out != c.expect {
			t.Errorf("Test %v - Expected: %v; Got %v", i, c.expect, out)
		}
	}
}

func TestIsIntaable(t *testing.T) {
	cases := []struct {
		b      BigInt
		expect bool
	}{
		{b: BigInt("0"), expect: true},
		{b: BigInt("-20"), expect: true},
		{b: BigInt("1000"), expect: true},
		{b: BigInt("-9223372036854775807"), expect: true},
		{b: BigInt("-9223372036854775808"), expect: false},
		{b: BigInt("18446744073709551615"), expect: true},
		{b: BigInt("18446744073709551616"), expect: false},
	}

	for i, c := range cases {
		out := c.b.IsIntable()
		if out != c.expect {
			t.Errorf("Test %v - Expected: %v; Got %v", i, c.expect, out)
		}
	}
}
