package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("../inputs/day_4.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}

	scanner := bufio.NewScanner(f)
	var total int = 0
	line_num := 0
	for scanner.Scan() {
		fmt.Printf("READING LINE %d\n", line_num)
		line := scanner.Text()
		// fmt.Printf("%d\n", len(line))
		var indices []int
		i := 0
		prev := 0
	exes:
		for {
			i = strings.Index(line[prev:], "X")
			// fmt.Printf("%d\n", i)
			if i == -1 || i == len(line)-1 {
				break exes
			}
			indices = append(indices, i+prev)
			prev += i + 1
		}
		fmt.Printf("%v\n", indices)
		line_num += 1

	}

	fmt.Printf("%d\n", total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
