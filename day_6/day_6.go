package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type direction struct {
	visited bool
	up      bool
	right   bool
	down    bool
	left    bool
}

var full_map [][]byte
var coords [][]int
var dirs [][]direction

func reset_all() {
	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords[0]); j++ {
			coords[i][j] = 0
			dirs[i][j] = direction{visited: false, up: false, right: false, left: false, down: false}
		}
	}
}

func print_coords() {
	for _, line := range coords {
		fmt.Printf("%v\n", line)
	}
}

func print_map() {
	for _, line := range full_map {
		fmt.Printf("%v\n", string(line))
	}
}

func move(x int, y int, direction string, loop bool) bool {
	period_byte := []byte(".")
	if x < 0 || y < 0 || x >= len(coords[0]) || y >= len(coords) {
		return loop
	}
	coords[y][x] = 1
	dirs[y][x].visited = true
	if y == 0 && direction == "up" {
		if dirs[y][x].up {
			loop = true
			return loop
		}
		return loop
	}
	if x == len(coords[0])-1 && direction == "right" {
		if dirs[y][x].right {
			loop = true
			return loop
		}
		return loop
	}
	if y == len(coords)-1 && direction == "down" {
		if dirs[y][x].down {
			loop = true
			return loop
		}
		return loop
	}
	if x == 0 && direction == "left" {
		if dirs[y][x].left {
			loop = true
			return loop
		}
		return loop
	}

	switch direction {
	case "up":
		if dirs[y][x].up {
			loop = true
			return loop
		}
		dirs[y][x].up = true
		if full_map[y-1][x] == period_byte[0] {
			loop = move(x, y-1, "up", loop)
		} else {
			loop = move(x, y, "right", loop)
		}
	case "right":
		if dirs[y][x].right {
			loop = true
			return loop
		}
		dirs[y][x].right = true
		if full_map[y][x+1] == period_byte[0] {
			loop = move(x+1, y, "right", loop)
		} else {
			loop = move(x, y, "down", loop)
		}
	case "down":
		if dirs[y][x].down {
			loop = true
			return loop
		}
		dirs[y][x].down = true
		if full_map[y+1][x] == period_byte[0] {
			loop = move(x, y+1, "down", loop)
		} else {
			loop = move(x, y, "left", loop)
		}
	case "left":
		if dirs[y][x].left {
			loop = true
			return loop
		}
		dirs[y][x].left = true
		if full_map[y][x-1] == period_byte[0] {
			loop = move(x-1, y, "left", loop)
		} else {
			loop = move(x, y, "up", loop)
		}
	}
	return loop
}

func main() {
	start := time.Now()
	f, _ := os.Open("../inputs/day_6.txt")

	scanner := bufio.NewScanner(f)
	var start_1 int
	var start_2 int
	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, "^")
		if pos != -1 {
			start_1 = pos
			start_2 = len(full_map)
			line = strings.Replace(line, "^", ".", -1)
		}
		full_map = append(full_map, []byte(line))
	}

	var orig_coords [][]int
	for i := 0; i < len(full_map); i++ {
		coords_row := make([]int, len(full_map[0]))
		orig_row := make([]int, len(full_map[0]))
		coords = append(coords, coords_row)
		orig_coords = append(orig_coords, orig_row)
		var dirs_row []direction
		for j := 0; j < len(full_map[0]); j++ {
			dir := direction{visited: false, up: false, right: false, left: false, down: false}
			dirs_row = append(dirs_row, dir)
		}
		dirs = append(dirs, dirs_row)
	}

	move(start_1, start_2, "up", false)

	var total_coords int = 0
	var total_obs int = 0
	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords[0]); j++ {
			orig_coords[i][j] = coords[i][j]
		}
	}

	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords[0]); j++ {
			total_coords += orig_coords[i][j]
			if orig_coords[i][j] == 1 {
				full_map[i][j] = 35
				reset_all()
				loop := move(start_1, start_2, "up", false)
				if loop {
					total_obs += 1
				}
				full_map[i][j] = 46
			}
		}
	}
	fmt.Printf("%d %d\n", total_coords, total_obs)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
