package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello World!")
}

type Point struct {
	X int
	Y int
}

func (p Point) GetArea(o Point) int {
	x := abs(p.X-o.X) + 1
	y := abs(p.Y-o.Y) + 1
	return x * y
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func ParseInput(content string) []Point {
	points := []Point{}
	splits := strings.Split(strings.Trim(content, " "), "\n")
	for _, split := range splits {
		if strings.Trim(split, " ") == "" {
			continue
		}
		coords := strings.Split(split, ",")
		if len(coords) != 2 {
			fmt.Fprintf(os.Stderr, "Failed to parse coords: %s\n", split)
			continue
		}
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse coords.X: %s\n", coords[0])
			continue
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse coords.Y: %s\n", coords[1])
			continue
		}
		points = append(points, Point{x, y})
	}

	return points
}

func FindMaxArea(points []Point) int {
	largest := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			area := p1.GetArea(p2)
			if area > largest {
				largest = area
			}
		}
	}

	return largest
}
