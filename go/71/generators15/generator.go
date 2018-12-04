package main

import (
	"fmt"
)

const (
	NUM_ROUNDS_1 int = 40000000
	NUM_ROUNDS_2 int = 5000000
	FACTOR_A     int = 16807
	FACTOR_B     int = 48271
	MOD          int = 2147483647
)

func part_one(numA int, numB int) int {
	count := 0

	for i := 0; i < NUM_ROUNDS_1; i++ {
		numA = (numA * FACTOR_A) % MOD
		numB = (numB * FACTOR_B) % MOD

		if numA&0xFFFF == numB&0xFFFF {
			count++
		}
	}

	return count
}

func part_two(numA int, numB int) int {
	seqA := make([]int, 0)
	seqB := make([]int, 0)
	count := 0

	for i := 0; i < NUM_ROUNDS_2; i++ {
		numA = (numA * FACTOR_A) % MOD

		if numA&3 == 0 {
			seqA = append(seqA, numA)
		} else {
			i--
		}
	}

	for i := 0; i < NUM_ROUNDS_2; i++ {
		numB = (numB * FACTOR_B) % MOD

		if numB&7 == 0 {
			seqB = append(seqB, numB)
		} else {
			i--
		}
	}

	for i := 0; i < NUM_ROUNDS_2; i++ {
		if seqA[i]&0xFFFF == seqB[i]&0xFFFF {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(part_one(116, 299))
	fmt.Println(part_two(116, 299))
	fmt.Println(part_two(65, 8921))
}
