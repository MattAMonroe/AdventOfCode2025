package main

import (
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	problems := ParseGrid(content)
	assert.Equal(t, 4, len(problems))

	sum := ComputeMath(problems)
	assert.Equal(t, 4277556, sum)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	problems := ParseGrid(content)

	sum := ComputeMath(problems)
	assert.Equal(t, 5171061464548, sum)
}

func TestRedColumns(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	grid := parseGrid(content)

	assert.Equal(t, 1, readIntFromColumn(grid, 0))
	assert.Equal(t, 24, readIntFromColumn(grid, 1))
	assert.Equal(t, 356, readIntFromColumn(grid, 2))
	assert.Equal(t, 369, readIntFromColumn(grid, 4))

	assert.Equal(t, 248, readIntFromColumn(grid, 5))
	assert.Equal(t, 8, readIntFromColumn(grid, 6))
	assert.Equal(t, 32, readIntFromColumn(grid, 8))
	assert.Equal(t, 581, readIntFromColumn(grid, 9))

	nextIdx := findNextOpOffset(grid[len(grid)-1][1:])
	assert.Equal(t, 3, nextIdx)

	nextIdx = findNextOpOffset(grid[len(grid)-1][5:])
	assert.Equal(t, 3, nextIdx)

	nextIdx = findNextOpOffset(grid[len(grid)-1][9:])
	assert.Equal(t, 3, nextIdx)

	nextIdx = findNextOpOffset(grid[len(grid)-1][13:])
	assert.Equal(t, -1, nextIdx)
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	problems := ParseVertical(content)
	sum := ComputeMath(problems)
	assert.Equal(t, 3263827, sum)
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	problems := ParseVertical(content)
	sum := ComputeMath(problems)
	assert.Equal(t, 10189959087258, sum)
}
