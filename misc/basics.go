// All GO prgrams start with package main
package misc

// Factored imports are the standard convention
import (
	"fmt"
	"math"
)

// Functions: We can make functions like this; this will not be exports because it's
// lower case
//
// Parameters: the parameter type goes after the variable name
// if all variables share the same type, the variables can be comma separated
// with the one type declaration following the last argument
//
// Return types: immediately affter the arguments before the curly brace
func add(x, y int) int {
	return x + y
}

// Main function is where all of it starts
func BasicLanguageFeatures() {
	// There are no semi-colons to end statements
	fmt.Println("This is how you print a line. Notice how it's a capital \"P\"")
	// Printf can do print format
	fmt.Printf("This must be print format. My name is: %s\n", "Justin")
	// Notice that exported functions and variables start with a capital letter
	fmt.Printf("This is pi: %g. This is the square root of pi: %g\n", math.Pi, math.Sqrt(math.Pi))
	// You prefix local variables with the keyword var or short-hand :=
	a, b := 1, 2 // no need to specify a type if you assign an initial value
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Println("Each", "argument", "is", "joined", "on", "spaces.")

	// here are all the types and a factored declaration
	var (
		isTrue bool   = true
		word   string = "hello"
		// This is only supported if the processor is 64 bit
		biggestNumber uint64 = 1<<64 - 1
		biggestByte   byte   = 1<<8 - 1 // byte is alias for uint8
	)

	fmt.Println(isTrue, word, biggestByte, biggestNumber)
	fmt.Printf("%v is of type %T\n", isTrue, isTrue)

	var (
		defaultString string  // "" empty string
		defaultNumber int     // 0
		defaultFloat  float64 // 0
		defaultBool   bool    // false
	)

	fmt.Println(defaultString, defaultNumber, defaultFloat, defaultBool)

	aFloat := 1.1
	aInt := 3

	fmt.Printf("%v - %T as %v - int; %v - %T as %v - float\n", aFloat, aFloat, uint(aFloat), aInt, aInt, float64(aInt))

	fmt.Printf("a is %v\n", a)

	// a is a variable and can be re-assigned
	a = 5

	fmt.Printf("a is now %v\n", a)

	// const keyword is used
	const unchangingFloat = 10.1
	fmt.Println(unchangingFloat)
}
