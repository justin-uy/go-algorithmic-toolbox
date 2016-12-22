package arraystring

// Reverse returns a new string where the runes are in reverse order
func Reverse(input string) string {
	runeSlice := RuneSlice([]rune(input))
	return string(ReverseSlice(runeSlice).(RuneSlice))
}

// SliceInterface is an interface for a variety of slice alias; this allow
// us to do some generic operations on slices without as much repeated code
type SliceInterface interface {
	Len() int
	Swap(i, j int)
	ToGenericSlice() []interface{}
}

// ReverseSlice will reverse the elements in the slice and return it
func ReverseSlice(input SliceInterface) SliceInterface {
	i := 0
	j := input.Len() - 1

	for i < j {
		input.Swap(i, j)
		i++
		j--
	}
	return input
}

// SlicesEqual is used to evaluate equality between two slices. Equality
// is defined as two slices having the same elements in the same order
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

// RuneSlice is an alias for []rune
type RuneSlice []rune

// Len returns the length of the slice
func (p RuneSlice) Len() int {
	return len(p)
}

// Swap will swap the values in a slice given to indices
func (p RuneSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// ToGenericSlice This essentially casts a RuneSlice to a generic slice
func (p RuneSlice) ToGenericSlice() []interface{} {
	genericSlice := make([]interface{}, p.Len())
	for i, v := range p {
		genericSlice[i] = v
	}

	return genericSlice
}

// StrSlice is an alias for []string
type StrSlice []string

// Len returns the length of the slice
func (p StrSlice) Len() int {
	return len(p)
}

// Swap will swap the values in a slice given to indices
func (p StrSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// ToGenericSlice This essentially casts a RuneSlice to a generic slice
func (p StrSlice) ToGenericSlice() []interface{} {
	genericSlice := make([]interface{}, p.Len())
	for i, v := range p {
		genericSlice[i] = v
	}

	return genericSlice
}

// IntSlice is an alias for []string
type IntSlice []int

// Len returns the length of the slice
func (p IntSlice) Len() int {
	return len(p)
}

// Swap will swap the values in a slice given to indices
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// ToGenericSlice This essentially casts a RuneSlice to a generic slice
func (p IntSlice) ToGenericSlice() []interface{} {
	genericSlice := make([]interface{}, p.Len())
	for i, v := range p {
		genericSlice[i] = v
	}

	return genericSlice
}
