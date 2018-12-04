package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part_one(input string) int {
	steps := 0

	banks := make([]int, len(strings.Fields(input)))
	banks_cache := make([][]int, 0)

	for i, bsize_str := range strings.Fields(input) {
		bsize, err := strconv.ParseInt(bsize_str, 0, 0)
		if err != nil {
			fmt.Println("cannot parse int", bsize_str)
		}

		banks[i] = int(bsize)
	}

	for {
		banks_copy := make([]int, len(banks))
		_ = copy(banks_copy, banks)
		banks_cache = append(banks_cache, banks_copy)

		emax := 0
		imax := 0
		for i, bsize := range banks {
			if bsize > emax {
				emax = bsize
				imax = i
			}
		}

		banks[imax] = 0
		for i := 0; i < len(banks); i++ {
			banks[i] += emax / len(banks)
			if (i-imax > 0 && i-imax <= emax%len(banks)) ||
				(i-imax <= 0 && i <= emax%len(banks)-(len(banks)-imax)) {
				banks[i]++
			}
		}

		exists := false
		var i int
		for i = 0; !exists && i < len(banks_cache); i++ {
			exists = true
			for j := 0; exists && j < len(banks); j++ {
				if banks[j] != banks_cache[i][j] {
					exists = false
				}
			}
		}

		steps++
		if exists {
			fmt.Println(len(banks_cache) - i)
			break
		}
	}

	return steps
}

func part_two(input string) int {
	return 0
}

func main() {
	fcontent, err := ioutil.ReadFile("memory.in")
	if err != nil {
		fmt.Println("error reading file")
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
