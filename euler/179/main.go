package main

import (
	"fmt"
	"os"
	"time"
)

const minInputNumber int = 3
const maxInputNumber int = 10000000

var divisors [maxInputNumber + 1]uint16

func main() {

	inputNumber := readInput()
	validateInput(inputNumber, minInputNumber, maxInputNumber)

	start := time.Now()
	populateDivisors(inputNumber)
	fmt.Print(countDivisors(inputNumber))
	duration := time.Since(start)
	fmt.Print(" ")
	fmt.Print(duration)
}

func readInput() int {
	var input int
	_, err := fmt.Scanf("%d", &input)

	if err != nil {
		fmt.Println("Invalid number entered")
		os.Exit(1)
	}

	return input
}

func validateInput(input int, min int, max int) {
	if input < min || input > max {
		fmt.Printf("Input number must be between %d and %d\n", min, max)
		os.Exit(1)
	}
}

func populateDivisors(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; i*j <= number; j++ {
			divisors[i*j]++
		}
	}
}

func countDivisors(number int) int {
	count := 0
	for i := 2; i < number; i++ {
		if divisors[i] == divisors[i+1] {
			count++
		}
	}

	return count
}
