package main

import (
	"fmt"
	"math"
)

const SIZE int = 1000000
const SIZE2 int = 100000000

var primes [SIZE]int
var NP int
var houses [SIZE2]int

func mk_sieve() {
	var sieve [SIZE]bool

	for i := 2; i < SIZE; i++ {
		if !sieve[i] {
			primes[NP] = i
			NP++
			for j := i * 2; j < SIZE; j += i {
				sieve[j] = true
			}
		}
	}

}

func part_one(input int) int {
	for i := 2; i < SIZE; i++ {
		sum := 1
		n := i
		for j := 0; primes[j]*primes[j] <= i; j++ {
			if i%primes[j] != 0 {
				continue
			}

			d := 1
			for n%primes[j] == 0 {
				d++
				n /= primes[j]
			}

			sum *= (int(math.Pow(float64(primes[j]), float64(d))) - 1) / (primes[j] - 1)
		}

		if n > 1 {
			sum *= n + 1
		}

		if sum > input {
			return i
		}
	}

	return 0
}

func part_two(input int) int {
	for i := 1; i < SIZE2/50; i++ {
		for j := 1; j <= 50; j++ {
			houses[i*j] += i * 11
		}
	}

	for i := 1; i < SIZE2; i++ {
		if houses[i] > input {
			return i
		}
	}

	return 0
}

func main() {
	mk_sieve()

	fmt.Println(part_one(36000000 / 10))
	fmt.Println(part_two(36000000))
}
