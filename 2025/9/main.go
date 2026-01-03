package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Rectangle struct {
	p1, p2 Point
}

func (r *Rectangle) area() int {
	width := int(math.Abs(float64(r.p1.x-r.p2.x))) + 1
	height := int(math.Abs(float64(r.p1.y-r.p2.y))) + 1
	return width * height
}

func main() {
	points := loadRedTilePoints()
	fmt.Printf("Number of points: %d\n", len(points))

	start := time.Now()
	result1 := naiveApproach(points)
	elapsed1 := time.Since(start)
	fmt.Printf("Naive O(n²) approach: %d (took %v)\n", result1, elapsed1)

	start = time.Now()
	result2 := convexHullApproach(points)
	elapsed2 := time.Since(start)
	fmt.Printf("Convex hull approach: %d (took %v)\n", result2, elapsed2)
}

func naiveApproach(points []Point) int {
	maxArea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			rect := Rectangle{points[i], points[j]}
			if area := rect.area(); area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func convexHullApproach(points []Point) int {
	hull := convexHull(points)
	maxArea := 0
	for i := range hull {
		for j := i + 1; j < len(hull); j++ {
			rect := Rectangle{hull[i], hull[j]}
			if area := rect.area(); area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func convexHull(points []Point) []Point {
	if len(points) <= 1 {
		return points
	}

	sorted := make([]Point, len(points))
	copy(sorted, points)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].x != sorted[j].x {
			return sorted[i].x < sorted[j].x
		}
		return sorted[i].y < sorted[j].y
	})

	var lower []Point
	for _, p := range sorted {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	var upper []Point
	for i := len(sorted) - 1; i >= 0; i-- {
		p := sorted[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	lower = lower[:len(lower)-1]
	upper = upper[:len(upper)-1]
	return append(lower, upper...)
}

func cross(a, b, c Point) int {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func loadRedTilePoints() []Point {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var result []Point
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		result = append(result, Point{
			atoi(parts[0]),
			atoi(parts[1]),
		})
	}
	return result
}

func atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}