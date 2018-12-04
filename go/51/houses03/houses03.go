package main

import (
	"fmt"
	"io/ioutil"
)

const SIZE int = 1000

func part_one(input string) int {
	delivered := make([][2 * SIZE]bool, 2*SIZE)
	result := 0

	si, sj := SIZE, SIZE
	ri, rj := SIZE, SIZE
	for s, c := range input {
		if !delivered[si][sj] {
			result++
		}

		delivered[si][sj] = true

		if !delivered[ri][rj] {
			result++
		}

		delivered[ri][rj] = true

		if s%2 == 0 {
			if c == 'v' {
				si++
			} else if c == '^' {
				si--
			} else if c == '<' {
				sj--
			} else if c == '>' {
				sj++
			}
		} else {
			if c == 'v' {
				ri++
			} else if c == '^' {
				ri--
			} else if c == '<' {
				rj--
			} else if c == '>' {
				rj++
			}
		}
	}

	return result
}

func main() {
	fcontent, err := ioutil.ReadFile("houses03.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
