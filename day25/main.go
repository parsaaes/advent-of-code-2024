package main

import (
	"fmt"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	in := input.ReadBulk()

	rawGrids := strings.Split(in, "\n\n")

	grids := make([][][]byte, len(rawGrids))

	for i := range rawGrids {
		grids[i] = make([][]byte, 0)

		gridLines := strings.Split(rawGrids[i], "\n")

		for j := range gridLines {
			grids[i] = append(grids[i], []byte(gridLines[j]))
		}
	}

	locks := make([][]int, 0)
	keys := make([][]int, 0)

	for i := range grids {
		grid := grids[i]

		if isLock(grid) {
			lockPattern := make([]int, 5)

			for q := 0; q < 5; q++ {
				for p := 1; p <= 6; p++ {
					if grid[p][q] == '#' {
						lockPattern[q]++
					}
				}
			}

			locks = append(locks, lockPattern)
		} else {
			keyPattern := make([]int, 5)

			for q := 0; q < 5; q++ {
				for p := 5; p >= 0; p-- {
					if grid[p][q] == '#' {
						keyPattern[q]++
					}
				}
			}

			keys = append(keys, keyPattern)
		}
	}

	total := 0
	for _, lock := range locks {
		for _, key := range keys {
			fit := true

			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					fit = false
					break
				}
			}

			if fit {
				total++
			}
		}
	}

	fmt.Println(total)
}

func isLock(grid [][]byte) bool {
	for i := 0; i < 5; i++ {
		if grid[0][i] != '#' {
			return false
		}
	}

	return true
}
