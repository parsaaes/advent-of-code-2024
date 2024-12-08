package main

import (
	"fmt"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	cells := input.ReadPixels()

	part1(cells)
	part2(cells)
}

func part1(cells [][]byte) {
	findAntinodes(cells, false)
}

func part2(cells [][]byte) {
	findAntinodes(cells, true)
}

func findAntinodes(cells [][]byte, anyGrid bool) {
	antinodeMap := make([][]byte, len(cells))
	for i := range cells {
		antinodeMap[i] = make([]byte, len(cells[i]))
		for j := range cells[i] {
			antinodeMap[i][j] = cells[i][j]
		}
	}

	for i := range cells {
		for j := range cells[i] {
			if cells[i][j] != '.' {
				goDown(cells, i, j, antinodeMap, anyGrid)
				goRight(cells, i, j, antinodeMap, anyGrid)
				goRTL(cells, i, j, antinodeMap, anyGrid)
				goLTR(cells, i, j, antinodeMap, anyGrid)
			}
		}
	}

	total := 0
	for i := range antinodeMap {
		for j := range antinodeMap[i] {
			if antinodeMap[i][j] == '#' {
				total++
			}
		}
	}

	fmt.Println(total)
}

func goDown(cells [][]byte, p, q int, antinodeMap [][]byte, anyGrid bool) {
	anthena := cells[p][q]

	for i := p + 1; i < len(cells); i++ {
		if cells[i][q] == anthena {
			diffY := i - p

			if anyGrid {
				for y := i; y < len(cells); y += diffY {
					antinodeMap[y][q] = '#'
				}

				for y := i; y >= 0; y -= diffY {
					antinodeMap[y][q] = '#'
				}
			} else {
				firstAntinodeY := i + diffY

				if firstAntinodeY < len(cells) {
					antinodeMap[firstAntinodeY][q] = '#'
				}

				secondAntinodeY := p - diffY

				if secondAntinodeY > 0 {
					antinodeMap[secondAntinodeY][q] = '#'
				}
			}
		}
	}
}

func goRight(cells [][]byte, p, q int, antinodeMap [][]byte, anyGrid bool) {
	anthena := cells[p][q]

	for j := q + 1; j < len(cells[p]); j++ {
		if cells[p][j] == anthena {
			diffX := j - q

			if anyGrid {
				for x := j; x < len(cells[p]); x += diffX {
					antinodeMap[p][x] = '#'
				}

				for x := j; x >= 0; x -= diffX {
					antinodeMap[p][x] = '#'
				}
			} else {
				firstAntinodeX := j + diffX

				if firstAntinodeX < len(cells[p]) {
					antinodeMap[p][firstAntinodeX] = '#'
				}

				secondAntinodeX := q - diffX

				if secondAntinodeX > 0 {
					antinodeMap[p][secondAntinodeX] = '#'
				}
			}
		}
	}
}

func goRTL(cells [][]byte, p, q int, antinodeMap [][]byte, anyGrid bool) {
	anthena := cells[p][q]

	for i := p + 1; i < len(cells); i++ {
		for j := q + 1; j < len(cells[i]); j++ {
			if cells[i][j] == anthena {
				diffY := i - p
				diffX := j - q

				if anyGrid {
					y := i
					x := j
					for y < len(cells) && x < len(cells[y]) {
						antinodeMap[y][x] = '#'

						y += diffY
						x += diffX
					}

					y = i
					x = j
					for y >= 0 && x >= 0 {
						antinodeMap[y][x] = '#'

						y -= diffY
						x -= diffX
					}
				} else {
					firstAntinodeY := i + diffY
					firstAntinodeX := j + diffX

					if firstAntinodeY < len(cells) && firstAntinodeX < len(cells[firstAntinodeY]) {
						antinodeMap[firstAntinodeY][firstAntinodeX] = '#'
					}

					secondAntinodeY := p - diffY
					secondAntinodeX := q - diffX

					if secondAntinodeY >= 0 && secondAntinodeX >= 0 {
						antinodeMap[secondAntinodeY][secondAntinodeX] = '#'
					}
				}
			}
		}
	}
}

func goLTR(cells [][]byte, p, q int, antinodeMap [][]byte, anyGrid bool) {
	anthena := cells[p][q]

	for i := p + 1; i < len(cells); i++ {
		for j := q - 1; j >= 0; j-- {
			if cells[i][j] == anthena {
				diffY := i - p
				diffX := j - q

				if anyGrid {
					y := i
					x := j
					for y < len(cells) && x >= 0 {
						antinodeMap[y][x] = '#'

						y += diffY
						x += diffX
					}

					y = i
					x = j
					for y >= 0 && x < len(cells[y]) {
						antinodeMap[y][x] = '#'

						y -= diffY
						x -= diffX
					}
				} else {
					firstAntinodeY := i + diffY
					firstAntinodeX := j + diffX

					if firstAntinodeY < len(cells) && firstAntinodeX >= 0 {
						antinodeMap[firstAntinodeY][firstAntinodeX] = '#'
					}

					secondAntinodeY := p - diffY
					secondAntinodeX := q - diffX

					if secondAntinodeY >= 0 && secondAntinodeX < len(cells[secondAntinodeY]) {
						antinodeMap[secondAntinodeY][secondAntinodeX] = '#'
					}
				}
			}
		}
	}
}
