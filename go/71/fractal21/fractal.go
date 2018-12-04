package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func rotate(from []string) []string {
	aux := make([][]byte, len(from))

	for i := 0; i < len(from); i++ {
		aux[i] = make([]byte, len(from))
	}

	for i := 0; i < len(aux); i++ {
		for j := 0; j < len(aux); j++ {
			aux[i][j] = from[len(aux)-1-j][i]
		}
	}

	auxs := make([]string, len(aux))
	for i := 0; i < len(aux); i++ {
		auxs[i] = string(aux[i])
	}

	return auxs
}

func hflip(from []string) []string {
	aux := make([][]byte, len(from))

	for i := 0; i < len(from); i++ {
		aux[i] = make([]byte, len(from))
	}

	for i := 0; i < len(aux); i++ {
		for j := 0; j < len(aux); j++ {
			aux[i][j] = from[i][len(aux)-1-j]
		}
	}

	auxs := make([]string, len(aux))
	for i := 0; i < len(aux); i++ {
		auxs[i] = string(aux[i])
	}

	return auxs
}

func vflip(from []string) []string {
	aux := make([][]byte, len(from))

	for i := 0; i < len(from); i++ {
		aux[i] = make([]byte, len(from))
	}

	for i := 0; i < len(aux); i++ {
		for j := 0; j < len(aux); j++ {
			aux[i][j] = from[len(aux)-1-i][j]
		}
	}

	auxs := make([]string, len(aux))
	for i := 0; i < len(aux); i++ {
		auxs[i] = string(aux[i])
	}

	return auxs
}

func part_one(input string, limit int) int {
	enhancements := make(map[string]string)

	for _, l := range strings.Split(input, "\n") {
		patterns := strings.Split(l, " => ")
		from, to := patterns[0], patterns[1]

		enhancements[from] = to
		enhancements[strings.Join(hflip(strings.Split(from, "/")), "/")] = to
		enhancements[strings.Join(vflip(strings.Split(from, "/")), "/")] = to

		from = strings.Join(rotate(strings.Split(from, "/")), "/")

		enhancements[from] = to
		enhancements[strings.Join(hflip(strings.Split(from, "/")), "/")] = to
		enhancements[strings.Join(vflip(strings.Split(from, "/")), "/")] = to

		from = strings.Join(rotate(strings.Split(from, "/")), "/")

		enhancements[from] = to
		enhancements[strings.Join(hflip(strings.Split(from, "/")), "/")] = to
		enhancements[strings.Join(vflip(strings.Split(from, "/")), "/")] = to

		from = strings.Join(rotate(strings.Split(from, "/")), "/")

		enhancements[from] = to
		enhancements[strings.Join(hflip(strings.Split(from, "/")), "/")] = to
		enhancements[strings.Join(vflip(strings.Split(from, "/")), "/")] = to
	}

	start := []string{".#.", "..#", "###"}

	for i := 0; i < limit; i++ {
		for d := 2; d <= 3; d++ {
			if len(start)%d == 0 {
				result := make([]string, (d+1)*len(start)/d)
				aux := make([]string, d)

				for j := 0; j < len(start); j += d {
					for k := 0; k < len(start); k += d {
						for l := 0; l < d; l++ {
							aux[l] = ""
							for m := 0; m < d; m++ {
								aux[l] += string(start[j+l][k+m])
							}
						}

						resaux := strings.Split(enhancements[strings.Join(aux, "/")], "/")
						for l := 0; l <= d; l++ {
							result[(j/d*(d+1))+l] += resaux[l]
						}
					}
				}

				start = result
				break
			}
		}
	}

	non := 0
	for i := 0; i < len(start); i++ {
		for j := 0; j < len(start); j++ {
			if start[i][j] == '#' {
				non++
			}
		}
	}

	return non
}

func main() {
	fcontent, err := ioutil.ReadFile("fractal.in")
	if err != nil {
		fmt.Println("error reading file", err)
	}

	input := string(fcontent)

	fmt.Println(part_one(input, 5))
	fmt.Println(part_one(input, 18))
}
