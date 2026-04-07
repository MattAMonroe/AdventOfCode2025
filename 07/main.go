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

func GetBeamStart(row []string) int {
	for i, char := range row {
		if char == "S" {
			return i
		}
	}

	return -1
}
