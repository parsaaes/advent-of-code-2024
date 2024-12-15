package main

import (
	"fmt"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type (
	cell struct{ i, j int }

	stack []cell
)

const (
	robot       = '@'
	empty       = '.'
	wall        = '#'
	box         = 'O'
	bigBoxLeft  = '['
	bigBoxRight = ']'
)

var (
	directionMap = map[byte]cell{
		'>': {i: 0, j: 1},
		'<': {i: 0, j: -1},
		'v': {i: 1, j: 0},
		'^': {i: -1, j: 0},
	}
)

func (s *stack) push(c cell) {
	*s = append(*s, c)
}

func (s *stack) pop() cell {
	if len(*s) == 0 {
		panic("empty stack pop")
	}

	result := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return result
}

func main() {
	in := input.ReadBulk()

	parts := strings.Split(in, "\n\n")

	rawGrid := strings.Split(parts[0], "\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")

	part1(rawGrid, moves)
	part2(rawGrid, moves)
}

func part1(rawGrid []string, moves string) {
	grid := make([][]byte, len(rawGrid))

	var initial cell

	for i := 0; i < len(rawGrid); i++ {
		grid[i] = make([]byte, len(rawGrid[i]))

		for j := 0; j < len(rawGrid[i]); j++ {
			grid[i][j] = rawGrid[i][j]
			if grid[i][j] == robot {
				initial = cell{i, j}
			}
		}
	}

	for i := 0; i < len(moves); i++ {
		dir := directionMap[moves[i]]

		initial = move(initial, dir, grid)
	}

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == box {
				total += 100*i + j
			}
		}
	}

	fmt.Println(total)
}

func part2(rawGrid []string, moves string) {
	grid := make([][]byte, len(rawGrid))

	var initial cell

	for i := 0; i < len(rawGrid); i++ {
		grid[i] = make([]byte, 2*len(rawGrid[i]))

		for j := 0; j < len(rawGrid[i]); j++ {
			switch rawGrid[i][j] {
			case wall:
				grid[i][2*j] = wall
				grid[i][2*j+1] = wall
			case box:
				grid[i][2*j] = bigBoxLeft
				grid[i][2*j+1] = bigBoxRight
			case empty:
				grid[i][2*j] = empty
				grid[i][2*j+1] = empty
			case robot:
				grid[i][2*j] = robot
				grid[i][2*j+1] = empty
				initial = cell{i, 2 * j}
			}
		}
	}

	for i := 0; i < len(moves); i++ {
		dir := directionMap[moves[i]]

		initial = moveBig(initial, dir, grid)
	}

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == bigBoxLeft {
				total += 100*i + j
			}
		}
	}

	fmt.Println(total)
}

func move(initial cell, direction cell, grid [][]byte) cell {
	newCell := cell{initial.i + direction.i, initial.j + direction.j}

	if grid[newCell.i][newCell.j] == empty {
		grid[initial.i][initial.j] = empty
		grid[newCell.i][newCell.j] = robot

		return newCell
	}

	if grid[newCell.i][newCell.j] == wall {
		return initial
	}

	if grid[newCell.i][newCell.j] == box {
		switch direction {
		case cell{0, 1}:
			shiftRight(newCell, grid)
		case cell{0, -1}:
			shiftLeft(newCell, grid)
		case cell{1, 0}:
			shiftDown(newCell, grid)
		case cell{-1, 0}:
			shiftUp(newCell, grid)
		}
	}

	if grid[newCell.i][newCell.j] == empty {
		grid[initial.i][initial.j] = empty
		grid[newCell.i][newCell.j] = robot

		return newCell
	}

	return initial
}

func moveBig(initial cell, direction cell, grid [][]byte) cell {
	newCell := cell{initial.i + direction.i, initial.j + direction.j}

	if grid[newCell.i][newCell.j] == empty {
		grid[initial.i][initial.j] = empty
		grid[newCell.i][newCell.j] = robot

		return newCell
	}

	if grid[newCell.i][newCell.j] == wall {
		return initial
	}

	if grid[newCell.i][newCell.j] == bigBoxLeft || grid[newCell.i][newCell.j] == bigBoxRight {
		switch direction {
		case cell{0, 1}:
			shiftRightBig(newCell, grid)
		case cell{0, -1}:
			shiftLeftBig(newCell, grid)
		case cell{1, 0}:
			shiftDownBig(newCell, grid)
		case cell{-1, 0}:
			shiftUpBig(newCell, grid)
		}
	}

	if grid[newCell.i][newCell.j] == empty {
		grid[initial.i][initial.j] = empty
		grid[newCell.i][newCell.j] = robot

		return newCell
	}

	return initial
}

func shiftLeft(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != box {
		panic("invalid shift")
	}

	j := firstBox.j
	for j > 0 {
		if grid[firstBox.i][j] != box {
			break
		}

		j--
	}

	if grid[firstBox.i][j] == empty {
		grid[firstBox.i][j] = box
		grid[firstBox.i][firstBox.j] = empty
	}
}

func shiftLeftBig(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != bigBoxRight {
		panic("invalid shift")
	}

	j := firstBox.j

	for j > 0 {
		if grid[firstBox.i][j] != bigBoxRight {
			break
		}

		j -= 2
	}

	if grid[firstBox.i][j] == empty {
		for q := j; q < firstBox.j; q++ {
			grid[firstBox.i][q] = grid[firstBox.i][q+1]
		}

		grid[firstBox.i][firstBox.j] = empty
	}
}

