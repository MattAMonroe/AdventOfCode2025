package main

import (
	"testing"

	"github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	countZeroTurns := CountZeros(content)

	assert.NotEqualf(t, "", content, "Didn't get file content, got \n===\n%s\n===\n", content)
	assert.Equalf(t, 3, countZeroTurns, "Didn't get correct number of Zero Turns, got %d", countZeroTurns)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	countZeroTurns := CountZeros(content)

	assert.NotEqualf(t, "", content, "Didn't get file content, got \n===\n%s\n===\n", content)
	assert.Equalf(t, 1059, countZeroTurns, "Didn't get correct number of Zero Turns, got %d", countZeroTurns)
}

func TestCountZeroRight(t *testing.T) {
	zeroTurns := CountZeros("R50")
	assert.Equalf(t, 1, zeroTurns, "Simple R50 failed")

	zeroTurns = CountZeros("R25\nR24")
	assert.Equalf(t, 0, zeroTurns, "Simple R25,R24 failed")

	zeroTurns = CountZeros("R25\nR24\nR1")
	assert.Equalf(t, 1, zeroTurns, "Simple R25,R24,R1 failed")

	zeroTurns = CountZeros("R25\nR24\nR2")
	assert.Equalf(t, 0, zeroTurns, "Simple R25,R24,R2 failed")
}

func TestCountZeroMultLeft(t *testing.T) {
	zeroTurns := CountZeros("L550")
	assert.Equalf(t, 1, zeroTurns, "Multiple L550 failed")
}

func TestCountZeroLeft(t *testing.T) {
	zeroTurns := CountZeros("L50")
	assert.Equalf(t, 1, zeroTurns, "Simple L50 failed")

	zeroTurns = CountZeros("L25\nL24")
	assert.Equalf(t, 0, zeroTurns, "Simple L25,L24 failed")

	zeroTurns = CountZeros("L25\nL24\nL1")
	assert.Equalf(t, 1, zeroTurns, "Simple L25,L24,L1 failed")

	zeroTurns = CountZeros("L25\nL24\nL2")
	assert.Equalf(t, 0, zeroTurns, "Simple L25,L24,L2 failed")
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	countZeroTurns := CountZeroPasses(content)

	assert.NotEqualf(t, "", content, "Didn't get file content, got \n===\n%s\n===\n", content)
	assert.Equalf(t, 6, countZeroTurns, "Didn't get correct number of Zero Turns, got %d", countZeroTurns)
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	countZeroTurns := CountZeroPasses(content)

	assert.NotEqualf(t, "", content, "Didn't get file content, got \n===\n%s\n===\n", content)
	assert.Equalf(t, 1059, countZeroTurns, "Didn't get correct number of Zero Turns, got %d", countZeroTurns)
}

func TestCountZeroPassesRight(t *testing.T) {
	zeroTurns := CountZeroPasses("R500")
	assert.Equalf(t, 5, zeroTurns, "Simple R500 failed")

	zeroTurns = CountZeroPasses("R25\nR24\nR2")
	assert.Equalf(t, 1, zeroTurns, "Simple R25,R24,R2 failed")

	zeroTurns = CountZeroPasses("R50\nR200")
	assert.Equalf(t, 3, zeroTurns, "Simple R50,R200 failed")
}

func TestCountZeroPassesLeft(t *testing.T) {
	zeroTurns := CountZeroPasses("L500")
	assert.Equalf(t, 5, zeroTurns, "Simple L500 failed")

	zeroTurns = CountZeroPasses("L25\nL24\nL2")
	assert.Equalf(t, 1, zeroTurns, "Simple L25,L24,L2 failed")

	zeroTurns = CountZeroPasses("L50\nL200")
	assert.Equalf(t, 3, zeroTurns, "Simple L50,L200 failed")
}

func TestProcessInput(t *testing.T) {
	content := "L28\nR16"
	commands := ProcessInput(content)
	assert.Equalf(t, 2, len(commands), "Didn't get correct number of commands from ProcessInput")
	assert.Equalf(t, DirectionFromInt(DirLeft), commands[0].direction, "Incorrect direction for first command")
	assert.Equalf(t, 28, commands[0].clicks, "Incorrect number of clicks for first command")
	assert.Equalf(t, DirectionFromInt(DirRight), commands[1].direction, "Incorrect direction for second command")
	assert.Equalf(t, 16, commands[1].clicks, "Incorrect number of clicks for second command")
}
