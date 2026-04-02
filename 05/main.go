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

type Range struct {
	start int
	end   int
}

type Ranges struct {
	r []Range
}

func (r Ranges) IsIncluded(num int) bool {
	for _, spread := range r.r {
		if num >= spread.start && num <= spread.end {
			return true
		}
	}

	return false
}

func ParseInput(content string) (Ranges, []int) {
	ranges := []Range{}
	nums := []int{}

	splits := strings.Split(content, "\n")
	mode := 0
	for _, line := range splits {
		if mode == 0 {
			if line == "" {
				mode = 1
				continue
			}
			s := strings.Split(line, "-")
			if len(s) != 2 {
				fmt.Fprintf(os.Stderr, "Invalid range input: %s\n", line)
				continue
			}
			start, err := strconv.Atoi(s[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid start value: %s from %s\n", s[0], line)
			}

			end, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid end value: %s from %s\n", s[1], line)
			}

			ranges = append(ranges, Range{start, end})
		}
		if mode == 1 {
			if line == "" {
				continue
			}
			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid num value: %s\n", line)
			}

			nums = append(nums, num)
		}
	}

	return Ranges{ranges}, nums
}

func FindIncludedValues(r Ranges, v []int) int {
	count := 0

	for _, num := range v {
		if r.IsIncluded(num) {
			count++
		}
	}

	return count
}
