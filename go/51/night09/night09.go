package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var used [20]bool
var sol [20]string
var cmin int
var distances map[string]map[string]int
var cities []string

func add_city(new_city string) {
	found := false
	for _, c := range cities {
		if c == new_city {
			found = true
		}
	}

	if !found {
		cities = append(cities, new_city)
	}
}

func bkt(level, ncities int) {
	if level == ncities {
		current := 0
		for i := 0; i < len(sol)-1; i++ {
			current += distances[sol[i]][sol[i+1]]
		}

		if current > cmin {
			cmin = current
		}
	} else {
		for i := 0; i < ncities; i++ {
			if !used[i] {
				sol[level] = cities[i]
				used[i] = true
				bkt(level+1, ncities)
				used[i] = false
			}
		}
	}
}

func part_one(input string) int {
	distances = make(map[string]map[string]int)
	cities = make([]string, 0)

	for _, l := range strings.Split(input, "\n") {
		var city1, city2 string
		var distance int
		fmt.Sscanf(l, "%s to %s = %d", &city1, &city2, &distance)

		_, ok := distances[city1]
		if !ok {
			distances[city1] = make(map[string]int)
		}
		distances[city1][city2] = distance

		_, ok = distances[city2]
		if !ok {
			distances[city2] = make(map[string]int)
		}
		distances[city2][city1] = distance

		add_city(city1)
		add_city(city2)
	}

	bkt(0, len(cities))

	return cmin
}

func main() {
	fcontent, err := ioutil.ReadFile("night09.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
