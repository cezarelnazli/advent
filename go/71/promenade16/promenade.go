package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part_one(start, input string) string {
	result := make([]byte, 16)
	copy(result, start)

	for _, cmd := range strings.Split(input, ",") {
		switch cmd[0] {
		case 's':
			num64, err := strconv.ParseInt(cmd[1:], 0, 0)
			if err != nil {
				fmt.Println("cannot convert int", err)
				continue
			}

			aux := make([]byte, 16)
			copy(aux, result[0:len(result)-int(num64)])
			for i := 0; i < int(num64); i++ {
				result[i] = result[len(result)-int(num64)+i]
			}

			for i := 0; i < len(result)-int(num64); i++ {
				result[int(num64)+i] = aux[i]
			}
		case 'x':
			var posA, posB int
			fmt.Sscanf(cmd, "x%d/%d", &posA, &posB)
			aux := result[posA]
			result[posA] = result[posB]
			result[posB] = aux
		case 'p':
			var progA, progB byte
			var posA, posB int
			progA = cmd[1]
			progB = cmd[3]

			for i, b := range result {
				if b == progA {
					posA = i
				}

				if b == progB {
					posB = i
				}
			}

			aux := result[posA]
			result[posA] = result[posB]
			result[posB] = aux
		}
	}

	return string(result)
}

func part_two(input string) string {
	result := "abcdefghijklmnop"

	for i := 1; i <= 1000000000%63; i++ {
		result = part_one(result, input)
	}

	return result
}

func main() {
	fcontent, err := ioutil.ReadFile("promenade.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one("abcdefghijklmnop", input))
	fmt.Println(part_two(input))
}
