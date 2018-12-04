package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const (
	UP    int = 0
	RIGHT int = 1
	DOWN  int = 2
	LEFT  int = 3
)

func make_matrix(input string) [][]byte {
	matrix := make([][]byte, 0)

	for _, l := range strings.Split(input, "\n") {
		row := make([]byte, 0)

		for _, c := range l {
			row = append(row, byte(c))
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func part_one(matrix [][]byte) string {
	r, c := 0, 0
	result := ""

	for i, e := range matrix[0] {
		if e == '|' {
			c = i
			break
		}
	}

	dir := DOWN

	loop := true
	steps := 0
	for loop {
		for matrix[r][c] != '+' && matrix[r][c] != ' ' {
			steps++
			if unicode.IsLetter(rune(matrix[r][c])) {
				result += string(matrix[r][c])
			}

			switch dir {
			case DOWN:
				r++
				break
			case UP:
				r--
				break
			case LEFT:
				c--
				break
			case RIGHT:
				c++
				break
			}
		}

		switch dir {
		case UP, DOWN:
			steps++
			if c > 0 && matrix[r][c-1] == '-' {
				dir = LEFT
				c--
			} else if c < len(matrix[r])-1 && matrix[r][c+1] == '-' {
				dir = RIGHT
				c++
			} else {
				loop = false
				steps--
			}
			break
		case LEFT, RIGHT:
			steps++
			if r > 0 && matrix[r-1][c] == '|' {
				dir = UP
				r--
			} else if r < len(matrix)-1 && matrix[r+1][c] == '|' {
				dir = DOWN
				r++
			} else {
				loop = false
				steps--
			}
			break
		}
	}

	fmt.Println(steps)
	return result
}

func main() {
	fcontent, err := ioutil.ReadFile("tubes.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	matrix := make_matrix(string(fcontent))
	fmt.Println(part_one(matrix))
}
