package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string) int {
	sum := 0

	for _, passphrase := range strings.Split(input, "\n") {
		seen := make(map[string]bool)
		sum += 1

		for _, word := range strings.Fields(passphrase) {
			_, found := seen[word]
			if found {
				sum -= 1
				break
			}

			seen[word] = true
		}
	}

	return sum
}

func part_two(input string) int {
	sum := 0

	for _, passphrase := range strings.Split(input, "\n") {
		seen := make(map[string]bool)
		letters := make([][26]int, 0)
		valid := true

		for _, word := range strings.Fields(passphrase) {
			_, found := seen[word]
			if found {
				valid = false
				break
			}

			var frequency [26]int

			for i := 0; i < len(word); i++ {
				frequency[word[i]-'a']++
			}

			for _, lfreq := range letters {
				if lfreq == frequency {
					valid = false
					break
				}
			}

			seen[word] = true
			letters = append(letters, frequency)

			if !valid {
				break
			}
		}

		if valid {
			sum++
		}
	}

	return sum
}

func main() {
	fcontent, err := ioutil.ReadFile("passphrase.in")
	if err != nil {
		fmt.Println("Cannot read file")
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
