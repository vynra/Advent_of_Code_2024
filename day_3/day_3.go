package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.ReadFile("../inputs/day_3.txt")
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}
	re := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	muls := re.FindAllString(string(f), -1)
	var total int = 0
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
	fmt.Printf("%d\n", total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
