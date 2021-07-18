package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	array, testcases := readInput()

	start := time.Now()
	for i := 0; i < int(testcases); i++ {
		larrysArrayV1(array[i])
	}
	duration := time.Since(start)
	fmt.Printf("Larrys Array v1: %s\n", duration)

	start = time.Now()
	for i := 0; i < int(testcases); i++ {
		larrysArrayV2(array[i])
	}
	duration = time.Since(start)
	fmt.Printf("Larrys Array v2: %s\n", duration)
}

func readInput() ([][]int32, int64) {
	file, err := os.Open("test-cases.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	testcases, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	array := make([][]int32, 0)

	for i := 0; i < int(testcases); i++ {
		scanner.Scan()
		arrayLength, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		tmp := make([]int32, 0)

		scanner.Scan()
		stringArray := strings.Split(strings.TrimRight(scanner.Text(), " \t\r\n"), " ")

		for j := 0; j < int(arrayLength); j++ {
			value, _ := strconv.ParseInt(stringArray[j], 10, 64)
			tmp = append(tmp, int32(value))
		}

		array = append(array, tmp)
	}

	return array, testcases
}

func larrysArrayV1(array []int32) string {
	inverions := 0

	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[i] < array[j] {
				inverions++
			}
		}
	}

	if inverions%2 == 0 {
		return "YES"
	}

	return "NO"
}

func larrysArrayV2(array []int32) string {
	var oddSwap uint8
	var i int32

	for ; int(i) < len(array); i++ {
		for array[i] != i+1 {
			array[i], array[array[i]-1] = array[array[i]-1], array[i]
			oddSwap = 1 - oddSwap
		}
	}

	if oddSwap == 0 {
		return "YES"
	}

	return "NO"
}
