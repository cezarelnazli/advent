package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string) int {
	n, ne, se, s, sw, nw := 0, 0, 0, 0, 0, 0

	for _, str := range strings.Split(strings.TrimSpace(input), ",") {
		switch str {
		case "n":
			n++
		case "s":
			s++
		case "ne":
			ne++
		case "se":
			se++
		case "nw":
			nw++
		case "sw":
			sw++
		}
	}

	res := 0

	if n < s {
		s -= n
		n = 0
	} else {
		n -= s
		s = 0
	}

	if ne < sw {
		sw -= ne
		ne = 0
	} else {
		ne -= sw
		sw = 0
	}

	if nw < se {
		se -= nw
		nw = 0
	} else {
		nw -= se
		se = 0
	}

	if se < n {
		res += se
		n -= se
		se = 0
	} else {
		res += n
		se -= n
		n = 0
	}

	if sw < n {
		res += sw
		n -= sw
		sw = 0
	} else {
		res += n
		sw -= n
		n = 0
	}

	if nw < s {
		res += nw
		s -= nw
		nw = 0
	} else {
		res += s
		nw -= s
		s = 0
	}

	if ne < s {
		res += ne
		s -= ne
		ne = 0
	} else {
		res += s
		ne -= s
		s = 0
	}

	if sw < se {
		res += sw
		se -= sw
		sw = 0
	} else {
		res += se
		sw -= se
		se = 0
	}

	if nw < ne {
		res += nw
		ne -= nw
		nw = 0
	} else {
		res += ne
		nw -= ne
		ne = 0
	}

	res += n + ne + nw + s + se + sw

	return res
}

func part_two(input string) int {
	max := 0

	for i, s := range input {
		if s == ',' {
			r := part_one(input[:i])
			if r > max {
				max = r
			}
		}
	}

	return max
}

func main() {
	fcontent, err := ioutil.ReadFile("hexes.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
