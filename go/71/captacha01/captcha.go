package main

import "fmt"
import "io/ioutil"

func part_one(input string) int {
	sum := 0

	if input[0] == input[len(input)-1] {
		sum += int(input[0] - '0')
	}

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			sum += int(input[i] - '0')
		}
	}

	return sum
}

func part_two(input string) int {
	sum := 0
	step := len(input) >> 1

	for i := 0; i < len(input)-step; i++ {
		if input[i] == input[i+step] {
			sum += int(input[i] - '0')
		}
	}

	return sum << 1
}

func main() {
	fcontent, err := ioutil.ReadFile("captcha.in")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
