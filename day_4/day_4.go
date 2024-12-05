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
	line_num := 0
	var lines []string
	var coords [][]int
	for scanner.Scan() {
		// fmt.Printf("READING LINE %d\n", line_num)
		line := scanner.Text()
		lines = append(lines, line)
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
		// fmt.Printf("%v\n", indices)
		coords = append(coords, indices)
		line_num += 1

	}

	var total int = 0
	for i := 0; i < 140; i++ {
		for j := 0; j < len(coords[i]); j++ {
			k := coords[i][j]
			if k > 2 {
				if lines[i][k-1] == 'M' && lines[i][k-2] == 'A' && lines[i][k-3] == 'S' {
					total += 1
				}
			}
			if k < 137 {
				if lines[i][k+1] == 'M' && lines[i][k+2] == 'A' && lines[i][k+3] == 'S' {
					total += 1
				}
			}
			if i > 2 {
				if lines[i-1][k] == 'M' && lines[i-2][k] == 'A' && lines[i-3][k] == 'S' {
					total += 1
				}
			}
			if i < 137 {
				if lines[i+1][k] == 'M' && lines[i+2][k] == 'A' && lines[i+3][k] == 'S' {
					total += 1
				}
			}
			if k > 2 && i > 2 {
				if lines[i-1][k-1] == 'M' && lines[i-2][k-2] == 'A' && lines[i-3][k-3] == 'S' {
					total += 1
				}
			}
			if k < 137 && i > 2 {
				if lines[i-1][k+1] == 'M' && lines[i-2][k+2] == 'A' && lines[i-3][k+3] == 'S' {
					total += 1
				}
			}
			if k > 2 && i < 137 {
				if lines[i+1][k-1] == 'M' && lines[i+2][k-2] == 'A' && lines[i+3][k-3] == 'S' {
					total += 1
				}
			}
			if k < 137 && i < 137 {
				if lines[i+1][k+1] == 'M' && lines[i+2][k+2] == 'A' && lines[i+3][k+3] == 'S' {
					total += 1
				}
			}
		}
	}

	// fmt.Printf("%v\n", coords)
	// fmt.Printf("%d\n", len(coords))
	fmt.Printf("%d\n", total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
