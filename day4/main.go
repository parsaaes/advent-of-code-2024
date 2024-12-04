package main

import (
	"fmt"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	lines := input.Read()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != 'X' && lines[i][j] != 'S' {
				continue
			}

			backward := lines[i][j] == 'S'

			if findHorizontal(lines, i, j, backward) {
				total++
			}

			if findVertical(lines, i, j, backward) {
				total++
			}

			if findLTRDiagonal(lines, i, j, backward) {
				total++
			}

			if findRTLDiagonal(lines, i, j, backward) {
				total++
			}
		}
	}

	fmt.Println(total)
}

func part2(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != 'A' {
				continue
			}

			if i-1 < 0 || i+1 >= len(lines) {
				continue
			}

			if j-1 < 0 || j+1 >= len(lines[i]) {
				continue
			}

			ltrDiagonal := string(lines[i-1][j-1]) + string(lines[i+1][j+1])
			rtlDiagonal := string(lines[i-1][j+1]) + string(lines[i+1][j-1])

			if ltrDiagonal != "MS" && ltrDiagonal != "SM" {
				continue
			}

			if rtlDiagonal != "MS" && rtlDiagonal != "SM" {
				continue
			}

			total++
		}
	}

	fmt.Println(total)
}

func findHorizontal(lines []string, p, q int, backward bool) bool {
	if q+3 >= len(lines[p]) {
		return false
	}

	if backward {
		return lines[p][q:q+4] == "SAMX"
	}

	return lines[p][q:q+4] == "XMAS"
}

func findVertical(lines []string, p, q int, backward bool) bool {
	if p+3 >= len(lines) {
		return false
	}

	str := string(lines[p][q]) + string(lines[p+1][q]) + string(lines[p+2][q]) + string(lines[p+3][q])

	if backward {
		return str == "SAMX"
	}

	return str == "XMAS"
}

func findLTRDiagonal(lines []string, p, q int, backward bool) bool {
	if p+3 >= len(lines) {
		return false
	}

	if q+3 >= len(lines[p+3]) {
		return false
	}

	str := string(lines[p][q]) + string(lines[p+1][q+1]) + string(lines[p+2][q+2]) + string(lines[p+3][q+3])

	if backward {
		return str == "SAMX"
	}

	return str == "XMAS"
}

func findRTLDiagonal(lines []string, p, q int, backward bool) bool {
	if p+3 >= len(lines) {
		return false
	}

	if q-3 < 0 {
		return false
	}

	str := string(lines[p][q]) + string(lines[p+1][q-1]) + string(lines[p+2][q-2]) + string(lines[p+3][q-3])

	if backward {
		return str == "SAMX"
	}

	return str == "XMAS"
}
