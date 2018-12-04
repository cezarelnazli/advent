package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Component struct {
	side1 int
	side2 int
}

var components []Component
var is_used []bool
var crt, max, max_level int

func backtrack(level, head int) {
	for i, c := range components {
		if (c.side1 == head || c.side2 == head) && !is_used[i] {
			saved_head := head
			crt += c.side1 + c.side2
			if crt > max && level >= max_level {
				max = crt
				max_level = level
			}
			is_used[i] = true

			if c.side1 == head {
				backtrack(level+1, c.side2)
			} else {
				backtrack(level+1, c.side1)
			}

			crt -= c.side1 + c.side2
			head = saved_head
			is_used[i] = false
		}
	}
}

func part_one(input string) int {
	components = make([]Component, 0)

	for _, l := range strings.Split(input, "\n") {
		var side1, side2 int
		fmt.Sscanf(l, "%d/%d", &side1, &side2)
		components = append(components, Component{side1: side1, side2: side2})
	}

	is_used = make([]bool, len(components))

	backtrack(0, 0)

	return max
}

func part_two(input string) int {
	return 0
}

func main() {
	fcontent, err := ioutil.ReadFile("moat.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
