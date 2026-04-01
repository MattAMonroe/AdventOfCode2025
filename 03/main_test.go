package main

import (
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	banks := ReadBatteryBanks(content)

	sum := SumBestCombos(banks)
	assert.Equal(t, 357, sum)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	banks := ReadBatteryBanks(content)

	sum := SumBestCombos(banks)
	assert.Equal(t, 17196, sum)
}

func TestParseBatteries(t *testing.T) {
	banks := ReadBatteryBanks("123\n456")
	assert.Equal(t, 2, len(banks))
	assert.Equal(t, 3, len(banks[0]))
	assert.Equal(t, 3, len(banks[1]))

	assert.Equal(t, 1, banks[0][0])
	assert.Equal(t, 2, banks[0][1])
	assert.Equal(t, 3, banks[0][2])

	assert.Equal(t, 4, banks[1][0])
	assert.Equal(t, 5, banks[1][1])
	assert.Equal(t, 6, banks[1][2])
}

func TestFindBestCombo(t *testing.T) {
	value := FindBestCombo([]int{1, 2, 3})
	assert.Equal(t, 23, value)

	value = FindBestCombo([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, 98, value)
	value = FindBestCombo([]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9})
	assert.Equal(t, 89, value)
	value = FindBestCombo([]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8})
	assert.Equal(t, 78, value)
	value = FindBestCombo([]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1})
	assert.Equal(t, 92, value)

}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")
}
