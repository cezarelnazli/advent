package main

import (
	"fmt"
)

var used []bool
var sol, containers []int
var current, num, objective, min_len int

func part_one(level int) {
	for i := sol[level] + 1; i <= len(containers); i++ {
		if !used[i-1] {
			used[i-1] = true
			sol[level+1] = i - 1
			current += containers[i-1]
			if current == objective {
				if level < min_len {
					min_len = level
					num = 1
				} else if level == min_len {
					num++
				}
			}
			part_one(level + 1)
			current -= containers[i-1]
			used[i-1] = false
		}
	}
}

func main() {
	containers = []int{
		11, 30, 47, 31, 32, 36, 3, 1, 5, 3,
		32, 36, 15, 11, 46, 26, 28, 1, 19, 3}

	used = make([]bool, len(containers), len(containers))
	sol = make([]int, len(containers)+1, len(containers)+1)
	objective = 150
	min_len = 30

	part_one(0)

	fmt.Println(num)
}
