package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	SIZE    int = 20000
	BURSTS  int = 10000
	BURSTS2 int = 10000000
)

const (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

const (
	CLEAN    int = 0
	WEAK     int = 1
	INFECTED int = 2
	FLAGGED  int = 3
)

func part_one(input string) int {
	is_infected := make([][]bool, SIZE)
	infected := 0
	lines, cols := 0, 0

	for i := 0; i < SIZE; i++ {
		is_infected[i] = make([]bool, SIZE)
	}

	for i, line := range strings.Split(input, "\n") {
		lines = i

		for j, c := range line {
			cols = j
			if c == '#' {
				is_infected[i+BURSTS][j+BURSTS] = true
			}
		}
	}

	i := BURSTS + (lines / 2)
	j := BURSTS + (cols / 2)
	dir := UP
	for b := 0; b < BURSTS; b++ {
		if is_infected[i][j] {
			dir = (dir + 1) % 4
		} else {
			if dir == UP {
				dir = LEFT
			} else {
				dir -= 1
			}
		}

		is_infected[i][j] = !is_infected[i][j]
		if is_infected[i][j] {
			infected++
		}

		if dir == UP {
			i--
		} else if dir == RIGHT {
			j++
		} else if dir == DOWN {
			i++
		} else if dir == LEFT {
			j--
		}
	}

	return infected
}

func part_two(input string) int {
	state := make([][]int, SIZE)
	infected := 0
	lines, cols := 0, 0

	for i := 0; i < SIZE; i++ {
		state[i] = make([]int, SIZE)
	}

	for i, line := range strings.Split(input, "\n") {
		lines = i

		for j, c := range line {
			cols = j
			if c == '#' {
				state[i+BURSTS][j+BURSTS] = INFECTED
			}
		}
	}

	i := BURSTS + (lines / 2)
	j := BURSTS + (cols / 2)
	dir := UP
	for b := 0; b < BURSTS2; b++ {
		if state[i][j] == CLEAN {
			if dir == UP {
				dir = LEFT
			} else {
				dir -= 1
			}
		} else if state[i][j] == INFECTED {
			dir = (dir + 1) % 4
		} else if state[i][j] == FLAGGED {
			if dir == UP {
				dir = DOWN
			} else if dir == DOWN {
				dir = UP
			} else if dir == LEFT {
				dir = RIGHT
			} else if dir == RIGHT {
				dir = LEFT
			}
		}

		if state[i][j] == CLEAN {
			state[i][j] = WEAK
		} else if state[i][j] == WEAK {
			state[i][j] = INFECTED
			infected++
		} else if state[i][j] == INFECTED {
			state[i][j] = FLAGGED
		} else if state[i][j] == FLAGGED {
			state[i][j] = CLEAN
		}

		if dir == UP {
			i--
		} else if dir == RIGHT {
			j++
		} else if dir == DOWN {
			i++
		} else if dir == LEFT {
			j--
		}
	}

	return infected
}

func main() {
	fcontent, err := ioutil.ReadFile("sporifica.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
