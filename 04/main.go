package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Hello World!")
}

type Coord struct {
	Row int
	Col int
}

type Grid struct {
	Rolls [][]int
}

func (g Grid) Columns() int {
	if g.Rows() == 0 {
		return 0
	}

	return len(g.Rolls[0])
}

func (g Grid) Rows() int {
	return len(g.Rolls)
}

func (g *Grid) AddRow(r []int) {
	g.Rolls = append(g.Rolls, r)
}

func (g *Grid) RemoveCoord(c Coord) bool {
	if g.Rolls[c.Row][c.Col] == 1 {
		g.Rolls[c.Row][c.Col] = 0
		return true
	}
	return false
}

func (g Grid) String() string {
	content := ""
	for _, row := range g.Rolls {
		for _, roll := range row {
			if roll == 0 {
				content += "."
			} else {
				content += "@"
			}
		}
		content += "\n"
	}
	return content
}

func ReadGrid(content string) Grid {
	grid := Grid{[][]int{}}

	splits := strings.Split(content, "\n")
	for _, split := range splits {
		if split == "" {
			continue
		}
		row := []int{}
		for _, char := range split {
			switch char {
			case '.':
				row = append(row, 0)
			case '@':
				row = append(row, 1)
			default:
				fmt.Fprintf(os.Stderr, "incorrect parsed character %c", char)
			}
		}
		grid.AddRow(row)
	}

	return grid
}

func GetFreeRolls(g Grid) []Coord {
	coords := []Coord{}
	for r, row := range g.Rolls {
		for c, value := range row {
			if value == 0 {
				continue
			}
			countNearby := 0
			if r > 0 {
				// row above
				if c > 0 {
					// top left
					countNearby += g.Rolls[r-1][c-1]
				}
				// above
				countNearby += g.Rolls[r-1][c]
				if c < g.Columns()-1 {
					// top right
					countNearby += g.Rolls[r-1][c+1]
				}
			}
			// left
			if c > 0 {
				// left
				countNearby += g.Rolls[r][c-1]
			}

			// right
			if c < g.Columns()-1 {
				countNearby += g.Rolls[r][c+1]
			}

			if r < g.Rows()-1 {
				// row below
				if c > 0 {
					// bottom left
					countNearby += g.Rolls[r+1][c-1]
				}
				// below
				countNearby += g.Rolls[r+1][c]
				if c < g.Columns()-1 {
					// bottom right
					countNearby += g.Rolls[r+1][c+1]
				}
			}
			if countNearby < 4 {
				coords = append(coords, Coord{r, c})
			}
		}
	}

	return coords
}

func CountRemoveRolls(g *Grid) int {
	countRemovedTotal := 0
	countRemovedRun := 1

	for countRemovedRun > 0 {
		removeCoords := GetFreeRolls(*g)
		for _, coord := range removeCoords {
			g.RemoveCoord(coord)
		}

		countRemovedRun = len(removeCoords)
		countRemovedTotal += len(removeCoords)
	}

	return countRemovedTotal
}
