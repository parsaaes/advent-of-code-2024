package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	line := input.ReadBulk()

	part1(line)
	part2(line)
}

func part1(line string) {
	numsStr := strings.Split(line, " ")

	result := []string{}

	blinkingTime := 25

	for x := 0; x < blinkingTime; x++ {
		for i := range numsStr {
			num, _ := strconv.Atoi(numsStr[i])

			if num == 0 {
				result = append(result, "1")
			} else if len(numsStr[i])%2 == 0 {
				num1 := strings.TrimLeft(numsStr[i][:len(numsStr[i])/2], "0")
				num2 := strings.TrimLeft(numsStr[i][len(numsStr[i])/2:], "0")

				if num1 == "" {
					num1 = "0"
				}

				if num2 == "" {
					num2 = "0"
				}

				result = append(result, num1)
				result = append(result, num2)
			} else {
				result = append(result, strconv.Itoa(num*2024))
			}
		}

		numsStr = make([]string, len(result))
		copy(numsStr, result)

		if x == blinkingTime-1 {
			break
		}

		result = []string{}
	}

	fmt.Println(len(result))
}

func part2(line string) {
	numsStr := strings.Split(line, " ")

	occurrences := map[string]int{}

	for i := range numsStr {
		occurrences[numsStr[i]] = 1
	}

	blinkingTime := 75

	for x := 0; x < blinkingTime; x++ {
		tempOccurrences := make(map[string]int, len(occurrences))

		for k, v := range occurrences {
			tempOccurrences[k] = v
		}

		for k, v := range occurrences {
			if v == 0 {
				continue
			}

			num, _ := strconv.Atoi(k)

			if k == "0" {
				tempOccurrences["1"] += v
			} else if len(k)%2 == 0 {
				num1 := strings.TrimLeft(k[:len(k)/2], "0")
				num2 := strings.TrimLeft(k[len(k)/2:], "0")

				if num1 == "" {
					num1 = "0"
				}

				if num2 == "" {
					num2 = "0"
				}

				tempOccurrences[num1] += v
				tempOccurrences[num2] += v
			} else {
				tempOccurrences[strconv.Itoa(num*2024)] += v
			}

			tempOccurrences[k] -= v
		}

		occurrences = tempOccurrences
	}

	total := 0
	for _, v := range occurrences {
		total += v
	}

	fmt.Println(total)
}
