package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string) int {
	total := 0

	for _, l := range strings.Split(input, "\n") {
		llen := len(l)

		escapes := 0
		for i := 0; i < llen; i++ {
			if l[i] == '\\' {
				i++
				escapes++
				if l[i] == 'x' {
					i += 2
				}
			} else if l[i] != '"' {
				escapes++
			}
		}

		total += llen - escapes
	}

	return total
}

func part_two(input string) int {
	diff := 0

	for _, l := range strings.Split(input, "\n") {
		diff += 2
		for i := 0; i < len(l); i++ {
			if l[i] == '\\' || l[i] == '"' {
				diff++
			}
		}
	}

	return diff
}

func main() {
	fcontent, err := ioutil.ReadFile("matchsticks08.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
