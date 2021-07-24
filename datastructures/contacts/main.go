package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"strconv"
    "strings"
)

func contacts(queries [][]string) []int32 {
	queriesIndex := make(map[string]int32)
	var results []int32

	for _, queriesRowItem := range queries {
		if queriesRowItem[0] == "add"  {
			for i := 1; i <= len(queriesRowItem[1]); i++ {
				queriesIndex[queriesRowItem[1][:i]] += 1
			}
		} else {
			results = append(results, queriesIndex[queriesRowItem[1]])
		}
	}

	return results
}

func main() {
	file, err := os.Open("input.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	queriesRows, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)

	if err != nil {
		log.Fatalf("failed reading first line: %s", err)
	}

	var queries [][]string
	
	for i := 0; i < int(queriesRows); i++ {
		scanner.Scan()

		queriesRowTemp := strings.Split(strings.TrimRight(scanner.Text()," \t\r\n"), " ")
		var queriesRow []string

		queriesRow = append(queriesRow, queriesRowTemp[0], queriesRowTemp[1])
        queries = append(queries, queriesRow)
	}
 
	file.Close()

	

	start := time.Now()
	contacts(queries)
	duration := time.Since(start)
	fmt.Println(duration)
}
