package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.ReadFile("../inputs/day_3.txt")
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}
	renablers := regexp.MustCompile(`do(n't)?\(\)`)
	indices := renablers.FindAllStringIndex(string(f), -1)
	laststart := 0
	result := make([]string, len(indices)+1)
	for i, element := range indices {
		result[i] = string(f[laststart:element[0]])
		laststart = element[0]
	}
	result[len(indices)] = string(f[laststart:])
	re := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	var total int = 0
	for j := 0; j < len(result); j++ {
		// fmt.Printf("%s\n", result[j])
		if !strings.Contains(result[j], "don't") {
			muls := re.FindAllString(string(result[j]), -1)
			var num1 int
			var num2 int
			for i := 0; i < len(muls); i++ {
				reint := regexp.MustCompile(`\d{1,3}`)
				nums := reint.FindAllString(muls[i], -1)
				// fmt.Printf("%v\n", nums)
				num1, err = strconv.Atoi(nums[0])
				num2, err = strconv.Atoi(nums[1])
				total += num1 * num2
			}
		}
	}
	// fmt.Printf("%v\n", indices)
	// fmt.Printf("%s\n", result[1])
	fmt.Printf("%d\n", total)

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)

}