func shiftRight(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != box {
		panic("invalid shift")
	}

	j := firstBox.j
	for j < len(grid[firstBox.i])-1 {
		if grid[firstBox.i][j] != box {
			break
		}

		j++
	}

	if grid[firstBox.i][j] == empty {
		grid[firstBox.i][j] = box
		grid[firstBox.i][firstBox.j] = empty
	}
}

func shiftRightBig(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != bigBoxLeft {
		panic("invalid shift")
	}

	j := firstBox.j

	for j < len(grid[firstBox.i])-1 {
		if grid[firstBox.i][j] != bigBoxLeft {
			break
		}

		j += 2
	}

	if grid[firstBox.i][j] == empty {
		for q := j; q > firstBox.j; q-- {
			grid[firstBox.i][q] = grid[firstBox.i][q-1]
		}

		grid[firstBox.i][firstBox.j] = empty
	}
}

func shiftUp(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != box {
		panic("invalid shift")
	}

	i := firstBox.i
	for i > 0 {
		if grid[i][firstBox.j] != box {
			break
		}

		i--
	}

	if grid[i][firstBox.j] == empty {
		grid[i][firstBox.j] = box
		grid[firstBox.i][firstBox.j] = empty
	}
}

func shiftUpBig(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != bigBoxLeft && grid[firstBox.i][firstBox.j] != bigBoxRight {
		panic("invalid shift")
	}

	// let's identify the box by its left side
	if grid[firstBox.i][firstBox.j] == bigBoxRight {
		firstBox.j--
	}

	biggerBox := map[cell]struct{}{}

	boxStack := make(stack, 0)

	boxStack.push(firstBox)

	for len(boxStack) != 0 {
		b := boxStack.pop()

		biggerBox[b] = struct{}{}

		/*
			[]
			[]
		*/
		if grid[b.i-1][b.j] == bigBoxLeft {
			boxStack.push(cell{
				i: b.i - 1,
				j: b.j,
			})
		}

		/*
			[]
			 []
		*/
		if grid[b.i-1][b.j-1] == bigBoxLeft {
			boxStack.push(cell{
				i: b.i - 1,
				j: b.j - 1,
			})
		}

		/*
			 []
			[]
		*/
		if grid[b.i-1][b.j+1] == bigBoxLeft {
			boxStack.push(cell{
				i: b.i - 1,
				j: b.j + 1,
			})
		}
	}

	hasObstacle := false

	for b := range biggerBox {
		if grid[b.i-1][b.j] == wall || grid[b.i-1][b.j+1] == wall {
			hasObstacle = true
			break
		}
	}

	if hasObstacle {
		return
	}

	for b := range biggerBox {
		grid[b.i][b.j] = empty
		grid[b.i][b.j+1] = empty
	}

	for b := range biggerBox {
		grid[b.i-1][b.j] = bigBoxLeft
		grid[b.i-1][b.j+1] = bigBoxRight
	}
}

func shiftDown(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != box {
		panic("invalid shift")
	}

	i := firstBox.i
	for i < len(grid)-1 {
		if grid[i][firstBox.j] != box {
			break
		}

		i++
	}

	if grid[i][firstBox.j] == empty {
		grid[i][firstBox.j] = box
		grid[firstBox.i][firstBox.j] = empty
	}
}

func shiftDownBig(firstBox cell, grid [][]byte) {
	if grid[firstBox.i][firstBox.j] != bigBoxLeft && grid[firstBox.i][firstBox.j] != bigBoxRight {
		panic("invalid shift")
	}

	// let's identify the box by its left side
	if grid[firstBox.i][firstBox.j] == bigBoxRight {
		firstBox.j--
	}

	biggerBox := map[cell]struct{}{}

	boxStack := make(stack, 0)

	boxStack.push(firstBox)

	for len(boxStack) != 0 {
		b := boxStack.pop()

		biggerBox[b] = struct{}{}

		/*
			[]
			[]
		*/
		if grid[b.i+1][b.j] == bigBoxLeft {
			boxStack.push(cell{
				i: b.i + 1,
				j: b.j,
			})
		}

		/*
			[]
			 []
		*/
		if grid[b.i+1][b.j+1] == bigBoxLeft {
			boxStack.push(cell{
				i: b.i + 1,
				j: b.j + 1,
			})
		}

		/*
			 []
			[]
		*/
		if grid[b.i+1][b.j-1] == bigBoxLeft {
			boxStack.push(cell{
				i: b.i + 1,
				j: b.j - 1,
			})
		}
	}

	hasObstacle := false

	for b := range biggerBox {
		if grid[b.i+1][b.j] == wall || grid[b.i+1][b.j+1] == wall {
			hasObstacle = true
			break
		}
	}

	if hasObstacle {
		return
	}

	for b := range biggerBox {
		grid[b.i][b.j] = empty
		grid[b.i][b.j+1] = empty
	}

	for b := range biggerBox {
		grid[b.i+1][b.j] = bigBoxLeft
		grid[b.i+1][b.j+1] = bigBoxRight
	}
}

func showGrid(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s", string(grid[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}
