package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/parsaaes/advent-of-code-2024/input"
)

const (
	gridY   = 103
	gridX   = 101
	seconds = 100
)

type cell struct {
	y, x int
}

func main() {
	lines := input.Read()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	counts := map[cell]int{}

	for i := range lines {
		initialState, velocity := parseLine(lines[i])

		final := move(initialState, velocity, seconds)

		// skip the middle
		if final.x == gridX/2 || final.y == gridY/2 {
			continue
		}

		counts[final]++
	}

	var a, b, c, d int

	for k, v := range counts {
		switch {
		case k.x < gridX/2 && k.y < gridY/2:
			a += v
		case k.x > gridX/2 && k.y < gridY/2:
			b += v
		case k.x < gridX/2 && k.y > gridY/2:
			c += v
		case k.x > gridX/2 && k.y > gridY/2:
			d += v
		}
	}

	fmt.Println(a * b * c * d)
}

func part2(lines []string) {
	f, err := os.OpenFile("log",
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()

	// empty the file
	if _, err = f.WriteString(""); err != nil {
		panic(err)
	}

	// by observing the initial ~300 you can find patterns to iterate faster
	for i := 0; i <= 10000; i += 1 {
		str := fmt.Sprintf("====== iteration[%d] ======\n", i)

		var grid [gridY][gridX]int

		for x := range lines {
			initialState, velocity := parseLine(lines[x])

			final := move(initialState, velocity, i)

			grid[final.y][final.x]++
		}

		for p := range grid {
			for q := range grid[p] {
				if grid[p][q] == 0 {
					str += "ـ"
				} else {
					str += "█"
				}
			}

			str += "\n"
		}

		str += "\n"

		if _, err = f.WriteString(str); err != nil {
			panic(err)
		}
	}
}

func parseLine(line string) (cell, cell) {
	re := regexp.MustCompile(`p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)

	matches := re.FindStringSubmatch(line)

	pX, _ := strconv.Atoi(matches[1])
	pY, _ := strconv.Atoi(matches[2])

	vX, _ := strconv.Atoi(matches[3])
	vY, _ := strconv.Atoi(matches[4])

	return cell{pY, pX}, cell{vY, vX}
}

func move(initialState cell, velocity cell, seconds int) cell {
	y := initialState.y + velocity.y*seconds
	x := initialState.x + velocity.x*seconds

	if y < 0 {
		shift := ((-1*y)/gridY + 1) * gridY
		y += shift
		y %= gridY
	}

	if y >= gridY {
		shift := ((y) / gridY) * gridY
		y -= shift
		y %= gridY
	}

	if x < 0 {
		shift := ((-1*x)/gridX + 1) * gridX
		x += shift
		x %= gridX
	}

	if x >= gridX {
		shift := ((x) / gridX) * gridX
		x -= shift
		x %= gridX
	}

	return cell{y, x}
}
