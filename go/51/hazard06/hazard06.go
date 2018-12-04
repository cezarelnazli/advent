package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const (
	OFF int = -1
	ON  int = 1
	TOG int = 2
)

const SIZE int = 1000

func part_one(input string) int {
	grid := make([][SIZE]int, SIZE)

	for _, instr := range strings.Split(input, "\n") {
		comp := strings.Split(instr, " ")
		var op, bi, bj, ei, ej int

		if comp[0] == "turn" {
			op = OFF
			if comp[1] == "on" {
				op = ON
			}

			fmt.Sscanf(comp[2], "%d,%d", &bi, &bj)
			fmt.Sscanf(comp[4], "%d,%d", &ei, &ej)
		} else if comp[0] == "toggle" {
			op = TOG

			fmt.Sscanf(comp[1], "%d,%d", &bi, &bj)
			fmt.Sscanf(comp[3], "%d,%d", &ei, &ej)
		}

		for i := bi; i <= ei; i++ {
			for j := bj; j <= ej; j++ {
				grid[i][j] = int(math.Max(0, float64(grid[i][j]+op)))
			}
		}
	}

	on := 0
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			on += grid[i][j]
		}
	}

	return on
}

func main() {
	fcontent, err := ioutil.ReadFile("hazard06.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
