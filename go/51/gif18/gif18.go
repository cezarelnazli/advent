package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const STEPS int = 100

type Dir struct {
	x int
	y int
}

var DIRS [8]Dir = [8]Dir{
	{x: -1, y: -1}, {x: 0, y: -1}, {x: +1, y: -1}, {x: -1, y: 0},
	{x: +1, y: 0}, {x: -1, y: 1}, {x: 0, y: 1}, {x: +1, y: +1}}

func part_one(input string, steps int) int {
	grid := make([][]bool, 0)
	grid_new := make([][]bool, 0)

	for i, l := range strings.Split(input, "\n") {
		grid = append(grid, make([]bool, len(l), len(l)))
		grid_new = append(grid_new, make([]bool, len(l), len(l)))

		for j, c := range l {
			if c == '#' {
				grid[i][j] = true
			}
		}
	}

	for s := 0; s < steps; s++ {
		for y, l := range grid {
			for x, _ := range l {
				adj := 0
				for _, d := range DIRS {
					if x+d.x >= 0 && x+d.x < len(l) &&
						y+d.y >= 0 && y+d.y < len(grid) {
						if grid[y+d.y][x+d.x] {
							adj++
						}
					}
				}

				if adj == 3 {
					grid_new[y][x] = true
				} else if grid[y][x] == true && adj == 2 {
					grid_new[y][x] = true
				} else {
					grid_new[y][x] = false
				}
			}
		}

		for y, l := range grid_new {
			for x, c := range l {
				grid[y][x] = c
			}
		}
	}

	num := 0
	for _, l := range grid {
		for _, c := range l {
			if c {
				num++
			}
		}
	}

	return num
}

func part_two(input string, steps int) int {
	grid := make([][]bool, 0)
	grid_new := make([][]bool, 0)

	for i, l := range strings.Split(input, "\n") {
		grid = append(grid, make([]bool, len(l), len(l)))
		grid_new = append(grid_new, make([]bool, len(l), len(l)))

		for j, c := range l {
			if c == '#' {
				grid[i][j] = true
			}
		}
	}

	grid[0][0] = true
	grid[0][len(grid)-1] = true
	grid[len(grid)-1][len(grid)-1] = true
	grid[len(grid)-1][0] = true

	for s := 0; s < steps; s++ {
		for y, l := range grid {
			for x, _ := range l {
				adj := 0
				for _, d := range DIRS {
					if x+d.x >= 0 && x+d.x < len(l) &&
						y+d.y >= 0 && y+d.y < len(grid) {
						if grid[y+d.y][x+d.x] {
							adj++
						}
					}
				}

				if adj == 3 || (x == 0 && y == len(l)-1) || (x == 0 && y == 0) || (x == len(l)-1 && y == 0) || (x == len(l)-1 && y == len(l)-1) {
					grid_new[y][x] = true
				} else if grid[y][x] == true && adj == 2 {
					grid_new[y][x] = true
				} else {
					grid_new[y][x] = false
				}
			}
		}

		for y, l := range grid_new {
			for x, c := range l {
				grid[y][x] = c
			}
		}
	}

	num := 0
	for _, l := range grid {
		for _, c := range l {
			if c {
				num++
			}
		}
	}

	return num
}

func main() {
	fcontent, err := ioutil.ReadFile("gif18.in")
	if err != nil {
		fmt.Println("error reading file", err)
	} else {
		input := string(fcontent)

		fmt.Println(part_one(input, STEPS))
		fmt.Println(part_two(input, STEPS))
	}
}
