package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pair struct {
	s string
	i int
}

func part_one(input string) int {
	nice := 0

	for _, s := range strings.Split(input, "\n") {
		vowels := 0
		double := false
		naughty := false

		for i := 0; i < len(s)-1; i++ {
			switch string(s[i : i+2]) {
			case "ab", "cd", "pq", "xy":
				naughty = true
				break
			}

			switch s[i] {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
				break
			}

			if s[i] == s[i+1] {
				double = true
			}
		}

		switch s[len(s)-1] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
			break
		}

		if vowels >= 3 && double && !naughty {
			nice++
		}
	}

	return nice
}

func part_two(input string) int {
	nice := 0

	for _, s := range strings.Split(input, "\n") {
		double := false
		groups := make([]Pair, 0)

		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				double = true
			}
		}

		found := false
		for i := 0; !found && i < len(s)-1; i++ {
			for j := 0; j < len(groups); j++ {
				if groups[j].s == string(s[i:i+2]) && groups[j].i+1 < i {
					found = true
					break
				}
			}

			groups = append(groups, Pair{s: string(s[i : i+2]), i: i})
		}

		if double && found {
			nice++
		}
	}

	return nice
}

func main() {
	fcontent, err := ioutil.ReadFile("intern05.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
