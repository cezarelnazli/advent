package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string) int {
	registers := make(map[string]int)
	all_max := 0

	for _, instr := range strings.Split(input, "\n") {
		var register, op, c_register, c_op string
		var val, c_val int

		fmt.Sscanf(instr, "%s %s %d if %s %s %d",
			&register, &op, &val,
			&c_register, &c_op, &c_val)

		do_op := false

		switch c_op {
		case ">":
			if registers[c_register] > c_val {
				do_op = true
			}
		case "<":
			if registers[c_register] < c_val {
				do_op = true
			}
		case "<=":
			if registers[c_register] <= c_val {
				do_op = true
			}
		case ">=":
			if registers[c_register] >= c_val {
				do_op = true
			}
		case "==":
			if registers[c_register] == c_val {
				do_op = true
			}
		case "!=":
			if registers[c_register] != c_val {
				do_op = true
			}
		}

		if do_op {
			switch op {
			case "inc":
				registers[register] += val
			case "dec":
				registers[register] -= val
			}

			if registers[register] > all_max {
				all_max = registers[register]
			}
		}
	}

	is_set := false
	max := 0
	for _, v := range registers {
		if !is_set || v > max {
			is_set = true
			max = v
		}
	}

	fmt.Println(all_max)
	return max
}

func main() {
	fcontent, err := ioutil.ReadFile("registers.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
