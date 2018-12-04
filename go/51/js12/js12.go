package main

import "fmt"
import "io/ioutil"
import "strconv"
import "unicode"

func part_one(input string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		aux := []byte("")
		j := i
		for unicode.IsDigit(rune(input[j])) || input[j] == '-' {
			aux = append(aux, input[j])
			j++
		}

		if i != j {
			num, err := strconv.ParseInt(string(aux), 0, 0)
			if err != nil {
				fmt.Println("cannot convert to int", err)
			} else {
				sum += int(num)
			}

			i = j
		}
	}

	return sum
}

func part_two(input string) int {
	binput := []byte(input)
	to_remove := 0

	for i := 0; i < len(binput)-4; i++ {

		fmt.Println(string(binput[i : i+6]))
		if string(binput[i:i+6]) == ":\"red\"" {
			to_remove = 1
			for j := i - 1; j >= 0 && to_remove > 0; j-- {
				if binput[j] == '}' {
					to_remove++
				} else if binput[j] == '{' {
					to_remove--
				}
				binput[j] = '#'
			}

			to_remove = 1
			j := i + 1
			for j < len(binput) && to_remove > 0 {
				if binput[j] == '{' {
					to_remove++
				} else if binput[j] == '}' {
					to_remove--
				}
				binput[j] = '#'
				j++
			}

			i = j
		}
	}

	fmt.Println(string(binput))

	return part_one(string(binput))
}

func main() {
	fcontent, err := ioutil.ReadFile("js12.in")
	if err != nil {
		fmt.Println("error reading file", err)
	} else {
		input := string(fcontent)

		fmt.Println(part_one(input))
		fmt.Println(part_two(input))
	}
}
