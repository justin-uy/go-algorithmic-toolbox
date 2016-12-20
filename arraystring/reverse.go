package arraystring

import (
  "strings"
)

func Reverse(input string) string {
  strSlice := ReverseSlice(StrSlice(strings.Split(input, "")))
  // StrSlice always implements Join and will not throw an error

  aStrSlice, _ := strSlice.(StrSlice)
  return strings.Join(aStrSlice, "")
}

type SliceInterface interface {
  Len() int
  Swap(i, j int)
  ToGenericSlice() []interface{}
}

func ReverseSlice(input SliceInterface) SliceInterface {
  i := 0
  j := input.Len() - 1

  for i < j {
    input.Swap(i, j)
    i++;
    j--;
  }
  return input
}

func SlicesEqual(a, b []interface{}) bool {
  if len(a) != len(b) {
    return false
  }

  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      return false
    }
  }

  return true
}

type StrSlice []string

func (p StrSlice) Len() int {
  return len(p)
}

func (p StrSlice) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}

func (p StrSlice) ToGenericSlice() []interface{} {
  genericSlice := make([]interface{}, p.Len())
  for i, v := range p {
    genericSlice[i] = v
  }

  return genericSlice
}

type IntSlice []int

func (p IntSlice) Len() int {
  return len(p)
}

func (p IntSlice) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}

func (p IntSlice) ToGenericSlice() []interface{} {
  genericSlice := make([]interface{}, p.Len())
  for i, v := range p {
    genericSlice[i] = v
  }

  return genericSlice
}
