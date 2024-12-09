package main

import (
	"fmt"
	"strconv"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	in := input.ReadBulk()

	fileID := 0
	diskMap := make([]int, 0)

	for i := range in {
		repeat, _ := strconv.Atoi(string(in[i]))

		if i%2 == 0 {
			// file space
			for p := 0; p < repeat; p++ {
				diskMap = append(diskMap, fileID)
			}

			fileID++
		} else {
			// free space
			for p := 0; p < repeat; p++ {
				diskMap = append(diskMap, -1)
			}
		}
	}

	part1(clone(diskMap))
	part2(clone(diskMap))
}

func clone(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)

	return result
}

func checksum(diskMap []int) int {
	result := 0

	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == -1 {
			continue
		}

		result += i * diskMap[i]
	}

	return result
}

func part1(diskMap []int) {
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] == -1 {
			j := len(diskMap) - 1

			for j >= 0 && diskMap[j] == -1 {
				j--
			}

			diskMap[i] = diskMap[j]
			diskMap = diskMap[:j]
		}
	}

	fmt.Println(checksum(diskMap))
}

func part2(diskMap []int) {
	type block struct {
		id, pos, length int
	}

	files := make([]block, 0)
	blanks := make([]block, 0)

	for i := 0; i < len(diskMap); i++ {
		pos := i
		id := diskMap[i]

		for i < len(diskMap) && diskMap[i] == id {
			i++
		}

		length := i - pos

		i--

		if diskMap[i] == -1 {
			blanks = append(blanks, block{
				id:     diskMap[i],
				pos:    pos,
				length: length,
			})
		} else {
			files = append(files, block{
				id:     diskMap[i],
				pos:    pos,
				length: length,
			})
		}
	}

	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(blanks); j++ {
			// we just need to move left files to right blanks
			if files[i].pos <= blanks[j].pos {
				break
			}

			// found a spot
			if files[i].length <= blanks[j].length {
				for p := 0; p < files[i].length; p++ {
					diskMap[files[i].pos+p] = -1
					diskMap[blanks[j].pos+p] = files[i].id
				}

				blanks[j].pos += files[i].length
				blanks[j].length -= files[i].length

				break
			}
		}
	}

	fmt.Println(checksum(diskMap))
}
