package main

import (
	"fmt"
)

const STEPS int = 12261543

func part_one() int {
	tape := make([]int, STEPS*2)
	to_write := [][2]int{
		'A': {1, 0},
		'C': {1, 0}}
	to_move := [][2]int{
		'A': {+1, -1},
		'B': {-1, +1},
		'C': {+1, -1},
		'D': {-1, -1},
		'E': {+1, +1},
		'F': {+1, +1}}
	next_state := [][2]byte{
		'A': {'B', 'C'},
		'B': {'A', 'C'},
		'C': {'A', 'D'},
		'D': {'E', 'C'},
		'E': {'F', 'A'},
		'F': {'A', 'E'}}

	cursor := STEPS
	state := byte('A')
	ones := 0

	for i := 0; i < STEPS; i++ {
		old_tape := tape[cursor]

		if state == 'A' || state == 'C' {
			if tape[cursor] == 1 && to_write[state][tape[cursor]] == 0 {
				ones--
			} else if tape[cursor] == 0 && to_write[state][tape[cursor]] == 1 {
				ones++
			}

			tape[cursor] = to_write[state][tape[cursor]]
		} else {
			if tape[cursor] == 0 {
				ones++
			}

			tape[cursor] = 1
		}

		new_cursor := cursor
		new_cursor += to_move[state][old_tape]

		state = next_state[state][old_tape]

		cursor = new_cursor
	}

	return ones
}

func main() {
	fmt.Println(part_one())
}
