package main

import "fmt"

func part_one(input string) string {
	binput := []byte(input)
	for {
		place := 7
		for place >= 0 && binput[place] == 'z' {
			binput[place] = 'a'
			place--
		}

		binput[place] = binput[place] + 1

		straight, confusing := false, false
		pairs := 0

		for i := 0; !confusing && i < len(binput); i++ {
			confusing = binput[i] == 'i' || binput[i] == 'o' || binput[i] == 'l'
		}

		for i := 0; !straight && i < len(binput)-2; i++ {
			straight = binput[i]+1 == binput[i+1] && binput[i]+2 == binput[i+2]
		}

		for i := 0; i < len(binput)-1; i++ {
			if binput[i] == binput[i+1] && binput[i] != binput[i-1] {
				pairs++
			}
		}

		if !confusing && straight && pairs >= 2 {
			return string(binput)
		}
	}
}

func main() {
	fmt.Println(part_one("vzbxkghb"))
	fmt.Println(part_one("vzbxxyzz"))
}
