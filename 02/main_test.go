package main

import (
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSimpleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqual(t, "", content, "Didn't read in sample file")

	ranges := ParseRanges(content)
	assert.Greater(t, len(ranges), 0, "Wasn't able to parse out ranges")

	sum := FindInvalidIDs(ranges)
	assert.Equal(t, 1227775554, sum)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqual(t, "", content, "Didn't read in problem file")

	ranges := ParseRanges(content)
	assert.Greater(t, len(ranges), 0, "Wasn't able to parse out ranges")

	sum := FindInvalidIDs(ranges)
	assert.Equal(t, 20223751480, sum)
}

func TestCheckInvalid(t *testing.T) {
	assert.Falsef(t, CheckInvalid(111), "Odd Number of Digits passed")
	assert.Truef(t, CheckInvalid(11), "Simple Case")
	assert.Truef(t, CheckInvalid(123123), "More difficult case")
	assert.Falsef(t, CheckInvalid(123321), "pallindrome passed")
}

func TestFindInvalidIDs(t *testing.T) {
	sum := FindInvalidIDs([]Range{Range{11, 22}})
	assert.Equalf(t, 33, sum, "Didn't get correct sum for 11-22 range")
}

func TestSimpleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqual(t, "", content, "Didn't read in sample file")

	ranges := ParseRanges(content)
	assert.Greater(t, len(ranges), 0, "Wasn't able to parse out ranges")

	sum := FindInvalidIDsExtra(ranges)
	assert.Equal(t, 4174379265, sum)
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqual(t, "", content, "Didn't read in sample file")

	ranges := ParseRanges(content)
	assert.Greater(t, len(ranges), 0, "Wasn't able to parse out ranges")

	sum := FindInvalidIDsExtra(ranges)
	assert.Equal(t, 30260171216, sum)
}

func TestCheckExtra(t *testing.T) {
	assert.Truef(t, CheckExtra(11), "Simple Case")
	assert.Truef(t, CheckExtra(123123), "More difficult case")
	assert.Falsef(t, CheckExtra(112), "almost pallindrome passed")
	assert.Falsef(t, CheckExtra(123321), "pallindrome passed")

	assert.Truef(t, CheckExtra(111), "Odd Number of repeating Digits failed")
	assert.Truef(t, CheckExtra(121212), "step down difficult case")
}
