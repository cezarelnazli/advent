package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string, start int) int {
	a, b := start, 0
	code := make([]string, 0)

	for _, instr := range strings.Split(input, "\n") {
		code = append(code, instr)
	}

	for i := 0; i < len(code); i++ {
		switch string(code[i][:3]) {
		case "hlf":
			if code[i][len(code[i])-1] == 'a' {
				a >>= 1
			} else {
				b >>= 1
			}
			break
		case "tpl":
			if code[i][len(code[i])-1] == 'a' {
				a *= 3
			} else {
				b *= 3
			}
			break
		case "inc":
			if code[i][len(code[i])-1] == 'a' {
				a++
			} else {
				b++
			}
			break
		case "jmp":
			var jmp int
			fmt.Sscanf(code[i], "jmp %d", &jmp)
			i += jmp - 1
			break
		case "jie":
			var jmp, offset int
			var r rune
			fmt.Sscanf(code[i], "jie %c, %d", &r, &jmp)
			if jmp < 0 {
				offset = -1
			} else {
				offset = +1
			}
			if r == 'a' && a%2 == 0 {
				i += jmp - offset
			} else if r == 'b' && b%2 == 0 {
				i += jmp - offset
			}
			break
		case "jio":
			var jmp, offset int
			var r rune
			fmt.Sscanf(code[i], "jio %c, %d", &r, &jmp)
			if jmp < 0 {
				offset = -1
			} else {
				offset = +1
			}
			if r == 'a' && a == 1 {
				i += jmp - offset
			} else if r == 'b' && b == 1 {
				i += jmp - offset
			}
			break
		}
	}

	return b
}

func main() {
	fcontent, err := ioutil.ReadFile("turing23.in")
	if err != nil {
		fmt.Println("error opening file", err)
	} else {
		input := string(fcontent)
		fmt.Println(part_one(input, 0))
		fmt.Println(part_one(input, 1))
	}
}
