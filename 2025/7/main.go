package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkError(error error) {
	if error != nil {
		panic(error)
	}
}

type Point struct {
	r, c int
}

func pt1(grid [][]rune) {

	rows := len(grid)
	cols := len(grid[0])

	currentPositions := make(map[int]bool)
	for i, el := range grid[0] {
		if el == 'S' {
			currentPositions[i] = true
			break
		}
	}

	triggeredSplits := make(map[Point]bool)
	for i := range rows {
		nextPositions := make(map[int]bool)

		for j := range currentPositions {
			if grid[i][j] == '^' {
				triggeredSplits[Point{i, j}] = true

				if i+1 < rows {
					if j-1 >= 0 {
						nextPositions[j-1] = true
					}
					if j+1 < cols {
						nextPositions[j+1] = true
					}
				}
			} else {
				// no split: continue downwards
				if i+1 < rows {
					nextPositions[j] = true
				}
			}
		}
		currentPositions = nextPositions
	}

	fmt.Println(len(triggeredSplits))
}

type Cell struct {
	r, c int
}

func pt2(grid [][]rune) {
	startCol := -1
	for i, el := range grid[0] {
		if el == 'S' {
			startCol = i
			break
		}
	}

	memo := make(map[Cell]int)

	result := countUniquePaths(1, startCol, grid, memo)
	fmt.Println(result)
}

func countUniquePaths(r, c int, grid [][]rune, memo map[Cell]int) int {
	if r >= len(grid) {
		return 1
	}

	if c < 0 || c >= len(grid[0]) {
		return 0
	}

	if val, exists := memo[Cell{r, c}]; exists {
		return val
	}

	result := 0
	if grid[r][c] == '^' {
		left := countUniquePaths(r+1, c-1, grid, memo)
		right := countUniquePaths(r+1, c+1, grid, memo)
		result = left + right
	} else {
		result = countUniquePaths(r+1, c, grid, memo)
	}

	memo[Cell{r, c}] = result

	return result
}

func main() {
	grid, error := loadGrid()
	checkError(error)
	pt2(grid)
}

func loadGrid() ([][]rune, error) {
	file, err := os.Open("input.txt")
	checkError(err)

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid, scanner.Err()
}
