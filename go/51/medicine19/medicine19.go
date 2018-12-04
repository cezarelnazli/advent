package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"unicode"
)

func one_round(start string, transforms *map[string][]string) map[string]bool {
	combos := make(map[string]bool)

	for i := 0; i < len(start); i++ {
		j := i + 1
		for j < len(start) && unicode.IsLower(rune(start[j])) {
			j++
		}

		aux := string(start[i:j])
		transform, ok := (*transforms)[aux]
		if ok {
			for _, t := range transform {
				combo := make([]byte, 0)
				combo = append(combo, start[:i]...)
				combo = append(combo, t...)
				if j < len(start) {
					combo = append(combo, start[j:]...)
				}

				combos[string(combo)] = true
			}
		}
	}

	return combos
}

func part_one(input string) int {
	transforms := make(map[string][]string)
	var start string

	for _, l := range strings.Split(input, "\n") {
		if len(l) > 1 {
			transform := strings.Split(l, " => ")
			if len(transform) == 2 {
				transforms[transform[0]] = append(transforms[transform[0]], transform[1])
			} else {
				aux := make([]byte, len(transform[0]))
				copy(aux, []byte(transform[0]))
				start = string(aux)
			}
		}
	}

	return len(one_round(start, &transforms))
}

func part_two(input string) int {
	transforms := make(map[string]string)
	var start string

	for _, l := range strings.Split(input, "\n") {
		if len(l) > 1 {
			start = l
			transform := strings.Split(l, " => ")
			if len(transform) == 2 {
				transforms[transform[1]] = transform[0]
			}
		}
	}

	step := 1
	perform := true
	for perform {
		if transforms[start] == "e" {
			break
		}

		keys := make([]string, 0)
		for k, _ := range transforms {
			keys = append(keys, k)
		}

		sort.Slice(keys, func(i, j int) bool { return len(keys[i]) > len(keys[j]) })

		for _, k := range keys {
			aux := strings.Replace(start, k, transforms[k], 1)
			if aux != start {
				step++
				start = aux
				break
			}
		}
	}

	return step
}

func main() {
	fcontent, err := ioutil.ReadFile("medicine19.in")
	if err != nil {
		fmt.Println("error reading file", err)
	} else {
		input := string(fcontent)

		fmt.Println(part_one(input))
		fmt.Println(part_two(input))
	}
}
