package main

import (
	"fmt"
)

func part_one(input int) int {
	buffer := make([]int, 2018)

	start_pos := 0
	buffer[start_pos] = 0

	for i := 1; i <= 2017; i++ {
		buffer[i] = (start_pos + input) % i
		if (start_pos+input)%i < i-1 {
			for j := i; j > (start_pos+input)%i; j-- {
				buffer[j] = buffer[j-1]
			}
		}

		buffer[(start_pos+input)%i+1] = i
		start_pos = (start_pos+input)%i + 1
	}

	return buffer[start_pos+1]
}

func part_two(input int) int {
	start_pos := 0
	result := 0

	for i := 1; i <= 50000000; i++ {
		if (start_pos+input)%i == 0 {
			result = i
		}

		start_pos = (start_pos+input)%i + 1
	}

	return result
}

func main() {
	fmt.Println(part_one(329))
	fmt.Println(part_two(329))
}
