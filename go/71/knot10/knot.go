package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const ListSize int = 256

type Square struct {
	i int
	j int
}

func one_round(lengths []byte, cpos *int, skip *int, list *[]int) {
	for _, length := range lengths {
		aux_list := make([]int, ListSize)

		limit := *cpos + int(length)
		limit_first := 0
		if limit > ListSize {
			limit_first = limit % ListSize
			limit = ListSize
		}

		j := 0
		for i := *cpos; i < limit; i++ {
			aux_list[j] = (*list)[i]
			j++
		}

		for i := 0; i < limit_first; i++ {
			aux_list[j] = (*list)[i]
			j++
		}

		j -= 1
		for i := *cpos; i < limit; i++ {
			(*list)[i] = aux_list[j]
			j--
		}

		for i := 0; i < limit_first; i++ {
			(*list)[i] = aux_list[j]
			j--
		}

		*cpos = (*cpos + int(length) + *skip) % ListSize
		(*skip)++
	}
}

func part_one(input string) int {
	list := make([]int, ListSize)
	lengths := make([]byte, 0)
	cpos := 0
	skip := 0

	for i := 0; i < ListSize; i++ {
		list[i] = i
	}

	for _, ls := range strings.Split(input, ",") {
		length8, err := strconv.ParseUint(strings.TrimSpace(ls), 0, 8)
		if err != nil {
			fmt.Println("error parsing int", err)
		}
		lengths = append(lengths, byte(length8))
	}

	one_round(lengths, &cpos, &skip, &list)

	return list[0] * list[1]
}

func part_two(input string) string {
	list := make([]int, ListSize)
	lengths := make([]byte, 0)
	cpos := 0
	skip := 0
	result := ""

	for i := 0; i < ListSize; i++ {
		list[i] = i
	}

	for _, b := range strings.TrimSpace(input) {
		lengths = append(lengths, byte(b))
	}

	lengths = append(lengths, []byte{17, 31, 73, 47, 23}...)

	for i := 0; i < 64; i++ {
		one_round(lengths, &cpos, &skip, &list)
	}

	for i := 0; i < 16; i++ {
		n := list[i<<4]
		for j := 1; j < 16; j++ {
			n ^= list[(i<<4)+j]
		}

		result = fmt.Sprintf("%s%0.2x", result, n)
	}

	return result
}

func day14() int {
	used := 0

	for i := 0; i < 128; i++ {
		result := part_two(fmt.Sprintf("hfdlxzhv-%d", i))
		for _, b := range result {
			switch b {
			case '1', '2', '4', '8':
				used += 1
			case '3', '5', '6', '9', 'a', 'c':
				used += 2
			case '7', 'b', 'd', 'e':
				used += 3
			case 'f':
				used += 4
			}
		}
	}

	return used
}

func day142() int {
	var grid [128][128]bool
	regions := 0

	for i := 0; i < 128; i++ {
		result := part_two(fmt.Sprintf("hfdlxzhv-%d", i))
		for j, b := range result {
			if b == '8' || b == '9' || b == 'a' || b == 'b' || b == 'c' || b == 'd' || b == 'e' || b == 'f' {
				grid[i][(j<<2)+0] = true
			}
			if b == '4' || b == '5' || b == '6' || b == '7' || b == 'c' || b == 'd' || b == 'e' || b == 'f' {
				grid[i][(j<<2)+1] = true
			}
			if b == '2' || b == '3' || b == '6' || b == '7' || b == 'a' || b == 'b' || b == 'e' || b == 'f' {
				grid[i][(j<<2)+2] = true
			}
			if b == '1' || b == '3' || b == '5' || b == '7' || b == '9' || b == 'b' || b == 'd' || b == 'f' {
				grid[i][(j<<2)+3] = true
			}
		}
	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if !grid[i][j] {
				continue
			}

			regions++

			queue := make([]Square, 0)
			head := 0

			queue = append(queue, Square{i: i, j: j})
			for head < len(queue) {
				elem := queue[head]
				head++
				grid[elem.i][elem.j] = false

				if elem.i > 0 && grid[elem.i-1][elem.j] {
					queue = append(queue, Square{i: elem.i - 1, j: elem.j})
				}

				if elem.i < 127 && grid[elem.i+1][elem.j] {
					queue = append(queue, Square{i: elem.i + 1, j: elem.j})
				}

				if elem.j > 0 && grid[elem.i][elem.j-1] {
					queue = append(queue, Square{i: elem.i, j: elem.j - 1})
				}

				if elem.j < 127 && grid[elem.i][elem.j+1] {
					queue = append(queue, Square{i: elem.i, j: elem.j + 1})
				}
			}
		}
	}

	return regions
}

func main() {
	fcontent, err := ioutil.ReadFile("knot.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
	fmt.Println(day14())
	fmt.Println(day142())
}
