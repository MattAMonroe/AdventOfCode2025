package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello World\n")
}

type Range struct {
	start int
	end   int
}

func ParseRange(s string) Range {
	splits := strings.Split(s, "-")

	if len(splits) < 2 {
		fmt.Fprintf(os.Stderr, "Range improperly formatted (%s)\n", s)
		return Range{0, 0}
	}

	lower, err := strconv.Atoi(splits[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse range start (%s)\n", splits[0])
		return Range{0, 0}
	}

	upper, err := strconv.Atoi(strings.ReplaceAll(splits[1], "\n", ""))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse range end (%s)\n", splits[1])
		return Range{0, 0}
	}

	return Range{lower, upper}
}

func ParseRanges(s string) []Range {
	splits := strings.Split(s, ",")
	ranges := []Range{}
	for _, split := range splits {
		if s == "" {
			// skip empty strings from newline, trialing comma, etc
			continue
		}
		r := ParseRange(split)
		if r.start == r.end {
			fmt.Fprintf(os.Stderr, "Start and end of range match, (%d, %d)\n", r.start, r.end)
			continue
		}

		ranges = append(ranges, r)
	}

	return ranges
}

func CheckInvalid(num int) bool {
	strNum := strconv.Itoa(num)
	strLen := len(strNum)
	if strLen%2 != 0 {
		// must have equal number of digits on both sides
		return false
	}

	midpoint := strLen / 2

	// fmt.Fprintf(os.Stderr, "Comparing values \"%s\" : \"%s\" \n", strNum[0:midpoint], strNum[midpoint:strLen])

	return strings.Compare(strNum[0:midpoint], strNum[midpoint:strLen]) == 0
}

func FindInvalidIDs(ranges []Range) int {
	sum := 0

	for _, r := range ranges {
		size := r.end - r.start + 1
		for i := range size {
			num := r.start + i
			if CheckInvalid(num) {
				sum += num
			}
		}
	}

	return sum
}
