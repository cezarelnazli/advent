package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Triple struct {
	x int
	y int
	z int
}

type Particle struct {
	p         Triple
	v         Triple
	a         Triple
	is_active bool
}

const TIMES int = 500

func part_one(input string) int {
	particles := make([]Particle, 0)

	for _, l := range strings.Split(input, "\n") {
		var px, py, pz, vx, vy, vz, ax, ay, az int

		fmt.Sscanf(l, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&px, &py, &pz, &vx, &vy, &vz, &ax, &ay, &az)

		particles = append(particles,
			Particle{
				p: Triple{x: px, y: py, z: pz},
				v: Triple{x: vx, y: vy, z: vz},
				a: Triple{x: ax, y: ay, z: az}})
	}

	for j := 0; j < TIMES; j++ {
		for i := 0; i < len(particles); i++ {
			particles[i].v.x += particles[i].a.x
			particles[i].v.y += particles[i].a.y
			particles[i].v.z += particles[i].a.z

			particles[i].p.x += particles[i].v.x
			particles[i].p.y += particles[i].v.y
			particles[i].p.z += particles[i].v.z
		}
	}

	imin := 0
	min := int(math.Abs(float64(particles[0].p.x))) + int(math.Abs(float64(particles[0].p.y))) + int(math.Abs(float64(particles[0].p.z)))
	for i := 1; i < len(particles); i++ {
		val := int(math.Abs(float64(particles[i].p.x))) + int(math.Abs(float64(particles[i].p.y))) + int(math.Abs(float64(particles[i].p.z)))
		if val < min {
			min = val
			imin = i
		}
	}

	return imin
}

func part_two(input string) int {
	particles := make([]Particle, 0)
	nleft := 0

	for _, l := range strings.Split(input, "\n") {
		var px, py, pz, vx, vy, vz, ax, ay, az int

		fmt.Sscanf(l, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&px, &py, &pz, &vx, &vy, &vz, &ax, &ay, &az)

		particles = append(particles,
			Particle{
				p:         Triple{x: px, y: py, z: pz},
				v:         Triple{x: vx, y: vy, z: vz},
				a:         Triple{x: ax, y: ay, z: az},
				is_active: true})

		nleft++
	}

	for j := 0; j < TIMES; j++ {
		collisions := make(map[Triple]int)
		to_delete := make(map[int]bool, 1000)

		for i := 0; i < len(particles); i++ {
			if particles[i].is_active {
				particles[i].v.x += particles[i].a.x
				particles[i].v.y += particles[i].a.y
				particles[i].v.z += particles[i].a.z

				particles[i].p.x += particles[i].v.x
				particles[i].p.y += particles[i].v.y
				particles[i].p.z += particles[i].v.z

				cd, collides := collisions[particles[i].p]
				if collides {
					to_delete[i] = true
					to_delete[cd] = true
				} else {
					collisions[particles[i].p] = i
				}
			}
		}

		for i, d := range to_delete {
			if d {
				particles[i].is_active = false
				nleft--
			}
		}
	}

	return nleft
}

func main() {
	fcontent, err := ioutil.ReadFile("particles.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
