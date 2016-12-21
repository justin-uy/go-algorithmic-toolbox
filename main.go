package main

import (
	"algorithmic_toolbox/week2"
	"fmt"
	"log"
)

func main() {
	for i := 1; i <= 92; i++ {
		fibOutput, err := week2.GetNthFibonacci(i)
		fibLastDigitOutput, err2 := week2.GetNthFibonacciLastDigit(i)

		if err != nil || err2 != nil {
			log.Fatal(err)
		}

		fmt.Printf("Fib - %v: %v\n", i, fibOutput)
		fmt.Printf("Fib - %v last digit: %v\n", i, fibLastDigitOutput)
	}
}
