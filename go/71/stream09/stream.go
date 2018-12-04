package main

import (
	"fmt"
	"io/ioutil"
)

func part_one(input string) (int, int) {
	level := 0
	sum := 0
	garbage := 0
	ignore_next := false
	in_garbage := false

	for _, c := range input {
		if ignore_next {
			ignore_next = false
			continue
		}

		if in_garbage {
			switch c {
			case '!':
				ignore_next = true
			case '>':
				in_garbage = false
			default:
				garbage++
			}

			continue
		}

		switch c {
		case '!':
			ignore_next = true
		case '<':
			in_garbage = true
		case '{':
			level++
			sum += level
		case '}':
			level--
		}
	}

	return sum, garbage
}

func main() {
	fcontent, err := ioutil.ReadFile("stream.in")
	if err != nil {
		fmt.Println("read error: ", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
