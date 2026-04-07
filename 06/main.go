package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello World!")
}

type Operation int

const (
	OpAdd Operation = iota
	OpMult
	OpUnknown
)

type Problem struct {
	Numbers []int
	Op      Operation
}

func ParseGrid(content string) []Problem {
	lines := strings.Split(content, "\n")
	problems := []Problem{}
	for _, line := range lines {
		corrected := line
		for strings.Contains(corrected, "  ") {
			corrected = strings.ReplaceAll(corrected, "  ", " ")
		}

		if strings.ReplaceAll(corrected, " ", "") == "" {
			continue
		}

		corrected = strings.Trim(corrected, " ")

		nums := strings.Split(corrected, " ")
		for i, entry := range nums {
			if strings.ReplaceAll(entry, " ", "") == "" {
				continue
			}

			if len(problems) <= i {
				problems = append(problems, Problem{[]int{}, OpUnknown})
			}

			num, err := strconv.Atoi(entry)
			if err != nil {
				if entry == "*" {
					problems[i].Op = OpMult
				} else if entry == "+" {
					problems[i].Op = OpAdd
				} else {
					fmt.Fprintf(os.Stderr, "Failed to parse entry %s: %v\n", entry, err)
					continue
				}
				continue
			}
			problems[i].Numbers = append(problems[i].Numbers, num)
		}
	}
	return problems
}

func ParseVertical(content string) []Problem {
	problems := []Problem{}

	grid := parseGrid(content)

	opRow := len(grid) - 1
	for i := range len(grid[0]) - 1 {
		switch grid[opRow][i] {
		case "*":
			nextOp := findNextOpOffset(grid[opRow][i+1:]) + 1 + i
			if nextOp == i {
				nextOp = len(grid[opRow]) + 1
			}
			spread := nextOp - 1 - i
			problem := Problem{[]int{}, OpMult}
			for j := range spread {
				col := nextOp - 2 - j
				num := readIntFromColumn(grid, col)
				problem.Numbers = append(problem.Numbers, num)
			}
			problems = append(problems, problem)
		case "+":
			nextOp := findNextOpOffset(grid[opRow][i+1:]) + 1 + i
			if nextOp == i {
				nextOp = len(grid[opRow]) + 1
			}
			spread := nextOp - 1 - i
			problem := Problem{[]int{}, OpAdd}
			for j := range spread {
				col := nextOp - 2 - j
				num := readIntFromColumn(grid, col)
				problem.Numbers = append(problem.Numbers, num)
			}
			problems = append(problems, problem)
		default:
			continue
		}
	}

	return problems
}

func parseGrid(content string) [][]string {
	lines := strings.Split(content, "\n")
	grid := [][]string{}

	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		slice := []string{}
		for _, digit := range line {
			slice = append(slice, string(digit))
		}

		grid = append(grid, slice)
	}

	return grid
}

func findNextOpOffset(row []string) int {
	for i := range len(row) {
		if row[i] == "*" || row[i] == "+" {
			return i
		}
	}
	return -1
}

func readIntFromColumn(grid [][]string, col int) int {
	// read the vertical column into an integer and return it
	numStr := ""
	for i := range len(grid) - 1 {
		numStr += grid[i][col]
	}

	numStr = strings.Trim(numStr, " ")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse column %d: %s\n", col, numStr)
	}

	return num
}

func ComputeMath(problems []Problem) int {
	sum := 0
	for _, problem := range problems {
		value := 0

		switch problem.Op {
		case OpMult:
			value = problem.Numbers[0]
			for _, p := range problem.Numbers[1:] {
				value *= p
			}
		case OpAdd:
			for _, p := range problem.Numbers {
				value += p
			}
		default:
			fmt.Fprintf(os.Stderr, "Unknown Operation\n")
		}

		sum += value
	}

	return sum
}
