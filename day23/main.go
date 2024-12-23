package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

func main() {
	lines := input.Read()

	connectionGraph := make(map[string]map[string]struct{})

	for i := range lines {
		computers := strings.Split(lines[i], "-")

		addToNetwork(connectionGraph, computers)
	}

	part1(connectionGraph)
	part2(connectionGraph)
}

func part1(graph map[string]map[string]struct{}) {
	fullGraphs := findFullGraphs(graph, 3)[3]

	total := 0
	for k := range fullGraphs {
		if k[0] == 't' || strings.Contains(k, ",t") {
			total++
		}
	}

	fmt.Println(total)
}

func part2(graph map[string]map[string]struct{}) {
	fullGraphs := findFullGraphs(graph, len(graph))

	maximum := math.MinInt64

	for k := range fullGraphs {
		if maximum < k {
			maximum = k
		}
	}

	for k := range fullGraphs[maximum] {
		fmt.Println(k)
	}
}

func marshalFullGraph(graph map[string]struct{}) string {
	result := make([]string, 0)

	for k := range graph {
		result = append(result, k)
	}

	sort.Strings(result)

	return strings.Join(result, ",")
}

func findFullGraphs(graph map[string]map[string]struct{}, limit int) map[int]map[string]struct{} {
	size := 1

	result := make(map[int]map[string]struct{})
	result[size] = make(map[string]struct{})

	// creating full graphs of size 1 (single nodes)
	for node := range graph {
		result[size][marshalFullGraph(map[string]struct{}{node: {}})] = struct{}{}
	}

	size++

	for size <= limit {
		for node := range graph {
			for smallerFG := range result[size-1] {
				if strings.Contains(smallerFG, node) {
					continue
				}

				shouldAdd := true
				fgToAdd := map[string]struct{}{}
				smallerFGNodes := strings.Split(smallerFG, ",")

				for _, smallerFGNode := range smallerFGNodes {
					if _, ok := graph[smallerFGNode][node]; !ok {
						shouldAdd = false
						break
					}

					fgToAdd[smallerFGNode] = struct{}{}
				}

				if shouldAdd {
					if result[size] == nil {
						result[size] = make(map[string]struct{})
					}

					fgToAdd[node] = struct{}{}

					result[size][marshalFullGraph(fgToAdd)] = struct{}{}
				}
			}
		}

		size++
	}

	return result
}

func addToNetwork(network map[string]map[string]struct{}, pair []string) {
	if _, ok := network[pair[0]]; !ok {
		network[pair[0]] = make(map[string]struct{})
	}

	if _, ok := network[pair[1]]; !ok {
		network[pair[1]] = make(map[string]struct{})
	}

	network[pair[0]][pair[1]] = struct{}{}
	network[pair[1]][pair[0]] = struct{}{}
}
