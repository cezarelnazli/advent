package main

import "fmt"
import "os"

type Player struct {
	atk  int
	def  int
	hp   int
	mana int
}

type Spell struct {
	atk   int
	def   int
	cost  int
	turns int
}

type ActiveSpell struct {
	spell Spell
	timer int
}

var cmin, total_mana int
var spells [5]Spell

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bkt(player, boss Player, player_turn bool, activeSpells [3]ActiveSpell, turn int) {

	if turn > 17 {
		return
	}

	if player_turn {
		player.hp--
	}

	if activeSpells[0].timer > 0 {
		player.def = activeSpells[0].spell.def
		activeSpells[0].timer--
	} else {
		player.def = 0
	}

	if activeSpells[1].timer > 0 {
		boss.hp -= activeSpells[1].spell.atk
		activeSpells[1].timer--
	}

	if activeSpells[2].timer > 0 {
		player.mana += activeSpells[2].spell.def
		activeSpells[2].timer--
	}

	if player.hp > 0 && boss.hp > 0 {
		if player_turn {
			for i, s := range spells {
				if s.cost <= player.mana {
					if i > 1 && activeSpells[i-2].timer > 0 {
						continue
					}

					new_player := player
					new_player.mana -= s.cost

					new_boss := boss

					total_mana += s.cost
					if i == 0 {
						new_boss.hp -= s.atk
					} else if i == 1 {
						new_boss.hp -= s.atk
						new_player.hp += s.def
					} else {
						activeSpells[i-2] = ActiveSpell{spell: s, timer: s.turns}
					}

					bkt(new_player, new_boss, false, activeSpells, turn+1)
					total_mana -= s.cost
					if i > 1 {
						activeSpells[i-2] = ActiveSpell{spell: s, timer: 0}
					}
				}
			}
		} else {
			paux := player.hp
			player.hp -= boss.atk - player.def
			bkt(player, boss, true, activeSpells, turn+1)
			player.hp = paux
		}
	} else {
		if player.hp > 0 && boss.hp <= 0 {
			if total_mana < cmin {
				cmin = total_mana
				fmt.Fprintf(os.Stderr, "%d\n", cmin)
			}
		}
	}
}

func part_one() int {
	cmin = 10000000
	spells[0] = Spell{atk: 4, def: 0, cost: 53, turns: 0}
	spells[1] = Spell{atk: 2, def: 2, cost: 73, turns: 0}
	spells[2] = Spell{atk: 0, def: 7, cost: 113, turns: 6}
	spells[3] = Spell{atk: 3, def: 0, cost: 173, turns: 6}
	spells[4] = Spell{atk: 0, def: 101, cost: 229, turns: 5}

	player := Player{atk: 0, def: 0, hp: 50, mana: 500}
	boss := Player{atk: 9, def: 0, hp: 58, mana: 0}

	bkt(player, boss, true,
		[3]ActiveSpell{ActiveSpell{}, ActiveSpell{}, ActiveSpell{}}, 0)

	return cmin
}

func main() {
	fmt.Println(part_one())
}
