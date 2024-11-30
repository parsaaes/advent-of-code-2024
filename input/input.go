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
