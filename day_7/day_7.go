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

func check_op(goal int, curr_val int, nums []string) bool {
	if curr_val > goal {
		return false
	}
	if len(nums) == 0 {
		if curr_val == goal {
			return true
		} else {
			return false
		}
	}
	var mult_curr_val int = curr_val
	var add_curr_val int = curr_val
	var cat_curr_val string
	num, _ := strconv.Atoi(nums[0])
	mult_curr_val *= num
	if check_op(goal, mult_curr_val, nums[1:]) {
		return true
	}
	add_curr_val += num
	if check_op(goal, add_curr_val, nums[1:]) {
		return true
	}
	cat_curr_val = strconv.Itoa(curr_val) + nums[0]
	temp, _ := strconv.Atoi(cat_curr_val)
	if check_op(goal, temp, nums[1:]) {
		return true
	}
	return false
}

func main() {

	start := time.Now()
	f, err := os.Open("../inputs/day_7.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}
	scanner := bufio.NewScanner(f)
	var total int = 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		goal, _ := strconv.Atoi(split[0])
		split2 := strings.Split(split[1], " ")
		start_num, _ := strconv.Atoi(split2[1])
		valid := check_op(goal, start_num, split2[2:])
		if valid {
			total += goal
		}
	}
	fmt.Printf("%d\n", total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
