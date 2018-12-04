package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	weight   int
	sum      int
	children []*Node
}

func part_one(input string) string {
	candidates := make([]string, 0)
	sustained := make([]string, 0)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")

		if len(parts) == 2 {
			var program string
			var weight int

			fmt.Sscanf(parts[0], "%s (%d)", &program, &weight)

			candidates = append(candidates, program)

			for _, s := range strings.Split(parts[1], ", ") {
				sustained = append(sustained, s)
			}

			for i := 0; i < len(candidates); i++ {
				for _, s := range sustained {
					if candidates[i] == s {
						candidates[len(candidates)-1], candidates[i] = candidates[i], candidates[len(candidates)-1]
						candidates = candidates[:len(candidates)-1]
						break
					}
				}
			}
		}
	}

	return candidates[0]
}

func make_sum(root *Node) {
	root.sum = root.weight

	for _, child := range root.children {
		make_sum(child)
		root.sum += child.sum
	}
}

func check_sum(root *Node) (int, bool) {
	queue := make([]*Node, 0)
	head := 0

	queue = append(queue, root)

	for head < len(queue) {
		node := queue[head]
		head++

		if len(node.children) > 0 {
			var benchmark int
			if node.children[0].sum == node.children[1].sum {
				benchmark = node.children[0].sum
			} else {
				if node.children[0].sum == node.children[2].sum {
					benchmark = node.children[0].sum
				} else {
					benchmark = node.children[1].sum
				}
			}

			for _, child := range node.children {
				queue = append(queue, child)

				if child.sum != benchmark {
					w, b := check_sum(child)
					if b {
						return child.weight - (child.sum - benchmark), false
					} else {
						return w, false
					}
				}
			}
		}
	}

	return root.weight, true
}

func part_two(input string) int {
	nodes := make(map[string]*Node)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")

		var program string
		var weight int

		fmt.Sscanf(parts[0], "%s (%d)", &program, &weight)

		node, exists := nodes[program]
		if exists {
			node.weight = weight
		} else {
			nodes[program] = &Node{weight: weight, children: make([]*Node, 0)}
			node = nodes[program]
		}

		if len(parts) == 2 {
			for _, s := range strings.Split(parts[1], ", ") {
				_, exists := nodes[s]
				if !exists {
					nodes[s] = &Node{children: make([]*Node, 0)}
				}

				nodes[program].children = append(nodes[program].children, nodes[s])
			}
		}
	}

	root := nodes[part_one(input)]
	make_sum(root)
	r, _ := check_sum(root)
	return r
}

func main() {
	fcontent, err := ioutil.ReadFile("towers.in")
	if err != nil {
		fmt.Println("error reading file")
	}

	input := string(fcontent)

	fmt.Println(part_one(input))
	fmt.Println(part_two(input))
}
