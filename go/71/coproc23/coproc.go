package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func getInt(elem string, regs *map[string]int) int {
	result := 0

	if unicode.IsLetter(rune(elem[0])) {
		result = (*regs)[elem]
	} else {
		int64, err := strconv.ParseInt(elem, 0, 0)
		if err != nil {
			fmt.Println("error converting int", err)
		}

		result = int(int64)
	}

	return result
}

func part_one(input string) int {
	instrs := make([]string, 0)
	regs := make(map[string]int)

	for _, l := range strings.Split(input, "\n") {
		instrs = append(instrs, l)
	}

	for _, l := range "abcdefgh" {
		regs[string(l)] = 0
	}

	nmul := 0

	for i := 0; i < len(instrs); i++ {
		elems := strings.Fields(instrs[i])

		switch elems[0] {
		case "set":
			val := getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "sub":
			val := regs[elems[1]]
			val -= getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "mul":
			val := regs[elems[1]]
			val *= getInt(elems[2], &regs)
			regs[elems[1]] = val
			nmul++
			break
		case "jnz":
			x := getInt(elems[1], &regs)
			if x != 0 {
				i += getInt(elems[2], &regs)
				if i < 0 || i >= len(instrs) {
					return nmul
				}

				i--
			}
			break
		}
	}

	return nmul
}

func part_two() int {
	res := 0

	for i := 107900; i <= 124900; i += 17 {
		for d := 2; d < i; d++ {
			if i%d == 0 {
				res++
				break
			}
		}
	}

	return res
}

func main() {
	fcontent, err := ioutil.ReadFile("coproc.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two())
}
