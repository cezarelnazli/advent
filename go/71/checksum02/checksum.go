package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part_one(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		all_items := strings.Fields(line)

		lmin, err := strconv.ParseInt(all_items[0], 0, 0)
		if err != nil {
			fmt.Println("cannot convert lmin")
		}

		lmax, err := strconv.ParseInt(all_items[0], 0, 0)
		if err != nil {
			fmt.Println("cannot convert lmax")
		}

		for _, num_s := range all_items[1:] {
			num, err := strconv.ParseInt(num_s, 0, 0)
			if err != nil {
				fmt.Println("cannot convert num", num_s)
			}

			if num < lmin {
				lmin = num
			}

			if num > lmax {
				lmax = num
			}
		}

		sum += int(lmax - lmin)
	}

	return sum
}

func gcd(n1, n2 int64) int64 {
	if n1%n2 == 0 {
		return n2
	}

	return gcd(n2, n1%n2)
}

func part_two(input string) int {
	var sum int64 = 0

	for _, line := range strings.Split(input, "\n") {
		all_items := strings.Fields(line)

		for i, n1_s := range all_items {
			n1, err := strconv.ParseInt(n1_s, 0, 0)
			if err != nil {
				fmt.Println("cannot convert num", n1_s)
			}

			for _, n2_s := range all_items[i+1:] {
				n2, err := strconv.ParseInt(n2_s, 0, 0)
				if err != nil {
					fmt.Println("cannot convert num", n1_s)
				}

				gcd := gcd(n1, n2)

				if gcd == n1 || gcd == n2 {
					if n1 > n2 {
						sum += n1 / n2
					} else {
						sum += n2 / n1
					}
				}
			}
		}
	}

	return int(sum)
}

func main() {
	fcontent, err := ioutil.ReadFile("checksum.in")
	if err != nil {
		fmt.Println("Error reading file")
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
