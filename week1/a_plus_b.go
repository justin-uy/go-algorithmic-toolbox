package week1

import (
  "strings"
  "strconv"
  "errors"
)

func APlusB(input string) (int, error) {
  values := strings.Split(input, " ")
  sum := 0

  if (len(values) != 2) {
    return 0, errors.New("Invalid input: incorrect number of arguments")
  }

  for i := 0; i < len(values); i++ {
    v, err := strconv.Atoi(values[i])
    if err != nil {
      return 0, err
    }

    sum += v
  }

  return sum, nil
}
