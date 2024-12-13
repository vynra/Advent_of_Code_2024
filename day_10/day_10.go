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
var curr_route []loc

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
func check_trail(y int, x int, curr_height int, thi int, curr_route []loc) bool {
	var curloc loc
	var up bool = false
	var right bool = false
	var down bool = false
	var left bool = false
	curloc_i := slices.IndexFunc(locs, func(l loc) bool { return l.x == x && l.y == y })
	for i := 0; i < len(curr_route); i++ {
		fmt.Printf("%v\n", curr_route[i])
	}
	// fmt.Printf("%d\n", curloc_i)
	if curloc_i == -1 {
		curloc = loc{x: x, y: y, visited: true}
		locs = append(locs, curloc)
		// fmt.Printf("NEW LOC\n")
		// fmt.Printf("%d, %d, %d, %d\n", x, y, curr_height, byte(curr_height+1))
		// fmt.Printf("left: %d, %d\n", lines[y][x-1], byte(curr_height+1))

		// fmt.Printf("Height: %d\n", curr_height)
		if x > 0 && lines[y][x-1]-48 == byte(curr_height+1) {
			fmt.Printf("LEFT\n")
			curr_route = append(curr_route, curloc)
			if check_trail(y, x-1, curr_height+1, thi, curr_route) {
				left = true
			}

		}
		// fmt.Printf("WHYYYYYYYY: %d\n", y)
		if y > 0 && lines[y-1][x]-48 == byte(curr_height+1) {
			fmt.Printf("UP\n")
			curr_route = append(curr_route, curloc)
			if check_trail(y-1, x, curr_height+1, thi, curr_route) {
				up = true
			}
		}
		if x < len(lines[y])-1 && lines[y][x+1]-48 == byte(curr_height+1) {
			fmt.Printf("RIGHT\n")
			fmt.Printf("%v\n", curloc)
			curr_route = append(curr_route, curloc)
			if check_trail(y, x+1, curr_height+1, thi, curr_route) {
				right = true
			}
		}
		if y < len(lines)-1 && lines[y+1][x]-48 == byte(curr_height+1) {
			fmt.Printf("DOWN\n")
			curr_route = append(curr_route, curloc)
			if check_trail(y+1, x, curr_height+1, thi, curr_route) {
				down = true
			}
		}
	} else {
		curloc = locs[curloc_i]
		if len(locs[curloc_i].peaks) > 0 {
			fmt.Printf("ALREADY VISITED PEAK??\n")
			trailheads[thi].peaks = append(trailheads[thi].peaks, locs[curloc_i].peaks...)
			for j := 0; j < len(curr_route); j++ {
				curloc_j := slices.IndexFunc(locs, func(l loc) bool { return l.x == curr_route[j].x && l.y == curr_route[j].y })
				locs[curloc_j].peaks = append(curr_route[j].peaks, curloc.peaks...)
			}
			return true
		}
	}
	if curr_height == 9 {
		temp_peak := peak{x: x, y: y}
		fmt.Printf("ITS 9 PEAK??\n")
		trailheads[thi].peaks = append(trailheads[thi].peaks, temp_peak)
		fmt.Printf("current route:\n")

		for j := 0; j < len(curr_route); j++ {
			curloc_j := slices.IndexFunc(locs, func(l loc) bool { return l.x == curr_route[j].x && l.y == curr_route[j].y })
			locs[curloc_j].peaks = append(curr_route[j].peaks, temp_peak)
		}
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
	}
	for i := 0; i < len(trailheads); i++ {
		check_trail(trailheads[i].y, trailheads[i].x, 0, i, curr_route)
	}
	for i := 0; i < len(trailheads); i++ {
		fmt.Printf("%v\n", trailheads[i])
	}
	fmt.Printf("\n LOCS\n")
	for i := 0; i < len(locs); i++ {
		fmt.Printf("%v\n", locs[i])
	}

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
