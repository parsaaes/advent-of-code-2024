package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	lines := input.Read()

	totalSequences := make(map[[4]int]int)

	total := 0
	for i := range lines {
		num, _ := strconv.Atoi(lines[i])

		sequences := make(map[[4]int]int)

		result := calculate(num, 2000, sequences, make([]int, 0))

		for k, v := range sequences {
			totalSequences[k] += v
		}

		total += result
	}

	maximum := math.MinInt64
	for _, v := range totalSequences {
		if v > maximum {
			maximum = v
		}
	}

	fmt.Println(total)
	fmt.Println(maximum)
}

func calculate(num int, limit int, sequences map[[4]int]int, previous []int) int {
	if limit == 0 {
		return num
	}

	lastDigit := num % 10

	num = ((num << 6) ^ num) & (16777216 - 1)
	num = ((num >> 5) ^ num) & (16777216 - 1)
	num = ((num << 11) ^ num) & (16777216 - 1)

	newDigit := num % 10

	diff := newDigit - lastDigit

	previous = append(previous, diff)

	if len(previous) > 4 {
		previous = previous[1:]
	}

	if len(previous) == 4 {
		var key [4]int
		copy(key[:], previous)

		if _, ok := sequences[key]; !ok {
			sequences[key] = newDigit
		}
	}

	return calculate(num, limit-1, sequences, previous)
}
