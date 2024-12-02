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

func check_line(nums []int) bool {

	var inc bool
	var dec bool
	inc = false
	dec = false
	var safe bool = false
	for i := 1; i < len(nums); i++ {
		dist_temp := nums[i] - nums[i-1]
		if dist_temp == 0 || dist_temp >= 4 || dist_temp <= -4 {
			inc = false
			dec = false
			break
		} else if dist_temp < 4 && dist_temp > 0 {
			if !inc && !dec {
				inc = true
			} else if dec {
				dec = false
				break
			}
		} else {
			if !inc && !dec {
				dec = true
			} else if inc {
				inc = false
				break
			}
		}
	}
	if inc || dec {
		safe = true
	}
	return safe
}

func main() {

	start := time.Now()
	f, err := os.Open("../inputs/day_2.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", f.Name(), err)
	}

	scanner := bufio.NewScanner(f)
	var safe_total int = 0
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Fields(line)
		var nums []int
		for i := 0; i < len(strs); i++ {
			var temp_num int
			temp_num, err = strconv.Atoi(strs[i])
			nums = append(nums, temp_num)
		}
		var safe bool = false
		safe = check_line(nums)
		if safe {
			safe_total += 1
		} else {
			var nums2 []int
			for i := 0; i < len(nums); i++ {
				nums2 = nil
				nums2 = append(nums2, nums[:i]...)
				nums2 = append(nums2, nums[i+1:]...)
				safe = check_line(nums2)
				if safe {
					safe_total += 1
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", safe_total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
