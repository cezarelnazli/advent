package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var names []string
var gain map[string]map[string]int
var cmax int
var sol [9]int
var used [9]bool

func add_person(person string) {
	found := false
	for _, n := range names {
		if n == person {
			found = true
			break
		}
	}

	if !found {
		names = append(names, person)
	}
}

func bkt(level int) {
	if level == len(names) {
		sum := gain[names[sol[0]]][names[sol[1]]]
		sum += gain[names[sol[0]]][names[sol[len(names)-1]]]
		sum += gain[names[sol[len(names)-1]]][names[sol[0]]]
		sum += gain[names[sol[len(names)-1]]][names[sol[len(names)-2]]]
		for i := 1; i < len(names)-1; i++ {
			sum += gain[names[sol[i]]][names[sol[i-1]]]
			sum += gain[names[sol[i]]][names[sol[i+1]]]
		}

		if sum > cmax {
			cmax = sum
		}
	} else {
		for i := 0; i < len(names); i++ {
			if !used[i] {
				used[i] = true
				sol[level] = i
				bkt(level + 1)
				used[i] = false
			}
		}
	}
}

func part_one(input string) int {
	gain = make(map[string]map[string]int)
	names = make([]string, 0)

	for _, l := range strings.Split(input, "\n") {
		var p1, p2, op string
		var amount int

		fmt.Sscanf(l,
			"%s would %s %d happiness units by sitting next to %s.",
			&p1, &op, &amount, &p2)

		add_person(p1)
		add_person(p2[:len(p2)-1])
		add_person("Me")

		_, ok := gain[p1]
		if !ok {
			gain[p1] = make(map[string]int)
		}

		if op == "gain" {
			gain[p1][p2[:len(p2)-1]] = amount
		} else {
			gain[p1][p2[:len(p2)-1]] = amount * -1
		}

		gain[p1]["Me"] = 0
	}

	gain["Me"] = make(map[string]int)

	bkt(0)

	return cmax
}

func part_two(input string) int {
	return 0
}

func main() {
	fcontent, err := ioutil.ReadFile("knights13.in")
	if err != nil {
		fmt.Println("error reading file", err)
	} else {
		input := string(fcontent)

		fmt.Println(part_one(input))
		fmt.Println(part_two(input))
	}
}
