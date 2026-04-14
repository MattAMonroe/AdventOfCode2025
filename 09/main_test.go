package main

import (
	"testing"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
	"github.com/stretchr/testify/assert"
)

func TestSampleP1(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	points := ParseInput(content)
	assert.Equal(t, 8, len(points))

	area := FindMaxArea(points)
	assert.Equal(t, 50, area)
}

func TestGetArea(t *testing.T) {
	p1 := Point{2, 5}
	p2 := Point{9, 7}
	p1p2Area := p1.GetArea(p2)
	assert.Equal(t, 24, p1p2Area)

	p2p1Area := p2.GetArea(p1)
	assert.Equal(t, 24, p2p1Area)

	p3 := Point{7, 1}
	p4 := Point{11, 7}
	p3p4Area := p3.GetArea(p4)
	assert.Equal(t, 35, p3p4Area)

	p4p3Area := p4.GetArea(p3)
	assert.Equal(t, 35, p4p3Area)

	p5 := Point{7, 3}
	p6 := Point{2, 3}
	p5p6Area := p5.GetArea(p6)
	assert.Equal(t, 6, p5p6Area)

	p6p5Area := p6.GetArea(p5)
	assert.Equal(t, 6, p6p5Area)

	p7 := Point{2, 5}
	p8 := Point{11, 1}
	p7p8Area := p7.GetArea(p8)
	assert.Equal(t, 50, p7p8Area)

	p8p7Area := p8.GetArea(p7)
	assert.Equal(t, 50, p8p7Area)

}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	points := ParseInput(content)
	assert.Equal(t, 496, len(points))

	area := FindMaxArea(points)
	assert.Equal(t, 4741451444, area)
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")
}
