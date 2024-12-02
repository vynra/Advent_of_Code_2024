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

func check_line(nums []int) ([]int, bool) {

	var inc bool
	var dec bool
	inc = false
	dec = false
	var safe bool = false
	var unsafes []int
	unsafes = nil
	for i := 1; i < len(nums); i++ {
		dist_temp := nums[i] - nums[i-1]
		if dist_temp == 0 || dist_temp >= 4 || dist_temp <= -4 {
			inc = false
			dec = false
			unsafes = append(unsafes, i)
		} else if dist_temp < 4 && dist_temp > 0 {
			if !inc && !dec {
				inc = true
			} else if dec {
				dec = false
				unsafes = append(unsafes, i)
			}
		} else {
			if !inc && !dec {
				dec = true
			} else if inc {
				inc = false
				unsafes = append(unsafes, i)
			}

		}
	}
	if inc || dec && len(unsafes) == 0 {
		fmt.Printf("%v\n", nums)
		fmt.Printf("%t %t\n", inc, dec)
		safe = true
	} else {
		fmt.Printf("%v\n", nums)

	}
	return unsafes, safe
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
		unsafes, safe := check_line(nums)
		if len(unsafes) == 0 && safe {
			safe_total += 1
			fmt.Printf("SAFE\n")
		} else if len(unsafes) == 2 {
			if unsafes[1]-unsafes[0] == 1 {
				var nums2 []int
				// fmt.Printf("%v\n", nums)
				var index int = unsafes[1]
				fmt.Printf("%d\n", index)

				nums2 = append(nums2, nums[:index]...)
				nums2 = append(nums2, nums[index+1:]...)
				unsafes, safe = check_line(nums2)
				if len(unsafes) == 0 && safe {
					safe_total += 1
					fmt.Printf("SAFE\n")
				} else {
					var nums2 []int
					// fmt.Printf("%v\n", nums)
					var index int = unsafes[0]
					fmt.Printf("%d\n", index)

					nums2 = append(nums2, nums[:index]...)
					nums2 = append(nums2, nums[index+1:]...)
					unsafes, safe = check_line(nums2)
					if len(unsafes) == 0 && safe {
						safe_total += 1
						fmt.Printf("SAFE\n")
					}

				}
			}
		} else if len(unsafes) == 1 {
			fmt.Printf("one unsafe\n")
			//remove right number
			var nums2 []int
			// fmt.Printf("%v\n", nums)
			var index int = unsafes[0]
			fmt.Printf("%d\n", index)

			nums2 = append(nums2, nums[:index]...)
			nums2 = append(nums2, nums[index+1:]...)
			unsafes2, safe2 := check_line(nums2)
			if len(unsafes2) == 0 && safe2 {
				safe_total += 1
				fmt.Printf("SAFE\n")
			} else {
				//remove left number
				nums2 = nil
				// fmt.Printf("%v\n", nums)
				index = unsafes[0] - 1
				fmt.Printf("%d\n", index)
				if index >= 0 {
					nums2 = append(nums2, nums[:index]...)
					nums2 = append(nums2, nums[index+1:]...)
					unsafes, safe = check_line(nums2)
					if len(unsafes) == 0 && safe {
						safe_total += 1
						fmt.Printf("SAFE\n")
					} else {

						//remove left number
						nums2 = nil
						// fmt.Printf("%v\n", nums)
						index = unsafes[0] - 2
						fmt.Printf("%d\n", index)
						if index >= 0 {

							nums2 = append(nums2, nums[:index]...)
							nums2 = append(nums2, nums[index+1:]...)
							unsafes, safe = check_line(nums2)
							if len(unsafes) == 0 && safe {
								safe_total += 1
								fmt.Printf("SAFE\n")
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", safe_total)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
