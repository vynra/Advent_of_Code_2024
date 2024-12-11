package main

import (
	// "bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
	"slices"
	"time"
)

func part_1(disk_map []int) int {
	var curr_i int = 0
	for i := len(disk_map) - 1; i > -1; i-- {
		if curr_i >= i {
			break
		}
		if disk_map[i] != -1 {
			idx := slices.Index(disk_map[curr_i:], -1)
			// fmt.Printf("%v\n", disk_map)
			idx += curr_i
			curr_i = idx + 1
			if curr_i >= i {
				break
			} else {
				disk_map[idx] = disk_map[i]
				disk_map[i] = -1
			}
		}
	}

	// fmt.Printf("%v\n", disk_map)
	var total int = 0
	for i := 0; i < len(disk_map); i++ {
		if disk_map[i] == -1 {
			break
		}
		total += i * disk_map[i]
	}
	return total
}

type blank struct {
	start  int
	length int
}

func remove_blank(blanks []blank, i int) []blank {
	return append(blanks[:i], blanks[i+1:]...)
}

func calc_blanks(disk_map []int) []blank {

	var blank_start int = 0
	var is_blank bool = false
	var blanks []blank
	for i := 0; i < len(disk_map); i++ {
		if disk_map[i] == -1 {
			if !is_blank {
				blank_start = i
				is_blank = true
			}
		} else {
			if is_blank {
				temp_blank := blank{start: blank_start, length: i - blank_start}
				blanks = append(blanks, temp_blank)
				is_blank = false
			}
		}
	}
	return blanks
}

func part_2(disk_map []int) int {
	// fmt.Printf("%v\n", disk_map)

	// fmt.Printf("%v\n", blanks)
	blanks := calc_blanks(disk_map)
	var file_end int = len(disk_map) - 1
	var is_file bool = true
	for i := len(disk_map) - 2; i > -1; i-- {
		blanks = calc_blanks(disk_map)
		// fmt.Printf("%v\n", blanks)
		// fmt.Printf("current: %d\n", disk_map[i])
		if disk_map[i] != -2 && disk_map[i] == disk_map[i+1] {
			if !is_file {
				file_end = i
				is_file = true
			}
		} else {
			if is_file {
				length := file_end - i
				// fmt.Printf("file length: %d\n", length)
				idx := slices.IndexFunc(blanks, func(b blank) bool { return b.length >= length })
				if idx != -1 {
					// fmt.Printf("blank: %v\n", blanks[idx])
					blank_index := blanks[idx].start
					// fmt.Printf("blank start: %d\n", blank_index)
					if blank_index < i {
						file_id := disk_map[i+1]
						for j := blank_index; j < length+blank_index; j++ {
							disk_map[j] = file_id
						}

						for j := i + 1; j < i+1+length; j++ {
							disk_map[j] = -1
						}
						//blanks stuff
						if length == blanks[idx].length {
							blanks = remove_blank(blanks, idx)
						} else {
							blanks[idx].start += length
							blanks[idx].length -= length
						}
					}

				}
				is_file = false
			}
		}

		if disk_map[i] != -1 && disk_map[i] != disk_map[i+1] {
			file_end = i
			is_file = true
		}
		// fmt.Printf("%v\n", disk_map)
	}

	// fmt.Printf("%v\n", disk_map)
	var total int = 0
	for i := 0; i < len(disk_map); i++ {
		if disk_map[i] != -1 {
			total += i * disk_map[i]
		}
	}
	return total
}

func main() {
	start := time.Now()
	f, err := os.ReadFile("../inputs/day_9.txt")

	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}
	var even bool = false
	var disk_map []int
	var curr_file int = 0
	// fmt.Printf("%s\n", f)

	for i := 0; i < len(f); i++ {
		// fmt.Printf("%d\n", int(f[i])-48)
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
	}
	disk_map_2 := make([]int, len(disk_map))
	for i := 0; i < len(disk_map); i++ {
		disk_map_2[i] = disk_map[i]
	}
	total := part_1(disk_map)
	total_2 := part_2(disk_map_2)
	// fmt.Printf("%v\n", disk_map_2)

	fmt.Printf("%d, %d\n", total, total_2)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
