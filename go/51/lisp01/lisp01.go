package main

import (
	"fmt"
	"io/ioutil"
)

func part_one(input string) int {
	floor := 0
	basement := false

	for i, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
			if floor == -1 && !basement {
				fmt.Println(i + 1)
				basement = true
			}
		}
	}

	return floor
}

func main() {
	fcontent, err := ioutil.ReadFile("lisp01.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
