package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var all_nums []int
var sol []int
var lqe int
var found bool
var used []bool

func bkt(l, np, target int) {
	if l == np {
		sum := 0
		qe := 1

		for _, w := range sol[1:np] {
			sum += all_nums[w]
			qe *= all_nums[w]
		}

		if sum == target {
			if qe < lqe {
				lqe = qe
			}
			found = true
		}
	} else {
		for i := sol[l-1] + 1; i <= len(all_nums); i++ {
			if !used[i-1] {
				used[i-1] = true
				sol[l] = i - 1
				bkt(l+1, np, target)
				used[i-1] = false
			}
		}
	}
}

func solve(input string, part int) int {
	var n, sum int
	all_nums = make([]int, 0)
	sol = make([]int, 0)
	used = make([]bool, 0)

	for _, l := range strings.Split(input, "\n") {
		fmt.Sscanf(l, "%d", &n)
		all_nums = append(all_nums, n)
		sol = append(sol, 0)
		used = append(used, false)
		sum += n
	}

	target := sum / part
	lqe = 500000000000
	found = false

	for i := 2; !found && i <= len(all_nums); i++ {
		bkt(1, i, target)
	}

	return lqe
}

func part_two(input string) int {
	return 0
}

func main() {
	fcontent, err := ioutil.ReadFile("balance24.in")
	if err != nil {
		fmt.Println("Error reading file", err)
	} else {
		input := string(fcontent)

		fmt.Println(solve(input, 3))
		fmt.Println(solve(input, 4))
	}
}
