package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	lines := input.ReadBulk()

	parts := strings.Split(lines, "\n\n")

	rules := make(map[int][]int)
	rawRules := strings.Fields(parts[0])

	for i := range rawRules {
		rawRule := strings.Split(rawRules[i], "|")

		a, _ := strconv.Atoi(rawRule[0])
		b, _ := strconv.Atoi(rawRule[1])

		rules[a] = append(rules[a], b)
	}

	prints := make([][]int, 0)
	printLines := strings.Fields(parts[1])

	for i := range printLines {
		numStrs := strings.Split(printLines[i], ",")

		numInts := make([]int, len(numStrs))

		for j := range numStrs {
			num, _ := strconv.Atoi(numStrs[j])
			numInts[j] = num
		}

		prints = append(prints, numInts)
	}

	incorrectPrints := part1(rules, prints)
	part2(rules, incorrectPrints)
}

func part1(rules map[int][]int, prints [][]int) [][]int {
	sum := 0

	incorrect := make([][]int, 0)

	for i := range prints {
		pageOrders := map[int]int{}

		pages := prints[i]

		for j := range pages {
			pageOrders[pages[j]] = j
		}

		isOk := true

	outer:
		for k, v := range pageOrders {
			befores := rules[k]

			for _, before := range befores {
				index, ok := pageOrders[before]

				if !ok {
					continue
				}

				if v > index {
					isOk = false
					break outer
				}
			}
		}

		if isOk {
			sum += pages[len(pages)/2]
		} else {
			incorrect = append(incorrect, prints[i])
		}
	}

	fmt.Println(sum)

	return incorrect
}

func part2(rules map[int][]int, prints [][]int) {
	sum := 0

	for p := range prints {
		sort.Slice(prints[p], func(i, j int) bool {
			rule, ok := rules[prints[p][i]]

			if !ok {
				return false
			}

			for _, q := range rule {
				if q == prints[p][j] {
					return true
				}
			}

			return false
		})

		sum += prints[p][len(prints[p])/2]
	}

	fmt.Println(sum)
}
