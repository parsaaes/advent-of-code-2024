package main

import (
	"fmt"
	"strconv"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type (
	cell struct {
		i, j int
	}

	stack []cell
)

func (s *stack) push(c cell) {
	*s = append(*s, c)
}

func (s *stack) pop() cell {
	if len(*s) == 0 {
		return cell{-1, -1}
	}

	result := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return result
}

func main() {
	grid := input.ReadPixels()

	cellVals := make(map[cell]int)

	for i := range grid {
		for j := range grid[i] {
			val, err := strconv.Atoi(string(grid[i][j]))
			if err != nil {
				continue
			}

			cellVals[cell{i, j}] = val
		}
	}

	directions := []cell{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	part1Scores := make(map[cell]map[cell]struct{})
	part2Scores := make(map[cell]int)

	for k, v := range cellVals {
		if v != 0 {
			continue
		}

		startCell := k

		s := &stack{}

		s.push(startCell)

		for len(*s) != 0 {
			currentCell := s.pop()
			currentVal := cellVals[currentCell]

			if currentVal == 9 {
				if _, ok := part1Scores[startCell]; !ok {
					part1Scores[startCell] = make(map[cell]struct{})
				}

				part1Scores[startCell][currentCell] = struct{}{}
				part2Scores[startCell]++
			}

			for i := range directions {
				potentialI := currentCell.i + directions[i].i
				potentialJ := currentCell.j + directions[i].j

				if potentialI < 0 || potentialJ < 0 {
					continue
				}

				if potentialI >= len(grid) || potentialJ >= len(grid[0]) {
					continue
				}

				potentialCell := cell{potentialI, potentialJ}

				potentialVal := cellVals[potentialCell]

				if potentialVal == currentVal+1 {
					s.push(potentialCell)
				}
			}
		}
	}

	part1 := 0
	part2 := 0

	for k, v := range part2Scores {
		part1 += len(part1Scores[k])
		part2 += v
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
