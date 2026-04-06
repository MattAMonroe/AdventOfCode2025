package main

import (
	// "fmt"
	// "os"
	"fmt"
	"os"
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	r, v := ParseInput(content)
	count := FindIncludedValues(r, v)
	assert.Equal(t, 3, count)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	r, v := ParseInput(content)
	count := FindIncludedValues(r, v)
	assert.Equal(t, 652, count)
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	ranges, _ := ParseInput(content)
	newRanges := CompactRanges(ranges)
	count := CountItemRanges(newRanges)
	assert.Equal(t, 14, count)
}

func TestCompactRangesSample(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	ranges, _ := ParseInput(content)
	assert.Equal(t, 4, len(ranges.r))
	newRanges := CompactRanges(ranges)
	assert.Equal(t, 2, len(newRanges.r))
}

func TestCompactRangesTouching(t *testing.T) {
	content := "3-4\n5-7"
	ranges, _ := ParseInput(content)
	assert.Equal(t, 2, len(ranges.r))
	newRanges := CompactRanges(ranges)
	assert.Equal(t, 1, len(newRanges.r))
	count := CountItemRanges(newRanges)
	assert.Equal(t, 5, count)
}

func TestRangesOverlap(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	ranges, _ := ParseInput(content)
	newRanges := CompactRanges(ranges)

	for _, origRange := range ranges.r {
		if newRanges.IsIncluded(origRange.start) && newRanges.IsIncluded(origRange.end) {
			continue
		}

		fmt.Fprintf(os.Stderr, "Couldn't find new Range for original Range: %v\n", origRange)
	}

	for _, newRange := range newRanges.r {
		if ranges.IsIncluded(newRange.start) && ranges.IsIncluded(newRange.end) {
			continue
		}

		fmt.Fprintf(os.Stderr, "Couldn't find new Range for New Range: %v\n", newRange)
	}
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	ranges, _ := ParseInput(content)
	newRanges := CompactRanges(ranges)

	count := CountItemRanges(newRanges)
	assert.Equal(t, 341753674214273, count)

}
