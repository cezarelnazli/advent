package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type Queue struct {
	elems []int
	head  int
	tail  int
}

const QUEUE_SIZE int = 10000

func getInt(elem string, regs *map[string]int) int {
	result := 0

	if unicode.IsLetter(rune(elem[0])) {
		result = (*regs)[elem]
	} else {
		int64, err := strconv.ParseInt(elem, 0, 0)
		if err != nil {
			fmt.Println("error converting int", err)
		}

		result = int(int64)
	}

	return result
}

func part_one(input string) int {
	instrs := make([]string, 0)
	regs := make(map[string]int)
	last_played := 0

	for _, l := range strings.Split(input, "\n") {
		instrs = append(instrs, l)
	}

	for _, l := range "abcdefghijklmnopqrstuvwxyz" {
		regs[string(l)] = 0
	}

	for i := 0; i < len(instrs); i++ {
		elems := strings.Fields(instrs[i])

		switch elems[0] {
		case "snd":
			last_played = getInt(elems[1], &regs)
			break
		case "set":
			val := getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "add":
			val := regs[elems[1]]
			val += getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "mul":
			val := regs[elems[1]]
			val *= getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "mod":
			val := regs[elems[1]]
			val %= getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "rcv":
			x := getInt(elems[1], &regs)
			if x != 0 {
				return last_played
			}
			break
		case "jgz":
			x := getInt(elems[1], &regs)
			if x != 0 {
				i += getInt(elems[2], &regs)
				if i < 0 || i >= len(instrs) {
					return 0
				}

				i--
			}
			break
		}
	}

	return last_played
}

func one_worker(input string, id int, acc *int, is_blocking *[]bool, queuer, queuew *Queue, done chan bool) {
	instrs := make([]string, 0)
	regs := make(map[string]int)

	for _, l := range strings.Split(input, "\n") {
		instrs = append(instrs, l)
	}

	for _, l := range "abcdefghijklmnopqrstuvwxyz" {
		regs[string(l)] = 0
	}
	regs["p"] = id

	for i := 0; i < len(instrs); i++ {
		elems := strings.Fields(instrs[i])

		switch elems[0] {
		case "snd":
			queuew.elems[queuew.tail] = getInt(elems[1], &regs)
			queuew.tail = (queuew.tail + 1) % QUEUE_SIZE
			*acc++
			break
		case "set":
			val := getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "add":
			val := regs[elems[1]]
			val += getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "mul":
			val := regs[elems[1]]
			val *= getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "mod":
			val := regs[elems[1]]
			val %= getInt(elems[2], &regs)
			regs[elems[1]] = val
			break
		case "rcv":
			for queuer.tail-queuer.head == 0 {
				(*is_blocking)[id] = true
				if (*is_blocking)[1-id] == true {
					done <- true
					return
				}
			}

			regs[elems[1]] = queuer.elems[queuer.head]
			queuer.head = (queuer.head + 1) % QUEUE_SIZE

			(*is_blocking)[id] = false
			break
		case "jgz":
			x := getInt(elems[1], &regs)
			if x > 0 {
				i += getInt(elems[2], &regs)
				if i < 0 || i >= len(instrs) {
					done <- true
					return
				}

				i--
			}
			break
		}
	}

	done <- true
}

func part_two(input string) int {
	is_blocking := make([]bool, 2)
	var acca, accb int
	queuea := Queue{elems: make([]int, QUEUE_SIZE)}
	queueb := Queue{elems: make([]int, QUEUE_SIZE)}
	donea := make(chan bool, 1)
	doneb := make(chan bool, 1)

	go one_worker(input, 0, &acca, &is_blocking, &queuea, &queueb, donea)
	go one_worker(input, 1, &accb, &is_blocking, &queueb, &queuea, doneb)

	<-donea
	<-doneb

	return accb
}

func main() {
	fcontent, err := ioutil.ReadFile("duet.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
