package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part_one(input string) int {
	instructions := make([]int, 0)
	steps := 1

	for _, s := range strings.Split(input, "\n") {
		instr, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			fmt.Println("cannot parse int")
			continue
		}

		instructions = append(instructions, int(instr))
	}

	ninstr := len(instructions)

	cpos := 0
	for {
		if cpos < 0 || cpos >= ninstr {
			break
		}

		instructions[cpos] += 1
		cpos += instructions[cpos] - 1
		steps++
	}

	return steps
}

func part_two(input string) int {
	instructions := make([]int, 0)
	steps := 0

	for _, s := range strings.Split(input, "\n") {
		instr, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			fmt.Println("cannot parse int")
			continue
		}

		instructions = append(instructions, int(instr))
	}

	ninstr := len(instructions)

	cpos := 0
	for {
		if cpos < 0 || cpos >= ninstr {
			break
		}

		if instructions[cpos] >= 3 {
			instructions[cpos] -= 1
			cpos += instructions[cpos] + 1
		} else {
			instructions[cpos] += 1
			cpos += instructions[cpos] - 1
		}

		steps++
	}

	return steps
}

func main() {
	fcontent, err := ioutil.ReadFile("maze.in")
	if err != nil {
		fmt.Println("Error reading input")
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
