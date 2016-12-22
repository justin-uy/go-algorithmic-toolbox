package stringutils

import (
	"learngo/arrayutils"
)

// Reverse returns a new string where the runes are in reverse order
func Reverse(input string) string {
	runeSlice := arrayutils.RuneSlice([]rune(input))
	return string(arrayutils.Reverse(runeSlice).(arrayutils.RuneSlice))
}
