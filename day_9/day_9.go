package main

import (
	// "bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.ReadFile("../inputs/day_9_ex.txt")

	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}
	var even bool = false
	var disk_map []int
	var curr_file int = 0
	// fmt.Printf("%s\n", f)

	for i := 0; i < len(f); i++ {
		fmt.Printf("%d\n", int(f[i])-48)
		if even {
			for j := 0; j < int(f[i])-48; j++ {
				disk_map = append(disk_map, -1)
			}
			even = false
		} else {
			for k := 0; k < int(f[i])-48; k++ {
				disk_map = append(disk_map, curr_file)
			}
			even = true
			curr_file += 1
		}
		fmt.Printf("%v\n", disk_map)
	}

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
