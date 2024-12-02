package main

import (
	//	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	start := time.Now()
	//dat, err := os.ReadFile("../inputs/day_1")
	//check(err)
	//fmt.Print(string(dat))
	f, err := os.Open("../inputs/day_1.txt")
	//check(err)
	var left []int
	var right []int
	b1 := make([]byte, 0)
	n1, err := f.Read(b1)
	var temp int

	var flag bool = true
	for flag {
		b1 = make([]byte, 5)
		n1, err = f.Read(b1)
		//		check(err)

		temp, err = strconv.Atoi(string(b1[:n1]))
		left = append(left, temp)

		b1 = make([]byte, 3)
		n1, err = f.Read(b1)
		//		check(err)

		b1 = make([]byte, 5)
		n1, err = f.Read(b1)
		//		check(err)
		temp, err = strconv.Atoi(string(b1[:n1]))
		right = append(right, temp)

		b1 = make([]byte, 1)
		n1, err = f.Read(b1)
		//		check(err)

		if err == io.EOF {
			flag = false
		}
	}

	//fmt.Printf("%v", left)
	//fmt.Printf("%v", right)
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// fmt.Printf("%v", left)
	//fmt.Printf("%v", right)

	var dist int = 0
	var dist_temp int
	var right_index int = 0
	var sim_score int = 0
	var occs int
	var same_flag bool = false
	for i := 0; i < len(left); i++ {
		//part 1
		dist_temp = left[i] - right[i]
		if dist_temp < 0 {
			dist_temp = -dist_temp
		}
		dist += dist_temp
		//part 2
		same_flag = false
		if i > 0 && left[i] == left[i-1] {
			sim_score += left[i] * occs
			same_flag = true
		}
		occs = 0
		var dist_temp2 int = left[i] - right[right_index]
		if dist_temp2 >= 0 && !same_flag {
			for dist_temp2 >= 0 {
				if dist_temp2 == 0 {
					occs += 1
				}
				if right_index < len(left)-1 {
					right_index += 1
					dist_temp2 = left[i] - right[right_index]
				} else {
					break
				}
			}
			sim_score += left[i] * occs
		}

	}
	// fmt.Printf("%d\n", sim_score)
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
	fmt.Printf("%d, %d\n", dist, sim_score)
}
