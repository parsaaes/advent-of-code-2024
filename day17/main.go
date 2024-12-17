package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/parsaaes/advent-of-code-2024/input"
)

type Computer struct {
	a, b, c int
	program []string
	pc      int
}

func (c *Computer) currentOperator() string {
	return c.program[c.pc]
}

func (c *Computer) currentOperand() string {
	return c.program[c.pc+1]
}

func (c *Computer) findComboOperand() int {
	operand := c.currentOperand()
	operandNum, _ := strconv.Atoi(operand)

	if 0 <= operandNum && operandNum <= 3 {
		return operandNum
	}

	switch operandNum {
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	}

	panic("invalid operand: " + operand)
}

func (c *Computer) findLiteralOperand() int {
	operand := c.currentOperand()
	operandNum, _ := strconv.Atoi(operand)

	return operandNum
}

func (c *Computer) adv() {
	numerator := c.a
	denominator := int(math.Pow(2, float64(c.findComboOperand())))

	c.a = numerator / denominator

	c.pc += 2
}

func (c *Computer) bdv() {
	numerator := c.a
	denominator := int(math.Pow(2, float64(c.findComboOperand())))

	c.b = numerator / denominator

	c.pc += 2
}

func (c *Computer) cdv() {
	numerator := c.a
	denominator := int(math.Pow(2, float64(c.findComboOperand())))

	c.c = numerator / denominator

	c.pc += 2
}

func (c *Computer) bxl() {
	c.b = c.b ^ c.findLiteralOperand()

	c.pc += 2
}

func (c *Computer) bst() {
	c.b = c.findComboOperand() % 8

	c.pc += 2
}

func (c *Computer) jnz() {
	if c.a != 0 {
		c.pc = c.findComboOperand()
	} else {
		c.pc += 2
	}
}

func (c *Computer) bxc() {
	c.b = c.b ^ c.c
	c.pc += 2
}

func (c *Computer) out() string {
	result := strconv.Itoa(c.findComboOperand()%8) + ","

	c.pc += 2

	return result
}

func main() {
	in := input.ReadBulk()

	computer := parseInput(in)

	part1(computer)
	part2(computer)
}

func part1(computer Computer) {
	out := ""

	for computer.pc < len(computer.program) {
		switch computer.currentOperator() {
		case "0":
			computer.adv()
		case "1":
			computer.bxl()
		case "2":
			computer.bst()
		case "3":
			computer.jnz()
		case "4":
			computer.bxc()
		case "5":
			out += computer.out()
		case "6":
			computer.bdv()
		case "7":
			computer.cdv()
		}
	}

	out = strings.TrimSuffix(out, ",")

	fmt.Println(out)
}

// This is based on my input and also examples :(
// After analyzing my program I found out that it's like a loop
// on register A which at the end of each run it will print register B.
// And registers B and C will be calculated based only on register A.
// So this only works on those inputs.
// These inputs will end like this 5 X 3 0.
func part2(computer Computer) {
	// some basic checks to find some possible unsolvable cases sooner
	if !(computer.program[len(computer.program)-1] == "0" &&
		computer.program[len(computer.program)-2] == "3" &&
		computer.program[len(computer.program)-4] == "5" &&
		haveAAsLoopCounter(computer)) {
		panic("cannot solve")
	}

	possibleAs := map[int]struct{}{
		0: {},
	}

	for i := len(computer.program) - 1; i >= 0; i-- {
		whatShouldBePrinted := computer.program[i]
		tempPossibleAs := map[int]struct{}{}

		for possibleABase := range possibleAs {
			for r := 0; r < 8; r++ {
				possibleA := possibleABase + r

				if possibleA == 0 {
					continue
				}

				if whatWillBePrinted(Computer{
					a:       possibleA,
					program: computer.program,
				}) == whatShouldBePrinted {
					tempPossibleAs[8*possibleA] = struct{}{}
				}
			}
		}

		possibleAs = tempPossibleAs
	}

	m := math.MaxInt

	for k := range possibleAs {
		possibleA := k / 8
		if possibleA < m {
			m = possibleA
		}
	}

	fmt.Println(m)
}

func haveAAsLoopCounter(computer Computer) bool {
	for i := 0; i < len(computer.program); i += 2 {
		if computer.program[i] == "0" {
			return true
		}
	}

	return false
}

func whatWillBePrinted(possibleComputer Computer) string {
	out := ""

	for possibleComputer.pc < len(possibleComputer.program) {
		switch possibleComputer.currentOperator() {
		case "0":
			possibleComputer.adv()
		case "1":
			possibleComputer.bxl()
		case "2":
			possibleComputer.bst()
		case "3":
			// skip jump to simulate the single run of the loop
			possibleComputer.pc += 2
		case "4":
			possibleComputer.bxc()
		case "5":
			out += possibleComputer.out()
		case "6":
			possibleComputer.bdv()
		case "7":
			possibleComputer.cdv()
		}
	}

	out = strings.TrimSuffix(out, ",")

	return out
}

func parseInput(in string) Computer {
	re := regexp.MustCompile(`Register A: ([0-9]+)\nRegister B: 0\nRegister C: 0\n\nProgram: (.+)`)

	matches := re.FindStringSubmatch(in)

	registerA, _ := strconv.Atoi(matches[1])
	program := matches[2]

	return Computer{
		a:       registerA,
		program: strings.Split(program, ","),
	}
}
