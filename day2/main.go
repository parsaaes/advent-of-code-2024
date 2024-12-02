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

	reports := make([][]int, 0)

	for i := range lines {
		levelStrs := strings.Fields(lines[i])

		levels := make([]int, 0)

		for j := range levelStrs {
			level, _ := strconv.Atoi(levelStrs[j])

			levels = append(levels, level)
		}

		reports = append(reports, levels)
	}

	part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	safeReports := 0

	for i := range reports {
		if isSafeReport(reports[i]) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func part2(reports [][]int) {
	safeReports := 0

	for i := range reports {
		// [*]x * * * * *
		if isSafeReport(reports[i][1:]) {
			safeReports++
			continue
		}

		safe := true
		increasing := (reports[i][1] - reports[i][0]) > 0

		for j := 1; j < len(reports[i]); j++ {
			diff := reports[i][j] - reports[i][j-1]
			absDiff := int(math.Abs(float64(diff)))

			if diff == 0 || (diff > 0) != increasing || absDiff > 3 {
				// * * * * * [*]x
				if j == len(reports[i])-1 {
					break
				}

				// * * [*]x * * *
				newReport1 := make([]int, len(reports[i])-1)
				copy(newReport1[:j], reports[i][:j])
				copy(newReport1[j:], reports[i][j+1:])

				// * *x [*] * * *
				newReport2 := make([]int, len(reports[i])-1)
				copy(newReport2[:j-1], reports[i][:j-1])
				copy(newReport2[j-1:], reports[i][j:])

				if isSafeReport(newReport1) || isSafeReport(newReport2) {
					break
				} else {
					safe = false
					break
				}
			}
		}

		if safe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func isSafeReport(report []int) bool {
	safe := true
	increasing := (report[1] - report[0]) > 0

	for j := 1; j < len(report); j++ {
		diff := report[j] - report[j-1]

		if increasing {
			if diff > 3 || diff <= 0 {
				safe = false
				break
			}
		} else {
			if diff < -3 || diff >= 0 {
				safe = false
				break
			}
		}
	}

	return safe
}
