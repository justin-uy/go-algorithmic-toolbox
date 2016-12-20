package arraystring

import (
  "testing"
  "fmt"
)

type reverseTestCase struct {
  input, expect string
}

var reverseTestCaseProvider = []reverseTestCase {
  reverseTestCase {
    input: "",
    expect: "",
  },
  reverseTestCase {
    input: "a",
    expect: "a",
  },
  reverseTestCase {
    input: "ab",
    expect: "ba",
  },
  reverseTestCase {
    input: "abc",
    expect: "cba",
  },
}

func TestReverse(t *testing.T) {
  for i, testCase := range reverseTestCaseProvider {
    out := Reverse(testCase.input)
    if out != testCase.expect {
      t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, testCase.expect, out))
    }
  }
}

type reverseSliceTestCase struct {
  input, expect SliceInterface
}

var reverseSliceTestCaseProvider = []reverseSliceTestCase {
  reverseSliceTestCase {
    input: StrSlice([]string{}),
    expect: StrSlice([]string{}),
  },
  reverseSliceTestCase {
    input: StrSlice([]string{"cat"}),
    expect: StrSlice([]string{"cat"}),
  },
  reverseSliceTestCase {
    input: StrSlice([]string{"cat", "dog"}),
    expect: StrSlice([]string{"dog", "cat"}),
  },
  reverseSliceTestCase {
    input: StrSlice([]string{"cat", "dog", "bat"}),
    expect: StrSlice([]string{"bat", "dog", "cat"}),
  },
  reverseSliceTestCase {
    input: IntSlice([]int{}),
    expect: IntSlice([]int{}),
  },
  reverseSliceTestCase {
    input: IntSlice([]int{1}),
    expect: IntSlice([]int{1}),
  },
  reverseSliceTestCase {
    input: IntSlice([]int{1, 2}),
    expect: IntSlice([]int{2, 1}),
  },
  reverseSliceTestCase {
    input: IntSlice([]int{1, 2, 3}),
    expect: IntSlice([]int{3, 2, 1}),
  },
}

func TestReverseSlice(t *testing.T) {
  for i, testCase := range reverseSliceTestCaseProvider {
    out := ReverseSlice(testCase.input)
    if !SlicesEqual(out.ToGenericSlice(), testCase.expect.ToGenericSlice()) {
      t.Error(fmt.Sprintf("Test %v - Expected %v; Got %v", i, testCase.expect, out))
    }
  }
}
