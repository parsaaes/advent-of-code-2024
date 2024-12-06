package input

import (
	"os"
	"strings"
)

const inputFileName = "in"

func Read() []string {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	strData := strings.TrimSpace(string(data))

	return strings.Split(strData, "\n")
}

func ReadPixels() [][]byte {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	strData := strings.TrimSpace(string(data))

	rawLines := strings.Split(strData, "\n")

	result := make([][]byte, 0)

	for i := range rawLines {
		line := make([]byte, 0)

		for j := range rawLines[i] {
			line = append(line, rawLines[i][j])
		}

		result = append(result, line)
	}

	return result
}

func ReadRaw() []string {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func ReadBulk() string {
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	return string(data)
}
