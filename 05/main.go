package main

import (
	"fmt"
	"math"
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

func (r Range) String() string {
	return fmt.Sprintf("[%d - %d]", r.start, r.end)
}

func (r Range) IsIncluded(num int) bool {
	if num >= r.start && num <= r.end {
		return true
	}
	return false
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

func (r Ranges) String() string {
	content := ""
	for _, spread := range r.r {
		content += fmt.Sprintf("%d - %d => %d\n", spread.start, spread.end, (spread.end - spread.start + 1))
	}

	return content
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

func CompactRanges(r Ranges) Ranges {
	ranges := r.r
	diff := len(ranges)
	for diff > 0 {
		beforeCount := len(ranges)
		rangesWrap := CompactRangesSingle(ranges)
		afterCount := len(rangesWrap.r)
		diff = beforeCount - afterCount
		ranges = rangesWrap.r
	}

	return Ranges{ranges}
}

func CompactRangesSingle(r []Range) Ranges {
	newRanges := []*Range{}
	for _, oldRange := range r {
		found := false
		for _, newRange := range newRanges {
			if newRange.start == oldRange.start && newRange.end == oldRange.end {
				found = true
				continue
			}
			if oldRange.IsIncluded(newRange.start) || oldRange.IsIncluded(newRange.end) || newRange.IsIncluded(oldRange.start) || newRange.IsIncluded(oldRange.end) || oldRange.end+1 == newRange.start || oldRange.start == newRange.end+1 {
				// merge the ranges
				newStart := int(math.Min(float64(oldRange.start), float64(newRange.start)))
				newEnd := int(math.Max(float64(oldRange.end), float64(newRange.end)))

				newRange.start = newStart
				newRange.end = newEnd

				found = true
			}
		}
		if !found {
			newRanges = append(newRanges, &Range{oldRange.start, oldRange.end})
		}
	}

	return Ranges{convertFromPointer(newRanges)}
}

func convertFromPointer(ptrs []*Range) []Range {
	ranges := []Range{}
	for _, r := range ptrs {
		ranges = append(ranges, *r)
	}

	return ranges
}

func CountItemRanges(r Ranges) int {
	sum := 0
	for _, spread := range r.r {
		sum += 1 + spread.end - spread.start
	}

	return sum
}
