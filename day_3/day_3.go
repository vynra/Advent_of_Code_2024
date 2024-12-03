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

func getSubtotal(result string) int {
	re := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	muls := re.FindAllString(string(result), -1)
	var total int = 0
	for i := 0; i < len(muls); i++ {
		reint := regexp.MustCompile(`\d{1,3}`)
		nums := reint.FindAllString(muls[i], -1)
		// fmt.Printf("%v\n", nums)
		num1, err := strconv.Atoi(nums[0])
		num2, err := strconv.Atoi(nums[1])
		if err == nil {
			total += num1 * num2
		}
	}
	return total
}

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
	var total_1 = 0
	var total_2 int = 0
	for j := 0; j < len(result); j++ {
		// fmt.Printf("%s\n", result[j])
		if !strings.Contains(result[j], "don't") {
			total_1 += getSubtotal(result[j])
			total_2 += getSubtotal(result[j])
		} else {
			total_1 += getSubtotal(result[j])
		}
	}
	// fmt.Printf("%v\n", indices)
	// fmt.Printf("%s\n", result[1])
	fmt.Printf("%d\n", total_1)
	fmt.Printf("%d\n", total_2)

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)

}
