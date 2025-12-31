package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// DSU (Disjoint Set Union) handles the "Circuit" logic efficiently
type DSU struct {
	parent []int
	size   []int
}

// NewDSU initializes every point in its own circuit of size 1
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent, size}
}

// Find locates the "root" ID of the circuit a point belongs to
// Uses path compression for O(1) average speed
func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i])
	return d.parent[i]
}

// Union merges the circuits containing points i and j
func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	// if Find(A) == Find(B) then A and B are in same circuit
	if rootI != rootJ {
		// Merge smaller set into larger set
		if d.size[rootI] < d.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		d.parent[rootJ] = rootI
		d.size[rootI] += d.size[rootJ]
		return true
	}
	return false
}

type Point struct {
	x, y, z int
	id      int // Unique ID 0..N-1
}

type Edge struct {
	u, v   int // IDs of points
	distSq int // Squared distance
}

type Pair struct {
	a, b int
}

func main() {
	points := loadPoints()

	// Generate all possible edges
	edges := make([]Edge, 0, len(points)*(len(points)-1)/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			d := (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)
			edges = append(edges, Edge{u: i, v: j, distSq: d})
		}
	}

	// Sort Edges by Distance (Shortest first)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distSq < edges[j].distSq
	})

	// Process the top 1000 edges
	dsu := NewDSU(len(points))
	// for pt2 we process all
	limit := max(len(edges), 1000)

	// DSU Union automatically handles creating, adding, and merging.
	joins := []Pair{}
	for _, edge := range edges[:limit] {
		if dsu.Union(edge.u, edge.v) {
			// for PT2, keep track of joins so we can multiply the last one
			joins = append(joins, Pair{points[edge.u].x, points[edge.v].x})
		}
	}

	// Collect sizes of all unique circuits
	// We use a map to ensure we only count each circuit root once
	circuitSizes := make(map[int]int)
	for i := range points {
		root := dsu.Find(i)
		// We use the root to identify the unique circuit
		circuitSizes[root] = dsu.size[root]
	}

	// Extract sizes to a slice and sort descending
	sizes := []int{}
	for _, s := range circuitSizes {
		sizes = append(sizes, s)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	// Multiply top 3
	result := 1
	for i, s := range sizes {
		if i >= 3 {
			break
		}
		result *= s
	}

	// fmt.Println("Result:", result)
	lastPair := joins[len(joins)-1]
	fmt.Println("Pt2 Result: ", lastPair.a*lastPair.b)
}

func loadPoints() []Point {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var points []Point
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		points = append(points, Point{
			x:  atoi(parts[0]),
			y:  atoi(parts[1]),
			z:  atoi(parts[2]),
			id: id,
		})
		id++
	}
	return points
}

func atoi(s string) int {
	val, _ := strconv.Atoi(strings.TrimSpace(s))
	return val
}
