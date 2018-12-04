package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string) int {
	depths := make([]int, 100)
	last := 0
	sum := 0

	for _, info := range strings.Split(input, "\n") {
		var level, depth int
		fmt.Sscanf(info, "%d: %d", &level, &depth)
		depths[level] = depth
		last = level
	}

	for i := 0; i <= last; i++ {
		if depths[i] > 0 {
			if i/(depths[i]-1)%2 == 0 {
				if i%(depths[i]-1) == 0 {
					sum += i * depths[i]
				}
			} else {
				if depths[i]-1-(i%depths[i]-1) == 0 {
					sum += i * depths[i]
				}
			}
		}
	}

	return sum
}

func part_two(input string) int {
	depths := make([]int, 100)
	last := 0

	for _, info := range strings.Split(input, "\n") {
		var level, depth int
		fmt.Sscanf(info, "%d: %d", &level, &depth)
		depths[level] = depth
		last = level
	}

	delay := 0
	for {
		hit := false
		for i := 0; i <= last; i++ {
			step := i + delay
			if depths[i] > 0 {
				if step/(depths[i]-1)%2 == 0 {
					if step%(depths[i]-1) == 0 {
						hit = true
						break
					}
				} else {
					if depths[i]-1-(step%depths[i]-1) == 0 {
						hit = true
						break
					}
				}
			}
		}

		if !hit {
			return delay
		}

		delay++
	}
}

func main() {
	fcontent, err := ioutil.ReadFile("firewall.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
