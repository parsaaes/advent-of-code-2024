package main

import (
	"fmt"
	"math"

	"github.com/parsaaes/advent-of-code-2024/input"
)

const (
	inf     = math.MaxInt64
	onTrack = '.'
	end     = 'E'
	wall    = '#'
)

type (
	cell struct {
		i, j int
	}

	cheat struct {
		target cell
		size   int
	}

	queue[T any] []T
)

var (
	right = cell{0, 1}
	left  = cell{0, -1}
	down  = cell{1, 0}
	up    = cell{-1, 0}
)

func (q *queue[T]) push(i T) {
	*q = append(*q, i)
}

func (q *queue[T]) pop() T {
	if len(*q) == 0 {
		panic("empty queue")
	}

	i := (*q)[0]
	*q = (*q)[1:]

	return i
}

func (c cell) isValid(grid [][]byte) bool {
	if c.i < 0 || c.i > len(grid)-1 || c.j < 0 || c.j > len(grid)-1 {
		return false
	}

	return true
}

func allDirections() []cell {
	return []cell{left, right, up, down}
}

func main() {
	grid := input.ReadPixels()

	var endCell cell

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == end {
				endCell = cell{i, j}
			}
		}
	}

	part1(grid, endCell)
	part2(grid, endCell)
}

func part1(grid [][]byte, endCell cell) {
	cheats := findCheats(grid, endCell, 2)

	total := 0

	for k, v := range cheats {
		if k >= 100 {
			total += v
		}
	}

	fmt.Println(total)
}

func part2(grid [][]byte, endCell cell) {
	cheats := findCheats(grid, endCell, 20)

	total := 0

	for k, v := range cheats {
		if k >= 100 {
			total += v
		}
	}

	fmt.Println(total)
}

func findCheats(grid [][]byte, endCell cell, limit int) map[int]int {
	distance, prev := traverse(grid, endCell)
	cheats := make(map[int]int)

	current := endCell

	for {
		possibleCheats := freeRun(grid, current, limit)

		for i := range possibleCheats {
			if (grid[possibleCheats[i].target.i][possibleCheats[i].target.j] == onTrack ||
				grid[possibleCheats[i].target.i][possibleCheats[i].target.j] == end) &&
				possibleCheats[i].size >= 2 {
				save := distance[current] - distance[possibleCheats[i].target] - possibleCheats[i].size
				cheats[save]++
			}
		}

		var ok bool

		current, ok = prev[current]
		if !ok {
			break
		}
	}

	return cheats
}

func traverse(grid [][]byte, start cell) (map[cell]int, map[cell]cell) {
	var (
		distance = make(map[cell]int)
		prev     = make(map[cell]cell)
	)

	for p := 0; p < len(grid); p++ {
		for q := 0; q < len(grid[p]); q++ {
			if grid[p][q] != wall {
				distance[cell{
					i: p,
					j: q,
				}] = inf
			}
		}
	}

	distance[start] = 0

	q := make(queue[cell], 0)

	q.push(start)

	for len(q) > 0 {
		current := q.pop()

		for _, dir := range allDirections() {
			neighbor := cell{
				i: current.i + dir.i,
				j: current.j + dir.j,
			}

			if !neighbor.isValid(grid) {
				continue
			}

			if grid[neighbor.i][neighbor.j] == wall {
				continue
			}

			if distance[current]+1 < distance[neighbor] {
				q.push(neighbor)

				distance[neighbor] = distance[current] + 1
				prev[current] = neighbor
			}
		}
	}

	return distance, prev
}

func freeRun(grid [][]byte, start cell, limit int) []cheat {
	visited := make(map[cell]bool)
	result := make([]cheat, 0)

	q := make(queue[cheat], 0)

	q.push(cheat{
		target: start,
		size:   0,
	})

	for len(q) > 0 {
		current := q.pop()

		if current.size > limit {
			continue
		}

		if visited[current.target] {
			continue
		}

		visited[current.target] = true

		result = append(result, current)

		for _, dir := range allDirections() {
			neighbor := cell{
				i: current.target.i + dir.i,
				j: current.target.j + dir.j,
			}

			if !neighbor.isValid(grid) {
				continue
			}

			q.push(cheat{
				target: neighbor,
				size:   current.size + 1,
			})
		}
	}

	return result
}
