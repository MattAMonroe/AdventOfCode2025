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
	assert.Greater(t, len(points), 0)

	q := Queue{[]Connection{}, 10}
	q.ComputeConnections(points)

	conns := q.GetConnections()
	assert.Equal(t, 10, len(conns))

	circuits := CreateCircuits(conns)
	product := CountCircuitNodes(circuits, 3)
	assert.Equal(t, 40, product)
}

func TestDistance(t *testing.T) {
	p1 := &Point{162, 817, 812, nil}
	p2 := &Point{425, 690, 689, nil}
	p3 := &Point{431, 825, 988, nil}

	// closest p1 - p2
	// next p1 - p3
	// final p2 - p3

	p1p2 := NewConnection(p1, p2)
	p1p3 := NewConnection(p1, p3)
	p2p3 := NewConnection(p2, p3)

	assert.True(t, p1p2.dist < p1p3.dist)
	assert.True(t, p1p3.dist < p2p3.dist)

	q := Queue{[]Connection{}, 10}
	q.Add(p2p3)
	q.Add(p1p2)
	q.Add(p1p3)

	q = Queue{[]Connection{}, 2}
	q.Add(p2p3)
	q.Add(p1p2)
	q.Add(p1p3)
}

func TestFullP1(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	points := ParseInput(content)
	assert.Greater(t, len(points), 0)

	q := Queue{[]Connection{}, 1000}
	q.ComputeConnections(points)

	conns := q.GetConnections()
	assert.Equal(t, 1000, len(conns))

	circuits := CreateCircuits(conns)
	product := CountCircuitNodes(circuits, 3)

	assert.Equal(t, 90036, product)
}

func TestSampleP2(t *testing.T) {
	content := aoclib.ReadFile("sample.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for sample.txt")

	points := ParseInput(content)
	assert.Greater(t, len(points), 0)

	dq := CreateAllConnections(points)
	assert.Equal(t, 190, dq.Len())

	circuit, conn := CreateSingleCircuit(dq, 20)
	assert.NotEqual(t, nil, circuit)
	assert.NotEqual(t, nil, conn)

	product := conn.p1.X * conn.p2.X
	assert.Equal(t, 25272, product)
}

func TestFullP2(t *testing.T) {
	content := aoclib.ReadFile("problem.txt")
	assert.NotEqualf(t, "", content, "Failed to read in file contents for problem.txt")

	points := ParseInput(content)
	assert.Greater(t, len(points), 0)

	dq := CreateAllConnections(points)
	assert.Equal(t, 499500, dq.Len())

	circuit, conn := CreateSingleCircuit(dq, 1000)
	assert.NotEqual(t, nil, circuit)
	assert.NotEqual(t, nil, conn)

	product := conn.p1.X * conn.p2.X
	assert.Equal(t, 6083499488, product)
}
