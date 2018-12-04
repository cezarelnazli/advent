package main

import "fmt"

func part_one(input string) int {
	for i := 0; i < 50; i++ {
		var aux []byte
		for j := 0; j < len(input); j++ {
			k := j + 1
			for k < len(input) && input[j] == input[k] {
				k++
			}

			aux = append(aux, byte('0'+(k-j)))
			aux = append(aux, input[j])
			j = k - 1
		}
		input = string(aux)
	}

	return len(input)
}

func main() {
	fmt.Println(part_one("1113222113"))
}
