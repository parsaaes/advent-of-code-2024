package main

import (
	"fmt"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type (
	Direction int

	Point struct {
		Y, X int
	}
)

const (
	Right  Direction = iota
	Bottom Direction = iota
	Left   Direction = iota
	Up     Direction = iota
)

func (d Direction) MoveFactor() (Y, X int) {
	switch d {
	case Right:
		return 0, 1
	case Bottom:
		return 1, 0
	case Left:
		return 0, -1
	case Up:
		return -1, 0
	default:
		return 0, 0
	}
}

func (d Direction) RotateRight() Direction {
	return (d + 1) % 4
}

func IsGuard(point byte) (bool, Direction) {
	switch point {
	case '>':
		return true, Right
	case '<':
		return true, Left
	case '^':
		return true, Up
	case 'v':
		return true, Bottom
	default:
		return false, -1
	}
}

func main() {
	board := input.ReadPixels()

	var (
		GuardDirection Direction
		GuardPoint     Point
	)

	for i := range board {
		for j := range board[i] {
			ok, direction := IsGuard(board[i][j])

			if ok {
				GuardDirection = direction
				GuardPoint = Point{i, j}

				break
			}
		}
	}

	visited := part1(board, GuardDirection, GuardPoint)
	part2(board, GuardDirection, GuardPoint, visited)
}

func part1(board [][]byte, GuardDirection Direction, GuardPoint Point) map[Point]struct{} {
	visited := make(map[Point]struct{})

	for {
		visited[GuardPoint] = struct{}{}

		moveY, moveX := GuardDirection.MoveFactor()
		nextY, nextX := GuardPoint.Y+moveY, GuardPoint.X+moveX

		// exiting the board
		if nextY >= len(board) || nextY < 0 {
			break
		}

		// exiting the board
		if nextX >= len(board[0]) || nextX < 0 {
			break
		}

		// hitting obstacle
		if board[nextY][nextX] == '#' {
			GuardDirection = GuardDirection.RotateRight()

			continue
		}

		// going forward
		if board[nextY][nextX] == '.' {
			board[GuardPoint.Y][GuardPoint.X] = '.'

			GuardPoint = Point{nextY, nextX}

			continue
		}
	}

	fmt.Println(len(visited))

	return visited
}

func part2(board [][]byte, InitialGuardDirection Direction, InitialGuardPoint Point, initialVisited map[Point]struct{}) {
	// removing the initial point
	delete(initialVisited, InitialGuardPoint)

	total := 0

	for potentialObstacle, _ := range initialVisited {
		// putting the obstacle
		board[potentialObstacle.Y][potentialObstacle.X] = 'O'

		GuardPoint := InitialGuardPoint
		GuardDirection := InitialGuardDirection

		visited := make(map[Point]map[Direction]struct{})

		for {
			if _, ok := visited[GuardPoint]; !ok {
				visited[GuardPoint] = map[Direction]struct{}{}
			}

			_, ok := visited[GuardPoint][GuardDirection]
			// loop detected
			if ok {
				total++
				break
			}

			visited[GuardPoint][GuardDirection] = struct{}{}

			moveY, moveX := GuardDirection.MoveFactor()
			nextY, nextX := GuardPoint.Y+moveY, GuardPoint.X+moveX

			// exiting the board
			if nextY >= len(board) || nextY < 0 {
				break
			}

			// exiting the board
			if nextX >= len(board[0]) || nextX < 0 {
				break
			}

			// hitting obstacle
			if board[nextY][nextX] == '#' || board[nextY][nextX] == 'O' {
				GuardDirection = GuardDirection.RotateRight()

				continue
			}

			// going forward
			if board[nextY][nextX] == '.' {
				board[GuardPoint.Y][GuardPoint.X] = '.'

				GuardPoint = Point{nextY, nextX}

				continue
			}
		}

		// removing the obstacle
		board[potentialObstacle.Y][potentialObstacle.X] = '.'
	}

	fmt.Println(total)
}
