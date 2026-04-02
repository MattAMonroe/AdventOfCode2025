package main

import (
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	g := ReadGrid(content)

	count := CountFreeRolls(g)
	assert.Equal(t, 13, count)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	g := ReadGrid(content)

	count := CountFreeRolls(g)
	assert.Equal(t, 1419, count)
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")
}
