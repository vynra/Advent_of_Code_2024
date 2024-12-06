package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	page_num int
	pages_b4 []int
}

func page_equals(page Page, i int) bool {
	return page.page_num == i
}

func pages_find_index(pages []Page, i int) int {
	for j := 0; j < len(pages); j++ {
		if page_equals(pages[j], i) {
			return j
		}
	}
	return -1
}

func part_1() {
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}

func part_2() {
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}

func main() {

	f, err := os.Open("../inputs/day_5_ex.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}
	var pages []Page
	scanner := bufio.NewScanner(f)
	var total = 0
	for scanner.Scan() {
		// fmt.Printf("READING LINE %d\n", line_num)
		line := scanner.Text()
		// fmt.Printf("%s\n", line)
		// parse second half
		split := strings.Split(line, "|")
		if len(split) == 1 {

			split := strings.Split(line, ",")
			if len(split) != 1 {
				fmt.Printf("%v\n", split)

				var line_valid = true
				split := strings.Split(line, ",")
				for i := 0; i < len(split); i++ {
					page_num, _ := strconv.Atoi(split[i])
					curr_page_i := pages_find_index(pages, page_num)
					curr_page := pages[curr_page_i]
					var page_valid = true
				check_line:
					for j := 0; j < len(curr_page.pages_b4); j++ {
					check_page:
						for k := i; k < len(split); k++ {
							if string(curr_page.pages_b4[j]) == split[k] {
								page_valid = false
								break check_page
							}
						}
						if !page_valid {
							line_valid = false
							break check_line
						}
					}
				}
				if line_valid {
					page_num, _ := strconv.Atoi(split[len(split)/2])
					total += page_num
					fmt.Printf("%d\n", total)
				}
			}

		} else {
			// parse first half, create pages
			first, _ := strconv.Atoi(split[0])
			second, _ := strconv.Atoi(split[1])
			first_page := &Page{}
			second_page := &Page{}
			temp_first_index := pages_find_index(pages, first)
			temp_second_index := pages_find_index(pages, second)
			if temp_first_index == -1 {
				page := Page{first, nil}
				pages = append(pages, page)
				first_page = &pages[len(pages)-1]
			} else {
				first_page = &pages[temp_first_index]
			}
			if temp_second_index == -1 {
				page := Page{second, nil}
				pages = append(pages, page)
				second_page = &pages[len(pages)-1]
			} else {
				second_page = &pages[temp_second_index]
			}
			if slices.Index(second_page.pages_b4, first_page.page_num) == -1 {
				second_page.pages_b4 = append(second_page.pages_b4, first_page.page_num)
			}
			// fmt.Printf("%d %d\n", first_page.page_num, second_page.page_num)
		}
	}
	for _, page := range pages {

		fmt.Printf("%v\n", page)
	}
	fmt.Printf("%d\n", total)
	part_1()
	part_2()
}
