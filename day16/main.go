package main

import (
	"fmt"
	"math"

	"github.com/parsaaes/advent-of-code-2024/input"
)

const (
	inf   = math.MaxInt64
	wall  = '#'
	empty = '.'
	start = 'S'
	end   = 'E'
)

type (
	point struct {
		i, j int
	}

	cell struct {
		point     point
		direction point
	}

	queue []cell

	stack []cell
)

var (
	right = point{0, 1}
	left  = point{0, -1}
	down  = point{1, 0}
	up    = point{-1, 0}
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

func allDirections() []point {
	return []point{left, right, up, down}
}

func possibleDirections(dir point) []point {
	switch dir {
	case right:
		return []point{up, down, right}
	case left:
		return []point{up, down, left}
	case up:
		return []point{left, right, up}
	case down:
		return []point{left, right, down}
	}

	panic("invalid direction")
}

func main() {
	grid := input.ReadPixels()

	var (
		startCell cell
		endPoint  point
		distance  = make(map[cell]int)
		prev      = make(map[cell]map[cell]struct{})
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case start:
				startCell = cell{
					point:     point{i, j},
					direction: right,
				}
			case end:
				endPoint = point{i, j}
				fallthrough
			case empty:
				for _, dir := range allDirections() {
					distance[cell{
						point:     point{i, j},
						direction: dir,
					}] = inf
				}
			}
		}
	}

	q := make(queue, 0)

	q.push(startCell)

	for len(q) > 0 {
		current := q.pop()

		for _, dir := range possibleDirections(current.direction) {
			neighbor := cell{
				point:     point{current.point.i + dir.i, current.point.j + dir.j},
				direction: dir,
			}

			if grid[neighbor.point.i][neighbor.point.j] == wall {
				continue
			}

			score := 1

			if dir != current.direction {
				score += 1000
			}

			if distance[current]+score < distance[neighbor] {
				q.push(neighbor)

				distance[neighbor] = distance[current] + score
				prev[neighbor] = map[cell]struct{}{
					current: {},
				}
			}

			if distance[current]+score == distance[neighbor] {
				if prev[neighbor] == nil {
					prev[neighbor] = map[cell]struct{}{}
				}

				prev[neighbor][current] = struct{}{}
			}
		}
	}

	bestPathEnds := make([]cell, 0)

	lowest := inf
	for _, dir := range allDirections() {
		end := cell{
			point: point{
				i: endPoint.i,
				j: endPoint.j,
			},
			direction: dir,
		}

		dist := distance[end]

		if dist < lowest {
			lowest = dist
			bestPathEnds = []cell{end}
		} else if dist == lowest {
			bestPathEnds = append(bestPathEnds, end)
		}
	}

	goodSeats := map[point]struct{}{}

	for _, end := range bestPathEnds {
		s := make(stack, 0)

		s.push(end)

		for len(s) != 0 {
			c := s.pop()

			goodSeats[point{c.point.i, c.point.j}] = struct{}{}

			for pr := range prev[c] {
				s.push(pr)
			}
		}
	}

	fmt.Println(lowest)
	fmt.Println(len(goodSeats))
}
