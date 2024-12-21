package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type (
	cell struct {
		i, j int
	}

	memoKey struct {
		current, target byte
		iteration       int
	}
)

var (
	numericKeypad = map[byte]cell{
		'7': cell{0, 0},
		'8': cell{0, 1},
		'9': cell{0, 2},
		'4': cell{1, 0},
		'5': cell{1, 1},
		'6': cell{1, 2},
		'1': cell{2, 0},
		'2': cell{2, 1},
		'3': cell{2, 2},
		'0': cell{3, 1},
		'A': cell{3, 2},
	}

	directionalKeyPad = map[byte]cell{
		'^': cell{0, 1},
		'A': cell{0, 2},
		'<': cell{1, 0},
		'v': cell{1, 1},
		'>': cell{1, 2},
	}

	numericPanic = cell{3, 0}
)

func main() {
	lines := input.Read()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0

	for i := range lines {
		psk := findAllPossibleNumericKeyboards(numericKeypad['A'], lines[i])

		shortestLen := math.MaxInt64

		for p := range psk {
			finalLen := 0

			current := byte('A')

			for q := range psk[p] {
				memo := make(map[memoKey]int)
				finalLen += shortest(current, psk[p][q], 2, memo)
				current = psk[p][q]
			}

			if finalLen < shortestLen {
				shortestLen = finalLen
			}
		}

		numericStr := strings.TrimSuffix(lines[i], "A")

		num, _ := strconv.Atoi(numericStr)

		total += num * shortestLen
	}

	fmt.Println(total)
}

func part2(lines []string) {
	total := 0

	for i := range lines {
		psk := findAllPossibleNumericKeyboards(numericKeypad['A'], lines[i])

		shortestLen := math.MaxInt64

		for p := range psk {
			finalLen := 0

			current := byte('A')

			for q := range psk[p] {
				memo := make(map[memoKey]int)
				finalLen += shortest(current, psk[p][q], 25, memo)
				current = psk[p][q]
			}

			if finalLen < shortestLen {
				shortestLen = finalLen
			}
		}

		numericStr := strings.TrimSuffix(lines[i], "A")

		num, _ := strconv.Atoi(numericStr)

		total += num * shortestLen
	}

	fmt.Println(total)
}

// a not so good brute force :D
func findAllPossibleNumericKeyboards(current cell, target string) []string {
	result := []string{""}

	for i := range target {
		tempResults := make([]string, 0)

		targetCell, ok := numericKeypad[target[i]]
		if !ok {
			panic("invalid target: " + target)
		}

		diff := cell{
			i: targetCell.i - current.i,
			j: targetCell.j - current.j,
		}

		xMoves := ""
		if diff.j < 0 {
			xMoves = strings.Repeat("<", -1*diff.j)
		} else {
			xMoves = strings.Repeat(">", diff.j)
		}

		yMoves := ""
		if diff.i < 0 {
			yMoves = strings.Repeat("^", -1*diff.i)
		} else {
			yMoves = strings.Repeat("v", diff.i)
		}

		ones := xMoves
		zeros := yMoves

		if len(yMoves) > len(ones) {
			ones = yMoves
			zeros = xMoves
		}

		possibilities := calculatePossibilities(len(ones), len(zeros)+len(ones))

		for x := range possibilities {
			temp := make([]byte, len(zeros)+len(ones))

			for y := range possibilities[x] {
				temp[possibilities[x][y]] = ones[0]
			}

			for y := range temp {
				if temp[y] == 0 {
					temp[y] = zeros[0]
				}
			}

			temp = append(temp, 'A')

			panicFree := true

			tmpCurrent := current

			for p := range temp {
				var move cell

				switch temp[p] {
				case '>':
					move = cell{
						i: 0,
						j: 1,
					}
				case '<':
					move = cell{
						i: 0,
						j: -1,
					}
				case '^':
					move = cell{
						i: -1,
						j: 0,
					}
				case 'v':
					move = cell{
						i: 1,
						j: 0,
					}
				case 'A':
					break
				}

				next := cell{
					i: tmpCurrent.i + move.i,
					j: tmpCurrent.j + move.j,
				}

				if next == numericPanic {
					panicFree = false
					break
				}

				tmpCurrent = next
			}

			if panicFree {
				for y := range result {
					tempResults = append(tempResults, result[y]+string(temp))
				}
			}
		}

		result = make([]string, len(tempResults))
		copy(result, tempResults)

		current = targetCell
	}

	return result
}

func calculatePossibilities(ones int, total int) [][]int {
	result := make([][]int, 0)

	totalMovesWillBeAdded := total

	maximumPossibility := int(math.Pow(2, float64(totalMovesWillBeAdded))) - 1

	for p := maximumPossibility; p >= 0; p-- {
		onePositions := make([]int, 0)

		for onePosition := totalMovesWillBeAdded - 1; onePosition >= 0; onePosition-- {
			test := p >> onePosition

			if test%2 == 1 {
				onePositions = append(onePositions, onePosition)
			}
		}

		if len(onePositions) == ones {
			result = append(result, onePositions)
		}
	}

	return result
}

