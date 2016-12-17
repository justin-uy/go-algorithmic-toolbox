package main

import (
	"algorithmic_toolbox/misc"
	"algorithmic_toolbox/week1"
	"algorithmic_toolbox/week2"
	"fmt"
	"log"
)

func printTestOutput(name string, input string, output int, expected int) {
	fmt.Printf(
		"~%v~\nInput: %v; Output: %v; Expected: %v\n",
		name,
		input,
		output,
		expected,
	)
}

func main() {
	aPlusBInput := "12 3"
	sum, err := week1.APlusB("12 3")

	if err != nil {
		log.Fatal(err)
	}

	printTestOutput("A plus B", aPlusBInput, sum, 12+3)

	maxPairProductInput := "12 4 8 23 1 -8 10"

	product, err := week1.MaxPairProduct(maxPairProductInput)

	if err != nil {
		log.Fatal(err)
	}

	printTestOutput("Max Pairwise Product", maxPairProductInput, product, 12*23)

	for i := 1; i <= 50; i++ {
		fibOutput, err := week2.GetNthFibonacci(i)
		fibLastDigitOutput, err2 := week2.GetNthFibonacciLastDigit(i)

		if err != nil || err2 != nil {
			log.Fatal(err)
		}

		fmt.Printf("Fib - %v: %v\n", i, fibOutput)
		fmt.Printf("Fib - %v last digit: %v\n", i, fibLastDigitOutput)
	}

	inputSums := [...]int{3, 6, 5, 10}
	inputArray := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(inputSums); i++ {
		sum := inputSums[i]
		hasMatch := misc.HasPairMatchingSum(inputArray[0:len(inputArray)], sum)

		if hasMatch {
			fmt.Printf("Array %v has a pair a values that sum up to %v\n", inputArray, sum)
		} else {
			fmt.Printf("Array %v doesn't have a pair a values that sum up to %v\n", inputArray, sum)
		}
	}
}
