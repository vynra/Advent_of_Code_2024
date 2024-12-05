package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	page_num int
	pages_b4 []*Page
}

func page_equals(page Page, i int) bool {
	return page.page_num == i
}

func pages_find_index(pages []*Page, i int) int {
	for j := 0; j < len(pages); j++ {
		if page_equals(*pages[j], i) {
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
	var pages []*Page
	scanner := bufio.NewScanner(f)
scanning:
	for scanner.Scan() {
		// fmt.Printf("READING LINE %d\n", line_num)
		line := scanner.Text()
		// fmt.Printf("%s\n", line)
		// parse first half, create pages
		split := strings.Split(line, "|")
		if len(split) != 2 {
			break scanning
		}
		first, err := strconv.Atoi(split[0])
		second, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Failed to convert nums %v", err)
		}
		first_page := &Page{}
		second_page := &Page{}
		temp_first_index := pages_find_index(pages, first)
		temp_second_index := pages_find_index(pages, second)
		if temp_first_index == -1 {
			page := Page{first, nil}
			pages = append(pages, &page)
			first_page = pages[len(pages)-1]
		} else {
			first_page = pages[temp_first_index]
		}
		if temp_second_index == -1 {
			page := Page{second, nil}
			pages = append(pages, &page)
			second_page = pages[len(pages)-1]
		} else {
			second_page = pages[temp_second_index]
		}
		if pages_find_index(first_page.pages_b4, second_page.page_num) == -1 {
			first_page.pages_b4 = append(first_page.pages_b4, second_page)
		}
		fmt.Printf("%d %d\n", first_page.page_num, second_page.page_num)
	}
	for _, page := range pages {
		fmt.Printf("%v\n", page)
	}
	part_1()
	part_2()
}
