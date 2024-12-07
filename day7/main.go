package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	lines := input.Read()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")

		testResult, _ := strconv.Atoi(parts[0])

		numsStr := strings.Fields(parts[1])

		nums := make([]int, len(numsStr))

		for i, s := range numsStr {
			nums[i], _ = strconv.Atoi(s)
		}

		operatorsLen := len(nums) - 1

		operatorCombinations := int(math.Pow(2, float64(operatorsLen)))

	outer:
		for i := operatorCombinations - 1; i >= 0; i-- {
			result := nums[0]

			// iterating over bits to calculate based on current combination
			for j := 0; j < operatorsLen; j++ {
				testBit := i >> (operatorsLen - 1 - j) // starting from the left most bit

				if testBit%2 == 0 {
					result += nums[j+1]
				} else {
					result *= nums[j+1]
				}

				if result > testResult {
					continue outer
				}

				if result == testResult && j == operatorsLen-1 {
					total += testResult
					break outer
				}
			}
		}
	}

	fmt.Println(total)
}

func part2(lines []string) {
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")

		testResult, _ := strconv.Atoi(parts[0])

		numStrs := strings.Fields(parts[1])

		nums := make([]int, len(numStrs))

		for i, s := range numStrs {
			nums[i], _ = strconv.Atoi(s)
		}

		operatorLen := len(nums) - 1

		operatorCombinations := int(math.Pow(3, float64(operatorLen)))

	outer:
		for i := operatorCombinations - 1; i >= 0; i-- {
			result := nums[0]

			for j := 0; j < operatorLen; j++ {
				testDigit := i / int(math.Pow(3, float64(operatorLen-1-j)))

				if testDigit%3 == 0 {
					result += nums[j+1]
				} else if testDigit%3 == 1 {
					result *= nums[j+1]
				} else {
					result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(nums[j+1]))
				}

				if result > testResult {
					continue outer
				}

				if result == testResult && j == operatorLen-1 {
					total += testResult
					break outer
				}
			}
		}
	}

	fmt.Println(total)
}
