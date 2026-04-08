package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Hello World!")
}

func ParseGrid(content string) [][]string {
	lines := strings.Split(strings.Trim(content, " "), "\n")

	grid := [][]string{}
	for _, line := range lines {
		row := []string{}
		if strings.Trim(line, " ") == "" {
			continue
		}

		for _, char := range line {
			row = append(row, string(char))
		}

		grid = append(grid, row)
	}

	return grid
}

func CountBeams(grid [][]string) int {
	countSplits := 0

	beamStart := GetBeamStart(grid[0])
	if beamStart == -1 {
		fmt.Fprintf(os.Stderr, "Failed to find Beam Start\n")
		return -1
	}
	beams := map[int]bool{}
	beams[beamStart] = true
	for _, row := range grid[1:] {
		newBeams := map[int]bool{}
		for beam := range beams {
			if row[beam] == "^" {
				if beam-1 >= 0 {
					newBeams[beam-1] = true
				}
				if beam+1 < len(row) {
					newBeams[beam+1] = true
				}
				countSplits += 1
			} else {
				newBeams[beam] = true
			}
		}

		beams = newBeams
	}

	return countSplits
}

func CountPaths(grid [][]string) int {
	beamStart := GetBeamStart(grid[0])
	if beamStart == -1 {
		fmt.Fprintf(os.Stderr, "Failed to find Beam Start\n")
		return -1
	}

	currentRow := []int{}
	previousRow := []int{}
	for range len(grid[0]) {
		currentRow = append(currentRow, 0)
		previousRow = append(previousRow, 0)
	}

	previousRow[beamStart] = 1

	for _, row := range grid[1:] {
		for i, char := range row {
			if char == "^" {
				currentRow[i] = 0
				continue
			}
			value := previousRow[i]
			if i-1 >= 0 && row[i-1] == "^" {
				value += previousRow[i-1]
			}

			if i+1 < len(row) && row[i+1] == "^" {
				value += previousRow[i+1]
			}

			currentRow[i] = value
		}
		temp := previousRow
		previousRow = currentRow
		currentRow = temp
	}

	sum := 0
	for _, val := range previousRow {
		sum += val
	}

	return sum
}

func GetBeamStart(row []string) int {
	for i, char := range row {
		if char == "S" {
			return i
		}
	}

	return -1
}
