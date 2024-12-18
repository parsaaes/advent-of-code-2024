package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

const (
	inf = math.MaxInt64
)

type (
	cell struct {
		i, j int
	}

	queue []cell
)

var (
	right = cell{0, 1}
	left  = cell{0, -1}
	down  = cell{1, 0}
	up    = cell{-1, 0}
)

func (q *queue) push(c cell) {
	*q = append(*q, c)
}

func (q *queue) pop() cell {
	if len(*q) == 0 {
		panic("empty queue")
	}

	c := (*q)[0]
	*q = (*q)[1:]

	return c
}

func allDirections() []cell {
	return []cell{left, right, up, down}
}

func main() {
	lines := input.Read()

	var grid [71][71]byte

	for i := 0; i < 1024; i++ {
		line := parseLine(lines[i])

		grid[line.i][line.j] = '#'
	}

	fmt.Println(findCost(grid))

	for i := 1024; i < len(lines); i++ {
		c := parseLine(lines[i])

		grid[c.i][c.j] = '#'

		cost := findCost(grid)

		if cost == inf {
			fmt.Println(lines[i])
			break
		}
	}
}

func findCost(grid [71][71]byte) int {
	var (
		start    = cell{0, 0}
		end      = cell{len(grid) - 1, len(grid) - 1}
		distance = make(map[cell]int)
	)

	for p := 0; p < len(grid); p++ {
		for q := 0; q < len(grid[p]); q++ {
			if grid[p][q] == 0 {
				distance[cell{
					i: p,
					j: q,
				}] = inf
			}
		}
	}

	distance[start] = 0

	q := make(queue, 0)

	q.push(start)

	for len(q) > 0 {
		current := q.pop()

		for _, dir := range allDirections() {
			neighbor := cell{
				i: current.i + dir.i,
				j: current.j + dir.j,
			}

			if neighbor.i < 0 || neighbor.i > len(grid)-1 || neighbor.j < 0 || neighbor.j > len(grid)-1 {
				continue
			}

			if grid[neighbor.i][neighbor.j] == '#' {
				continue
			}

			if distance[current]+1 < distance[neighbor] {
				q.push(neighbor)

				distance[neighbor] = distance[current] + 1
			}
		}
	}

	return distance[end]
}

func parseLine(line string) cell {
	cellNums := strings.Split(line, ",")

	x, _ := strconv.Atoi(cellNums[0])
	y, _ := strconv.Atoi(cellNums[1])

	return cell{y, x}
}
