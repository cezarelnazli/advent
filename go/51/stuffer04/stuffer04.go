package main

import (
	"crypto/md5"
	"fmt"
)

func part_one(input string) int {
	i := 1
	for {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		if sum[0] == 0x00 && sum[1] == 0x00 && sum[2] == 0x00 {
			break
		}

		i++
	}

	return i
}

func main() {
	fmt.Println(part_one("yzbqklnj"))
}
