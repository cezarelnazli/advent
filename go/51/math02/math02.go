package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part_one(input string) (int, int) {
	wrap := 0
	ribbon := 0

	for _, box := range strings.Split(input, "\n") {
		sides := strings.Split(box, "x")

		aux, err := strconv.ParseInt(sides[0], 0, 0)
		if err != nil {
			fmt.Println("cannot convert", sides[0])
		}
		l := int(aux)

		aux, err = strconv.ParseInt(sides[1], 0, 0)
		if err != nil {
			fmt.Println("cannot convert", sides[1])
		}
		w := int(aux)

		aux, err = strconv.ParseInt(sides[2], 0, 0)
		if err != nil {
			fmt.Println("cannot convert", sides[2])
		}
		h := int(aux)

		min := l * w
		if l*h < min {
			min = l * h
		}

		if w*h < min {
			min = w * h
		}

		wrap += min + 2*l*w + 2*w*h + 2*h*l

		ribbon += l * w * h
		if l <= w && w <= h {
			ribbon += 2*l + 2*w
		} else if l <= h && h <= w {
			ribbon += 2*l + 2*h
		} else if w <= l && l <= h {
			ribbon += 2*w + 2*l
		} else if w <= h && h <= l {
			ribbon += 2*w + 2*h
		} else if h <= l && l <= w {
			ribbon += 2*h + 2*l
		} else if h <= w && w <= l {
			ribbon += 2*h + 2*w
		}
	}

	return ribbon, wrap
}

func main() {
	fcontent, err := ioutil.ReadFile("math02.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