// this could be simpler :D
func shortest(current, target byte, iteration int, memo map[memoKey]int) (res int) {
	defer func() {
		memo[memoKey{
			current:   current,
			target:    target,
			iteration: iteration,
		}] = res
	}()

	end := directionalKeyPad[target]

	diff := cell{
		i: end.i - directionalKeyPad[current].i,
		j: end.j - directionalKeyPad[current].j,
	}

	if iteration == 0 {
		return 1
	}

	if result, ok := memo[memoKey{current, target, iteration}]; ok {
		return result
	}

	switch {
	case diff.i == 0 && diff.j == 0:
		return shortest('A', 'A', iteration-1, memo)
	case diff.i == 0 && diff.j == 1:
		return shortest('A', '>', iteration-1, memo) + shortest('>', 'A', iteration-1, memo)
	case diff.i == 0 && diff.j == -1:
		return shortest('A', '<', iteration-1, memo) + shortest('<', 'A', iteration-1, memo)
	case diff.i == 1 && diff.j == 0:
		return shortest('A', 'v', iteration-1, memo) + shortest('v', 'A', iteration-1, memo)
	case diff.i == -1 && diff.j == 0:
		return shortest('A', '^', iteration-1, memo) + shortest('^', 'A', iteration-1, memo)
	case diff.i == 0 && diff.j == 2:
		return shortest('A', '>', iteration-1, memo) + shortest('>', '>', iteration-1, memo) + shortest('>', 'A', iteration-1, memo)
	case diff.i == 0 && diff.j == -2:
		return shortest('A', '<', iteration-1, memo) + shortest('<', '<', iteration-1, memo) + shortest('<', 'A', iteration-1, memo)
	case directionalKeyPad[current] == cell{0, 1} && directionalKeyPad[target] == cell{1, 0}:
		return shortest('A', 'v', iteration-1, memo) + shortest('v', '<', iteration-1, memo) + shortest('<', 'A', iteration-1, memo)
	case directionalKeyPad[current] == cell{1, 0} && directionalKeyPad[target] == cell{0, 1}:
		return shortest('A', '>', iteration-1, memo) + shortest('>', '^', iteration-1, memo) + shortest('^', 'A', iteration-1, memo)
	case directionalKeyPad[current] == cell{0, 1} && directionalKeyPad[target] == cell{1, 2}:
		return int(math.Min(
			float64(shortest('A', 'v', iteration-1, memo)+shortest('v', '>', iteration-1, memo)+shortest('>', 'A', iteration-1, memo)),
			float64(shortest('A', '>', iteration-1, memo)+shortest('>', 'v', iteration-1, memo)+shortest('v', 'A', iteration-1, memo)),
		))
	case directionalKeyPad[current] == cell{1, 2} && directionalKeyPad[target] == cell{0, 1}:
		return int(math.Min(
			float64(shortest('A', '<', iteration-1, memo)+shortest('<', '^', iteration-1, memo)+shortest('^', 'A', iteration-1, memo)),
			float64(shortest('A', '^', iteration-1, memo)+shortest('^', '<', iteration-1, memo)+shortest('<', 'A', iteration-1, memo)),
		))
	case directionalKeyPad[current] == cell{0, 2} && directionalKeyPad[target] == cell{1, 1}:
		return int(math.Min(
			float64(shortest('A', 'v', iteration-1, memo)+shortest('v', '<', iteration-1, memo)+shortest('<', 'A', iteration-1, memo)),
			float64(shortest('A', '<', iteration-1, memo)+shortest('<', 'v', iteration-1, memo)+shortest('v', 'A', iteration-1, memo)),
		))
	case directionalKeyPad[current] == cell{1, 1} && directionalKeyPad[target] == cell{0, 2}:
		return int(math.Min(
			float64(shortest('A', '^', iteration-1, memo)+shortest('^', '>', iteration-1, memo)+shortest('>', 'A', iteration-1, memo)),
			float64(shortest('A', '>', iteration-1, memo)+shortest('>', '^', iteration-1, memo)+shortest('^', 'A', iteration-1, memo)),
		))
	case directionalKeyPad[current] == cell{0, 2} && directionalKeyPad[target] == cell{1, 0}:
		return int(math.Min(
			float64(shortest('A', '<', iteration-1, memo)+shortest('<', 'v', iteration-1, memo)+shortest('v', '<', iteration-1, memo)+shortest('<', 'A', iteration-1, memo)),
			float64(shortest('A', 'v', iteration-1, memo)+shortest('v', '<', iteration-1, memo)+shortest('<', '<', iteration-1, memo)+shortest('<', 'A', iteration-1, memo)),
		))
	case directionalKeyPad[current] == cell{1, 0} && directionalKeyPad[target] == cell{0, 2}:
		return int(math.Min(
			float64(shortest('A', '>', iteration-1, memo)+shortest('>', '^', iteration-1, memo)+shortest('^', '>', iteration-1, memo)+shortest('>', 'A', iteration-1, memo)),
			float64(shortest('A', '>', iteration-1, memo)+shortest('>', '>', iteration-1, memo)+shortest('>', '^', iteration-1, memo)+shortest('^', 'A', iteration-1, memo)),
		))
	}

	panic("invalid iteration")
}
