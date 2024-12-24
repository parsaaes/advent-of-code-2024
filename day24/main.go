package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type Gate struct {
	inputA, inputB string
	operator       string
}

func main() {
	in := input.ReadBulk()

	parts := strings.Split(in, "\n\n")

	inputBits := make(map[string]int)
	gates := make(map[Gate]string)

	for _, val := range strings.Split(parts[0], "\n") {
		k, v := parseValues(val)
		inputBits[k] = v
	}

	for _, val := range strings.Split(parts[1], "\n") {
		k, v := parseGates(val)
		gates[v] = k
	}

	resultBits := make([]string, 0)

	for _, v := range gates {
		if v[0] == 'z' {
			resultBits = append(resultBits, v)
		}
	}

	part1(gates, resultBits, inputBits)
	part2(gates, resultBits)
}

func part1(gates map[Gate]string, resultBits []string, inputBits map[string]int) {
	sort.Strings(resultBits)

	for _, k := range resultBits {
		calculate(k, inputBits, gates)
	}

	result := 0
	for i := 0; i < len(resultBits); i++ {
		result += int(math.Pow(2, float64(i))) * inputBits[fmt.Sprintf("z%02d", i)]
	}

	fmt.Println(result)
}

func part2(gates map[Gate]string, resultBits []string) {
	changes, carry := checkHalfAdder(gates)

	for i := 1; i < len(resultBits)-1; i++ {
		ch, c := checkFullAdder(gates, i, carry)

		changes = append(changes, ch...)

		carry = c
	}

	if carry != fmt.Sprintf("z%02d", len(resultBits)-1) {
		changes = append(changes, []string{fmt.Sprintf("z%02d", len(resultBits)-1), carry}...)

		swapGates(gates, fmt.Sprintf("z%02d", len(resultBits)-1), carry)
	}

	sort.Strings(changes)

	fmt.Println(strings.Join(changes, ","))
}

func checkHalfAdder(gates map[Gate]string) ([]string, string) {
	changes := make([]string, 0)

	firstXOR := Gate{
		inputA:   "x00",
		inputB:   "y00",
		operator: "XOR",
	}

	firstAND := Gate{
		inputA:   "x00",
		inputB:   "y00",
		operator: "AND",
	}

	if gates[firstXOR] != "z00" {
		changes = append(changes, []string{"z00", gates[firstXOR]}...)

		swapGates(gates, "z00", gates[firstXOR])
	}

	return changes, gates[firstAND]
}

func checkFullAdder(gates map[Gate]string, bit int, carry string) ([]string, string) {
	changes := make([]string, 0)

	inputsXOR := Gate{
		inputA:   fmt.Sprintf("x%02d", bit),
		inputB:   fmt.Sprintf("y%02d", bit),
		operator: "XOR",
	}

	inputsAndCarryXOR := findGateByOneInput(gates, gates[inputsXOR], "XOR")

	if inputsAndCarryXOR == nil {
		correctInputsAndCarryXOR := findGateByOneInput(gates, carry, "XOR")

		inputsXORToChange := correctInputsAndCarryXOR.inputA
		if inputsXORToChange == carry {
			inputsXORToChange = correctInputsAndCarryXOR.inputB
		}

		changes = append(changes, []string{inputsXORToChange, gates[inputsXOR]}...)

		swapGates(gates, gates[inputsXOR], inputsXORToChange)

		inputsAndCarryXOR = findGateByOneInput(gates, carry, "XOR")
	}

	inputsAndCarryAND := findGateByOneInput(gates, gates[inputsXOR], "AND")
	if inputsAndCarryAND == nil {
		panic("not implemented") // no case was found in the input
	}

	if gates[*inputsAndCarryXOR] != fmt.Sprintf("z%02d", bit) {
		changes = append(changes, []string{fmt.Sprintf("z%02d", bit), gates[*inputsAndCarryXOR]}...)

		swapGates(gates, fmt.Sprintf("z%02d", bit), gates[*inputsAndCarryXOR])
	}

	carryOfXOR := inputsAndCarryXOR.inputA
	carryOFAND := inputsAndCarryAND.inputA

	if inputsAndCarryXOR.inputA == gates[inputsXOR] {
		carryOfXOR = inputsAndCarryXOR.inputB
		carryOFAND = inputsAndCarryAND.inputB
	}

	if carryOfXOR != carry {
		panic("not implemented") // no case was found in the input
	}

	if carryOFAND != carry {
		panic("not implemented") // no case was found in the input
	}

	inputs := sortInputs(carry, gates[inputsXOR])

	carryANDGate := Gate{
		inputA:   inputs[0],
		inputB:   inputs[1],
		operator: "AND",
	}

	possibleORGate1 := findGateByOneInput(gates, gates[carryANDGate], "OR")
	if possibleORGate1 == nil {
		panic("not implemented") // no case was found in the input
	}

	inputsANDGate := Gate{
		inputA:   fmt.Sprintf("x%02d", bit),
		inputB:   fmt.Sprintf("y%02d", bit),
		operator: "AND",
	}

	possibleORGate2 := findGateByOneInput(gates, gates[inputsANDGate], "OR")
	if possibleORGate2 == nil {
		panic("not implemented") // no case was found in the input
	}

	inputs = sortInputs(gates[carryANDGate], gates[inputsANDGate])

	orGate := Gate{
		inputA:   inputs[0],
		inputB:   inputs[1],
		operator: "OR",
	}

	if orGate != *possibleORGate1 {
		panic("not implemented") // no case was found in the input
	}

	if orGate != *possibleORGate2 {
		panic("not implemented") // no case was found in the input
	}

	return changes, gates[orGate]
}

func swapGates(gates map[Gate]string, x string, y string) {
	g1 := findGateByOutput(gates, x)
	g2 := findGateByOutput(gates, y)

	gates[g1] = y
	gates[g2] = x
}

func sortInputs(in ...string) []string {
	sort.Strings(in)

	return in
}

func findGateByOneInput(gates map[Gate]string, input string, operation string) *Gate {
	for k := range gates {
		if k.operator == operation {
			if k.inputA == input || k.inputB == input {
				return &k
			}
		}
	}

	return nil
}

func findGateByOutput(gates map[Gate]string, output string) Gate {
	for k, v := range gates {
		if v == output {
			return k
		}
	}

	panic("not found")
}

func calculate(key string, calculated map[string]int, gates map[Gate]string) int {
	if val, ok := calculated[key]; ok {
		return val
	}

	var gate Gate
	for k, v := range gates {
		if v == key {
			gate = k
		}
	}

	a := calculate(gate.inputA, calculated, gates)
	b := calculate(gate.inputB, calculated, gates)

	result := 0

	switch gate.operator {
	case "AND":
		if a == 1 && b == 1 {
			result = 1
		}
	case "OR":
		if a == 1 || b == 1 {
			result = 1
		}
	case "XOR":
		if a != b {
			result = 1
		}
	}

	calculated[key] = result

	return result
}

func parseValues(line string) (string, int) {
	subParts := strings.Split(line, ":")

	number, _ := strconv.Atoi(strings.TrimSpace(subParts[1]))

	return subParts[0], number
}

func parseGates(line string) (string, Gate) {
	re := regexp.MustCompile(`(.*) (.*) (.*) -> (.*)`)

	matches := re.FindStringSubmatch(line)

	sortedInputs := sortInputs(matches[1], matches[3])

	return matches[4], Gate{
		inputA:   sortedInputs[0],
		inputB:   sortedInputs[1],
		operator: matches[2],
	}
}
