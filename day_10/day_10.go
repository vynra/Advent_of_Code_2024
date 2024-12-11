package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strconv"
	// "strings"
	"slices"
	"time"
)

var lines []string
var trailheads []trailhead
var peaks []peak
var locs []loc

type loc struct {
	x       int
	y       int
	visited bool
	peaks   []peak
}

type peak struct {
	x int
	y int
}

type trailhead struct {
	x     int
	y     int
	peaks []peak
}

// inputs current location
func check_trail(x int, y int, curr_height int, thi int, curr_route []loc) bool {
	var curloc loc
	var up bool = false
	var right bool = false
	var down bool = false
	var left bool = false
	curloc_i := slices.IndexFunc(locs, func(l loc) bool { return l.x == x && l.y == y })
	// fmt.Printf("%d\n", curloc_i)
	if curloc_i == -1 {
		curloc = loc{x: x, y: y, visited: true}
		locs = append(locs, curloc)
		// fmt.Printf("%v\n", locs)

		if x > 0 && lines[y][x-1] == byte(curr_height+1) {
			if check_trail(y, x-1, curr_height+1, thi, curr_route) {
				curr_route = append(curr_route, curloc)
				left = true
			}

		}
		if y > 0 && lines[y-1][x] == byte(curr_height+1) {
			if check_trail(y-1, x, curr_height+1, thi, curr_route) {
				curr_route = append(curr_route, curloc)
				down = true
			}
		}
		if x < len(lines[y])-1 && lines[y][x+1] == byte(curr_height+1) {
			if check_trail(y, x+1, curr_height+1, thi, curr_route) {
				curr_route = append(curr_route, curloc)
				right = true
			}
		}
		if y < len(lines)-1 && lines[y+1][x] == byte(curr_height+1) {
			if check_trail(y+1, x, curr_height+1, thi, curr_route) {
				curr_route = append(curr_route, curloc)
				up = true
			}
		}
	} else {
		curloc = locs[curloc_i]
		if len(locs[curloc_i].peaks) > 0 {
			fmt.Printf("ALREADY VISITED PEAK??\n")
			trailheads[thi].peaks = append(trailheads[thi].peaks, locs[curloc_i].peaks...)
			return true
		}
	}
	if curr_height == 9 {
		temp_peak := peak{x: x, y: y}
		fmt.Printf("ITS 9 PEAK??\n")
		trailheads[thi].peaks = append(trailheads[thi].peaks, temp_peak)
		return true
	}
	return up || right || down || left
}

func main() {
	start := time.Now()

	f, err := os.Open("../inputs/day_10_ex.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		for i := 0; i < len(line); i++ {
			temp, _ := strconv.Atoi(string(line[i]))
			if temp == 0 {
				// fmt.Printf("TRAILHEAD")
				temp_th := trailhead{x: i, y: len(lines) - 1}
				trailheads = append(trailheads, temp_th)

			} else if temp == 9 {
				temp_peak := peak{x: i, y: len(lines) - 1}
				peaks = append(peaks, temp_peak)
			}
		}

		for i := 0; i < len(trailheads); i++ {
			var curr_route []loc
			check_trail(trailheads[i].x, trailheads[i].y, 0, i, curr_route)
		}
	}

	fmt.Printf("%v\n", trailheads)
	fmt.Printf("%v\n", locs)

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
