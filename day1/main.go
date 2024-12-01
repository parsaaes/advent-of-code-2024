package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	lines := input.Read()

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i := range lines {
		row := strings.Fields(lines[i])

		leftID, _ := strconv.Atoi(row[0])
		rightID, _ := strconv.Atoi(row[1])

		left[i] = leftID
		right[i] = rightID
	}

	part1(left, right)
	part2(left, right)
}

func part1(left, right []int) {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(totalDistance)
}

func part2(left, right []int) {
	occurrences := map[int]int{}

	for i := range right {
		occurrences[right[i]] += 1
	}

	similarityScore := 0

	for i := range left {
		similarityScore += left[i] * occurrences[left[i]]
	}

	fmt.Println(similarityScore)
}
