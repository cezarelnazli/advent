package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const SIZE int = 2000

func part_one(input string) int {
	adjecency := make([][SIZE]bool, SIZE)
	nlinked := 0

	for _, ln := range strings.Split(input, "\n") {
		ln_split := strings.Split(ln, " <-> ")
		left, right := ln_split[0], ln_split[1]
		right_elems := strings.Split(right, ", ")

		lnode64, err := strconv.ParseInt(left, 0, 0)
		if err != nil {
			fmt.Println("error parsing lnode", err)
			continue
		}
		lnode := int(lnode64)

		for _, e := range right_elems {
			rnode64, err := strconv.ParseInt(e, 0, 0)
			if err != nil {
				fmt.Println("error parsing rnode", err)
				continue
			}
			rnode := int(rnode64)

			adjecency[lnode][rnode] = true
			adjecency[rnode][lnode] = true
			adjecency[rnode][rnode] = true
			adjecency[lnode][lnode] = true
		}
	}

	visited := make([]bool, SIZE)
	queue := make([]int, SIZE*SIZE)
	connected := 0

	for start := 0; start < SIZE; start++ {
		if !visited[start] {
			connected++
			head := 0
			tail := 0

			queue[tail] = start
			tail++

			for head < tail {
				elem := queue[head]
				head++

				for i := 0; i < SIZE; i++ {
					if adjecency[elem][i] && !visited[i] {
						visited[i] = true
						queue[tail] = i
						tail++
					}
				}

				visited[elem] = true
			}
		}

		if start == 0 {
			for i := 0; i < SIZE; i++ {
				if visited[i] {
					nlinked++
				}
			}

			fmt.Println(nlinked)
		}
	}

	return connected
}

func part_two(input string) int {
	return 0
}

func main() {
	fcontent, err := ioutil.ReadFile("salesman.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
