package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type (
	location struct {
		y, x int
	}

	clawMachine struct {
		a, b, prize location
	}
)

func main() {
	in := input.ReadBulk()

	rawClawMachines := strings.Split(in, "\n\n")

	clawMachines := make([]clawMachine, len(rawClawMachines))

	for i := range rawClawMachines {
		rawClawMachineLines := strings.Split(rawClawMachines[i], "\n")

		clm := clawMachine{
			a:     parseLine(rawClawMachineLines[0]),
			b:     parseLine(rawClawMachineLines[1]),
			prize: parseLine(rawClawMachineLines[2]),
		}

		clawMachines[i] = clm
	}

	part1(clawMachines)
	part2(clawMachines)
}

func part1(clawMachines []clawMachine) {
	total := 0

	for i := range clawMachines {
		a, b := solveGreedy(clawMachines[i])

		total += 3*(a) + b
	}

	fmt.Println(total)
}

func part2(clawMachines []clawMachine) {
	/*
		xA*a+xB*b=prizeX
		yA*a+yB*b=prizeY
	*/

	var total int64

	for i := range clawMachines {
		prizeX := clawMachines[i].prize.x + 10000000000000
		prizeY := clawMachines[i].prize.y + 10000000000000

		determinant := clawMachines[i].a.x*clawMachines[i].b.y - clawMachines[i].b.x*clawMachines[i].a.y

		// zero or infinite results
		if determinant == 0 {
			if prizeX%clawMachines[i].b.x == 0 {
				numberOfB := prizeX / clawMachines[i].b.x

				if prizeY == numberOfB*clawMachines[i].b.y {
					total += int64(numberOfB)
				}
			} else if prizeX%clawMachines[i].a.x == 0 {
				numberOfA := prizeX / clawMachines[i].a.x

				if prizeY == numberOfA*clawMachines[i].a.y {
					total += int64(numberOfA * 3)
				}
			}

			continue
		}

		a := float64(prizeX*clawMachines[i].b.y-clawMachines[i].b.x*prizeY) / float64(determinant)
		b := float64(clawMachines[i].a.x*prizeY-prizeX*clawMachines[i].a.y) / float64(determinant)

		// moves should be integers
		if a == float64(int64(a)) && b == float64(int64(b)) {
			total += 3*int64(a) + int64(b)
		}
	}

	fmt.Println(total)
}

func parseLine(line string) location {
	re := regexp.MustCompile(`.*: X[+=]([0-9]+), Y[+=]([0-9]+)`)

	matches := re.FindStringSubmatch(line)

	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])

	return location{
		x: x,
		y: y,
	}
}

func solveGreedy(clm clawMachine) (int, int) {
	a := clm.a
	b := clm.b
	prize := clm.prize

	initialB := prize.x / b.x

	aResult, bResult := 0, 0

	for potentialB := initialB; potentialB >= 0; potentialB-- {
		resultX := potentialB * b.x
		resultY := potentialB * b.y

		if resultX == prize.x && resultY == prize.y {
			return 0, potentialB
		}

		if resultX < prize.x && resultY < prize.y {
			remainingX := prize.x - resultX
			remainingY := prize.y - resultY

			howManyX := remainingX / a.x
			howManyY := remainingY / a.y

			if remainingX%a.x == 0 && remainingY%a.y == 0 && howManyX == howManyY {
				aResult = howManyX
				bResult = potentialB

				return aResult, bResult
			}
		}
	}

	return 0, 0
}
