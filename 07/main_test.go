package main

import (
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	grid := ParseGrid(content)
	assert.Greater(t, len(grid), 1)
	splits := CountBeams(grid)
	assert.Equal(t, 21, splits)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	grid := ParseGrid(content)
	assert.Greater(t, len(grid), 1)
	splits := CountBeams(grid)
	assert.Equal(t, 1687, splits)
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	grid := ParseGrid(content)
	assert.Greater(t, len(grid), 1)
	splits := CountPaths(grid)
	assert.Equal(t, 40, splits)
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	grid := ParseGrid(content)
	assert.Greater(t, len(grid), 1)
	splits := CountPaths(grid)
	assert.Equal(t, 390684413472684, splits)
}
