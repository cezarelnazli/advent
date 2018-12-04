package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func getInt(input string, regs *map[string]uint16) (uint16, bool) {
	if unicode.IsDigit(rune(input[0])) {
		aux, err := strconv.ParseInt(input, 0, 16)
		if err != nil {
			return 0, false
		} else {
			return uint16(aux), true
		}
	} else {
		aux, ok := (*regs)[input]
		if ok {
			return aux, true
		} else {
			return 0, false
		}
	}
}

func part_one(input string) uint16 {
	regs := make(map[string]uint16)
	lines := strings.Split(input, "\n")

	for {
		change := false

		for _, l := range lines {
			sides := strings.Split(l, " -> ")
			lhs, rhs := sides[0], sides[1]

			lhs_items := strings.Fields(lhs)

			if rhs == "b" {
				regs["b"] = 46065
				continue
			}

			if len(lhs_items) == 1 {
				aux, ok := getInt(lhs_items[0], &regs)
				if ok {
					regs[rhs] = aux
				} else {
					change = true
				}
			} else if len(lhs_items) == 2 {
				aux, ok := getInt(lhs_items[1], &regs)
				if ok {
					regs[rhs] = ^aux
				} else {
					change = true
				}
			} else {
				aux1, ok := getInt(lhs_items[0], &regs)
				if ok {
					aux2, ok := getInt(lhs_items[2], &regs)
					if ok {
						switch lhs_items[1] {
						case "AND":
							regs[rhs] = aux1 & aux2
							break
						case "OR":
							regs[rhs] = aux1 | aux2
							break
						case "RSHIFT":
							regs[rhs] = aux1 >> aux2
							break
						case "LSHIFT":
							regs[rhs] = aux1 << aux2
							break
						}
					} else {
						change = true
					}
				} else {
					change = true
				}
			}
		}

		if !change {
			break
		}
	}

	return regs["a"]
}

func main() {
	fcontent, err := ioutil.ReadFile("assembly07.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
}
