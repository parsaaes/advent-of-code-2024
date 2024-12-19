package main

import (
	"fmt"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	in := input.ReadBulk()

	availableTowels, wantedTowels := parseInput(in)

	part1(availableTowels, wantedTowels)
	part2(availableTowels, wantedTowels)
}

func parseInput(in string) (map[string]struct{}, []string) {
	parts := strings.Split(in, "\n\n")

	availableTowelsStr := parts[0]

	availableTowelsList := strings.Split(availableTowelsStr, ", ")

	availableTowels := map[string]struct{}{}

	for i := range availableTowelsList {
		availableTowels[availableTowelsList[i]] = struct{}{}
	}

	wantedTowels := strings.Fields(parts[1])

	return availableTowels, wantedTowels
}

func part1(availableTowels map[string]struct{}, wantedTowels []string) {
	total := 0

	memo := make(map[string]int)

	for i := range wantedTowels {
		count := findWays(wantedTowels[i], availableTowels, memo)

		if count > 0 {
			total++
		}
	}

	fmt.Println(total)
}

func part2(availableTowels map[string]struct{}, wantedTowels []string) {
	total := 0

	memo := make(map[string]int)

	for i := range wantedTowels {
		count := findWays(wantedTowels[i], availableTowels, memo)

		total += count
	}

	fmt.Println(total)
}

func findWays(towel string, availableTowels map[string]struct{}, memo map[string]int) int {
	if result, ok := memo[towel]; ok {
		return result
	}

	total := 0
	if _, ok := availableTowels[towel]; ok {
		total++
	}

	for i := 1; i < len(towel); i++ {
		if _, ok := availableTowels[towel[:i]]; !ok {
			continue
		}

		total += findWays(towel[i:], availableTowels, memo)
	}

	memo[towel] = total

	return total
}
