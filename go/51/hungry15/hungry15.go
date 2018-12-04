package main

import (
	"fmt"
)

type Ingredient struct {
	capc int
	dur  int
	fla  int
	tex  int
	cal  int
}

func part_one(ingredients *[4]Ingredient, account_cal bool) int {
	max := 0

	for i1 := 0; i1 <= 100; i1++ {
		for i2 := 0; i2 <= 100-i1; i2++ {
			for i3 := 0; i3 <= 100-i1-i2; i3++ {
				i4 := 100 - i1 - i2 - i3

				if account_cal {
					cal := i1*(*ingredients)[0].cal + i2*(*ingredients)[1].cal +
						i3*(*ingredients)[2].cal + i4*(*ingredients)[3].cal
					if cal != 500 {
						continue
					}
				}

				capc := i1*(*ingredients)[0].capc + i2*(*ingredients)[1].capc +
					i3*(*ingredients)[2].capc + i4*(*ingredients)[3].capc
				if capc <= 0 {
					continue
				}

				dur := i1*(*ingredients)[0].dur + i2*(*ingredients)[1].dur +
					i3*(*ingredients)[2].dur + i4*(*ingredients)[3].dur
				if dur <= 0 {
					continue
				}

				fla := i1*(*ingredients)[0].fla + i2*(*ingredients)[1].fla +
					i3*(*ingredients)[2].fla + i4*(*ingredients)[3].fla
				if fla <= 0 {
					continue
				}

				tex := i1*(*ingredients)[0].tex + i2*(*ingredients)[1].tex +
					i3*(*ingredients)[2].tex + i4*(*ingredients)[3].tex
				if tex <= 0 {
					continue
				}

				if capc*dur*fla*tex > max {
					max = capc * dur * fla * tex
				}
			}
		}
	}

	return max
}

func main() {
	var ingredients [4]Ingredient

	ingredients[0] = Ingredient{capc: 5, dur: -1, fla: 0, tex: 0, cal: 5}
	ingredients[1] = Ingredient{capc: -1, dur: 3, fla: 0, tex: 0, cal: 1}
	ingredients[2] = Ingredient{capc: 0, dur: -1, fla: 4, tex: 0, cal: 6}
	ingredients[3] = Ingredient{capc: -1, dur: 0, fla: 0, tex: 2, cal: 8}

	fmt.Println(part_one(&ingredients, false))
	fmt.Println(part_one(&ingredients, true))
}
