package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello World!")
}

// DA PLAN
// X - Parse our input, get the list of points
// Compute distance between all points, keep shortest N (10 for sample, 1000 for actual)
// from shortes to longest, make circuits, skip if both points are in circuit
// multiply size of 3 largest circuits

type Point struct {
	X       int
	Y       int
	Z       int
	circuit *Circuit
}

func (p *Point) IsEqual(other *Point) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func (p *Point) DistTo(o *Point) float64 {
	xDiff := p.X - o.X
	yDiff := p.Y - o.Y
	zDiff := p.Z - o.Z

	return math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff))
}

func ParsePoint(str string) *Point {
	splits := strings.Split(str, ",")
	if len(splits) != 3 {
		return &Point{0, 0, 0, nil}
	}

	x, err := strconv.Atoi(splits[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Parse X from %s\n", str)
		return &Point{0, 0, 0, nil}
	}

	y, err := strconv.Atoi(splits[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Parse Y from %s\n", str)
		return &Point{0, 0, 0, nil}
	}

	z, err := strconv.Atoi(splits[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Parse Z from %s\n", str)
		return &Point{0, 0, 0, nil}
	}

	return &Point{x, y, z, nil}
}

type Circuit struct {
	nodes map[string]*Point
}

func (c *Circuit) Add(p *Point) {
	p.circuit = c
	c.nodes[p.String()] = p
}

func (c Circuit) HasPoint(p *Point) bool {
	_, present := c.nodes[p.String()]
	return present
}

func (c *Circuit) Merge(o *Circuit) {
	if o == nil {
		return
	}
	for _, node := range o.nodes {
		c.Add(node)
	}

	clear(o.nodes)
}

func (c *Circuit) String() string {
	builder := ""
	for key := range c.nodes {
		if len(builder) > 0 {
			builder += "-"
		}
		builder += key
	}

	return "[" + builder + "]"
}

type Connection struct {
	p1   *Point
	p2   *Point
	dist float64
}

func NewConnection(p1 *Point, p2 *Point) Connection {
	return Connection{p1, p2, p1.DistTo(p2)}
}

type Queue struct {
	// conns should be sorted shortest to longest
	conns  []Connection
	maxLen int
}

func (q *Queue) Add(c Connection) {
	inserted := false
	if len(q.conns)+1 <= q.maxLen {
		for i, o := range q.conns {
			if c.dist < o.dist {
				inserted = true
				q.conns = slices.Insert(q.conns, i, c)
				break
			}
		}
		if !inserted {
			q.conns = append(q.conns, c)
		}
		return
	}

	if c.dist >= q.conns[len(q.conns)-1].dist {
		return
	}

	q.conns = q.conns[:len(q.conns)-1]
	for i, o := range q.conns {
		if c.dist < o.dist {
			inserted = true
			q.conns = slices.Insert(q.conns, i, c)
			break
		}
	}
	if !inserted {
		q.conns = append(q.conns, c)
	}
}

func (q *Queue) GetConnections() []Connection {
	return q.conns
}

func (q *Queue) ComputeConnections(points []*Point) {
	for i, p1 := range points {
		if i == len(points) {
			continue
		}
		for _, p2 := range points[i+1:] {
			dist := p1.DistTo(p2)
			conn := Connection{p1, p2, dist}
			q.Add(conn)
		}
	}
}

func CreateCircuits(conns []Connection) []*Circuit {
	circuits := []*Circuit{}
	for _, conn := range conns {
		if conn.p1.circuit == nil && conn.p2.circuit == nil {
			// new points not in circuits, need to create a new circuit
			newcircuit := &Circuit{nodes: map[string]*Point{}}
			newcircuit.Add(conn.p1)
			newcircuit.Add(conn.p2)
			circuits = append(circuits, newcircuit)
			continue
		}
		if conn.p1.circuit == conn.p2.circuit {
			continue
		}
		if conn.p1.circuit != nil {
			conn.p1.circuit.Merge(conn.p2.circuit)
			conn.p1.circuit.Add(conn.p2)
		} else {
			conn.p2.circuit.Add(conn.p1)
		}
	}

	slices.SortFunc(circuits, func(a, b *Circuit) int {
		return cmp.Compare(len(b.nodes), len(a.nodes))
	})

	return circuits
}

func CountCircuitNodes(circuits []*Circuit, num int) int {
	count := int(math.Min(float64(num), float64(len(circuits))))
	product := 1

	for i := range count {
		product *= len(circuits[i].nodes)
	}

	return product
}

func ParseInput(content string) []*Point {
	points := []*Point{}

	splits := strings.Split(strings.Trim(content, " "), "\n")
	for _, split := range splits {
		strPoint := strings.Trim(split, " ")
		if strPoint == "" {
			continue
		}
		point := ParsePoint(strPoint)
		if point.X == 0 || point.Y == 0 || point.Z == 0 {
			fmt.Fprintf(os.Stderr, "Invalid point received for %s: %v\n", strPoint, point)
		}

		points = append(points, point)
	}

	return points
}
