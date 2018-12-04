package main

import (
	"fmt"
)

const (
	START = 20151125
	MUL   = 252533
	MOD   = 33554393
)

func part_one(lf, cf int) int {
	cl := 2
	new_num := START

	for {
		i := cl
		j := 1

		for i > 0 {
			new_num = ((new_num % MOD) * (MUL % MOD)) % MOD

			if i == lf && j == cf {
				return new_num
			}

			i--
			j++
		}

		cl++
	}
}

func main() {
	fmt.Println(part_one(3010, 3019))
}
