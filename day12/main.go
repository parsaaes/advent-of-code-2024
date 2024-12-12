package main

import (
	"fmt"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type (
	point struct {
		i, j int
	}

	cell struct {
		point
		val     byte
		visited bool
	}

	queue []*cell

	region struct {
		perimeter, area, sides int
		borderGrid             [][]byte
	}
)

const (
	downFacingFence  = 'v'
	upFacingFence    = '^'
	rightFacingFence = '<'
	leftFacingFence  = '>'
	visitedFence     = 'x'
)

func (r *region) findSides() {
	for i := range r.borderGrid {
		for j := range r.borderGrid[i] {
			switch r.borderGrid[i][j] {
			case downFacingFence, upFacingFence:
				r.traverseSideHorizontally(point{i, j}, r.borderGrid[i][j])
			case rightFacingFence, leftFacingFence:
				r.traverseSideVertically(point{i, j}, r.borderGrid[i][j])
			default:
				continue
			}

			r.sides++
		}
	}
}

func (r *region) traverseSideHorizontally(pnt point, fenceType byte) {
	for x := pnt.j; x < len(r.borderGrid[pnt.i]); x += 2 {
		if r.borderGrid[pnt.i][x] == fenceType {
			r.borderGrid[pnt.i][x] = visitedFence
		} else {
			break
		}
	}
}

func (r *region) traverseSideVertically(pnt point, fenceType byte) {
	for x := pnt.i; x < len(r.borderGrid); x += 2 {
		if r.borderGrid[x][pnt.j] == fenceType {
			r.borderGrid[x][pnt.j] = visitedFence
		} else {
			break
		}
	}
}

func (r *region) addFence(pnt, dir point) {
	fenceMap := map[point]byte{
		point{1, 0}:  '^',
		point{-1, 0}: 'v',
		point{0, 1}:  '<',
		point{0, -1}: '>',
	}

	r.borderGrid[pnt.i+dir.i][pnt.j+dir.j] = fenceMap[dir]
}

func (q *queue) push(c *cell) {
	*q = append(*q, c)
}

func (q *queue) pop() *cell {
	if len(*q) == 0 {
		panic("empty queue")
	}

	c := (*q)[0]
	*q = (*q)[1:]

	return c
}

func main() {
	rawGrid := input.ReadPixels()

	grid := make([][]*cell, len(rawGrid))

	for i := range rawGrid {
		grid[i] = make([]*cell, len(rawGrid[i]))

		for j := range rawGrid[i] {
			grid[i][j] = &cell{
				point: point{
					i: i,
					j: j,
				},
				val: rawGrid[i][j],
			}
		}
	}

	regions := make([]region, 0)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j].visited {
				continue
			}

			source := grid[i][j]

			/* border grids works like this:
			          .-.-.
			AA        |A|A|
			AA  =>    .-.-.
					  |A|A|
				      .-.-.
			*/

			r := region{
				borderGrid: make([][]byte, 2*len(grid)+1),
			}

			for p := range r.borderGrid {
				r.borderGrid[p] = make([]byte, 2*len(grid[i])+1)
			}

			q := make(queue, 0)

			q.push(source)

			for len(q) > 0 {
				c := q.pop()

				if c.visited {
					continue
				}

				c.visited = true

				r.area++

				directions := []point{
					{0, 1}, {1, 0}, {0, -1}, {-1, 0},
				}

				for _, d := range directions {
					potentialI := c.i + d.i
					potentialJ := c.j + d.j

					cellPositionInBorderGrid := point{
						i: 2*c.i + 1,
						j: 2*c.j + 1,
					}

					if potentialI < 0 || potentialJ < 0 ||
						potentialI >= len(grid) || potentialJ >= len(grid[0]) {
						r.addFence(cellPositionInBorderGrid, d)
						r.perimeter++
						continue
					}

					potentialCell := grid[potentialI][potentialJ]

					if potentialCell.val == source.val {
						q.push(potentialCell)
					} else {
						r.addFence(cellPositionInBorderGrid, d)
						r.perimeter++
					}
				}
			}

			r.findSides()

			regions = append(regions, r)
		}
	}

	part1Total, part2Total := 0, 0

	for i := range regions {
		part1Total += regions[i].area * regions[i].perimeter
		part2Total += regions[i].area * regions[i].sides
	}

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}
