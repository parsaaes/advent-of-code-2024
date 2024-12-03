package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	raw := input.ReadBulk()

	part1(raw)
	part2(raw)
}

func part1(raw string) {
	sum := 0

	for i := 0; i < len(raw); i++ {
		nextMul := strings.Index(raw[i:], "mul")
		mulIndex := nextMul + i
		i += nextMul + 3

		if nextMul == -1 {
			break
		}

		open := strings.Index(raw[mulIndex:], "(") + mulIndex
		end := strings.Index(raw[mulIndex:], ")") + mulIndex

		if !(open == mulIndex+3 && end > open+3) {
			continue
		}

		potentialNums := strings.Split(raw[open+1:end], ",")

		if len(potentialNums) != 2 {
			continue
		}

		left, err := strconv.Atoi(potentialNums[0])
		if err != nil {
			continue
		}

		right, err := strconv.Atoi(potentialNums[1])
		if err != nil {
			continue
		}

		sum += left * right
	}

	fmt.Println(sum)
}

func part2(raw string) {
	sum := 0

	enabled := true
	enableDisableInstruction := 0

	for i := 0; i < len(raw); i++ {
		nextMul := strings.Index(raw[i:], "mul")
		mulIndex := nextMul + i
		i += nextMul + 3

		lastDo := strings.LastIndex(raw[enableDisableInstruction:i], "do()")
		lastDont := strings.LastIndex(raw[enableDisableInstruction:i], "don't()")

		if lastDo != -1 {
			lastDo += +enableDisableInstruction
			enableDisableInstruction = lastDo + 4
			enabled = true
		}

		if lastDont != -1 {
			lastDont += enableDisableInstruction
		}

		if lastDont > lastDo {
			enableDisableInstruction = lastDont + 7
			enabled = false
		}

		if !enabled {
			continue
		}

		if nextMul == -1 {
			break
		}

		open := strings.Index(raw[mulIndex:], "(") + mulIndex
		end := strings.Index(raw[mulIndex:], ")") + mulIndex

		if !(open == mulIndex+3 && end > open+3) {
			continue
		}

		potentialNums := strings.Split(raw[open+1:end], ",")

		if len(potentialNums) != 2 {
			continue
		}

		left, err := strconv.Atoi(potentialNums[0])
		if err != nil {
			continue
		}

		right, err := strconv.Atoi(potentialNums[1])
		if err != nil {
			continue
		}

		sum += left * right
	}

	fmt.Println(sum)
}
