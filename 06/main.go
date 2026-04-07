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
