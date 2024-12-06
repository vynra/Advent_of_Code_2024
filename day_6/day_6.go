package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var full_map []string
var coords [][]int

func print_coords() {
	for _, line := range coords {
		fmt.Printf("%v\n", line)
	}
}

func print_map() {
	for _, line := range full_map {
		fmt.Printf("%v\n", line)
	}
}

func move(x int, y int, direction string) {
	period_byte := []byte(".")
	fmt.Printf("%v\n", period_byte)
	if x < 0 || y < 0 || x >= len(coords[0]) || y >= len(coords) {
		return
	}
	coords[y][x] = 1

	if y == 0 && direction == "up" {
		coords[y][x] = 1
		return
	}
	if x == len(coords[0])-1 && direction == "right" {
		coords[y][x] = 1
		return
	}
	if y == len(coords)-1 && direction == "down" {
		coords[y][x] = 1
		return
	}
	if x == 0 && direction == "left" {
		coords[y][x] = 1
		return
	}

	fmt.Printf("%v\n", full_map[x][y+1])
	switch direction {
	case "up":
		if full_map[y-1][x] == period_byte[0] {
			fmt.Printf("true\n")
			move(x, y-1, "up")
		} else {
			move(x, y, "right")
		}
	case "right":
		if full_map[y][x+1] == period_byte[0] {
			fmt.Printf("true\n")
			move(x+1, y, "right")
		} else {
			move(x, y, "down")
		}
	case "down":
		if full_map[y+1][x] == period_byte[0] {
			fmt.Printf("true\n")
			move(x, y+1, "down")
		} else {
			move(x, y, "left")
		}
	case "left":
		if full_map[y][x-1] == period_byte[0] {
			fmt.Printf("true\n")
			move(x-1, y, "left")
		} else {
			move(x, y, "up")
		}
	}
}

func part_1() {
	start := time.Now()
	f, _ := os.Open("../inputs/day_6.txt")

	scanner := bufio.NewScanner(f)
	var start_1 int
	var start_2 int
	for scanner.Scan() {
		// fmt.Printf("READING LINE %d\n", line_num)
		line := scanner.Text()
		pos := strings.Index(line, "^")
		if pos != -1 {
			start_1 = pos
			start_2 = len(full_map)
			line = strings.Replace(line, "^", ".", -1)
		}
		full_map = append(full_map, line)
	}
	for i := 0; i < len(full_map); i++ {
		coords_row := make([]int, len(full_map[0]))
		coords = append(coords, coords_row)
	}

	move(start_1, start_2, "up")

	print_coords()
	print_map()
	var total int = 0
	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords[0]); j++ {
			total += coords[i][j]
		}
	}

	fmt.Printf("%d\n", total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}

func part_2() {

	start := time.Now()
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)

}

func main() {
	part_1()
	part_2()
}
