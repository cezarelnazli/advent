package main

import "fmt"

type Player struct {
	atk int
	def int
	hp  int
}

type Item struct {
	atk  int
	def  int
	cost int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func turns_to_beat(w, a, r1, r2, vs Item) int {
	return vs.cost / max(1, w.atk+r1.atk+r2.atk-vs.def)
}

func part_one() int {
	weapons := []Item{
		Item{atk: 4, def: 0, cost: 8},
		Item{atk: 5, def: 0, cost: 10},
		Item{atk: 6, def: 0, cost: 25},
		Item{atk: 7, def: 0, cost: 40},
		Item{atk: 8, def: 0, cost: 74}}

	armour := []Item{
		Item{atk: 0, def: 0, cost: 0},
		Item{atk: 0, def: 1, cost: 13},
		Item{atk: 0, def: 2, cost: 31},
		Item{atk: 0, def: 3, cost: 53},
		Item{atk: 0, def: 4, cost: 75},
		Item{atk: 0, def: 5, cost: 102}}

	rings := []Item{
		Item{atk: 0, def: 0, cost: 0},
		Item{atk: 0, def: 0, cost: 0},
		Item{atk: 1, def: 0, cost: 25},
		Item{atk: 2, def: 0, cost: 50},
		Item{atk: 3, def: 0, cost: 100},
		Item{atk: 0, def: 1, cost: 20},
		Item{atk: 0, def: 2, cost: 40},
		Item{atk: 0, def: 3, cost: 80}}

	boss := Item{atk: 8, def: 2, cost: 100}

	cmin := 1000
	for _, w := range weapons {
		for _, a := range armour {
			for i, r1 := range rings {
				for _, r2 := range rings[i+1:] {
					if turns_to_beat(w, a, r1, r2, boss) <= 100/max(1, boss.atk-(a.def+r1.def+r2.def)) {
						if w.cost+a.cost+r1.cost+r2.cost < cmin {
							cmin = w.cost + a.cost + r1.cost + r2.cost
						}
					}
				}
			}
		}
	}

	return cmin
}

func part_two() int {
	weapons := []Item{
		Item{atk: 4, def: 0, cost: 8},
		Item{atk: 5, def: 0, cost: 10},
		Item{atk: 6, def: 0, cost: 25},
		Item{atk: 7, def: 0, cost: 40},
		Item{atk: 8, def: 0, cost: 74}}

	armour := []Item{
		Item{atk: 0, def: 0, cost: 0},
		Item{atk: 0, def: 1, cost: 13},
		Item{atk: 0, def: 2, cost: 31},
		Item{atk: 0, def: 3, cost: 53},
		Item{atk: 0, def: 4, cost: 75},
		Item{atk: 0, def: 5, cost: 102}}

	rings := []Item{
		Item{atk: 0, def: 0, cost: 0},
		Item{atk: 0, def: 0, cost: 0},
		Item{atk: 1, def: 0, cost: 25},
		Item{atk: 2, def: 0, cost: 50},
		Item{atk: 3, def: 0, cost: 100},
		Item{atk: 0, def: 1, cost: 20},
		Item{atk: 0, def: 2, cost: 40},
		Item{atk: 0, def: 3, cost: 80}}

	boss := Item{atk: 8, def: 2, cost: 100}

	cmax := 0
	for _, w := range weapons {
		for _, a := range armour {
			for i, r1 := range rings {
				for _, r2 := range rings[i+1:] {
					if turns_to_beat(w, a, r1, r2, boss) > 100/max(1, boss.atk-(a.def+r1.def+r2.def)) {
						if w.cost+a.cost+r1.cost+r2.cost > cmax {
							cmax = w.cost + a.cost + r1.cost + r2.cost
						}
					}
				}
			}
		}
	}

	return cmax
}

func main() {
	fmt.Println(part_one())
	fmt.Println(part_two())
}
