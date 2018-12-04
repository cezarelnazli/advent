package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part_one(input string, travel int) int {
	max := 0

	for _, l := range strings.Split(input, "\n") {
		var who string
		var speed, time, rest int
		fmt.Sscanf(l,
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&who, &speed, &time, &rest)

		cnt := travel / (time + rest) * (speed * time)
		if time < travel%(time+rest) {
			cnt += time * speed
		} else {
			cnt += (travel % (time + rest)) * speed
		}

		if cnt > max {
			max = cnt
		}
	}

	return max
}

type Reindeer struct {
	speed int
	time  int
	rest  int
}

func part_two(input string, travel int) int {
	reindeers := make([]Reindeer, 0)
	leads := make([]int, 0)

	for _, l := range strings.Split(input, "\n") {
		var who string
		var speed, time, rest int
		fmt.Sscanf(l,
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&who, &speed, &time, &rest)
		reindeers = append(reindeers,
			Reindeer{speed: speed, time: time, rest: rest})
		leads = append(leads, 0)
	}

	for t := 1; t <= travel; t++ {
		max := 0
		for _, r := range reindeers {
			cnt := t / (r.time + r.rest) * (r.speed * r.time)
			if r.time < t%(r.time+r.rest) {
				cnt += r.time * r.speed
			} else {
				cnt += (t % (r.time + r.rest)) * r.speed
			}

			if cnt > max {
				max = cnt
			}
		}

		for i, r := range reindeers {
			cnt := t / (r.time + r.rest) * (r.speed * r.time)
			if r.time < t%(r.time+r.rest) {
				cnt += r.time * r.speed
			} else {
				cnt += (t % (r.time + r.rest)) * r.speed
			}

			if cnt == max {
				leads[i]++
			}
		}
	}

	max := 0
	for _, l := range leads {
		if l > max {
			max = l
		}
	}

	return max
}

func main() {
	fcontent, err := ioutil.ReadFile("olympics14.in")
	if err != nil {
		fmt.Println("error reading file", err)
	} else {
		input := string(fcontent)

		fmt.Println(part_one(input, 2503))
		fmt.Println(part_two(input, 2503))
	}
}
