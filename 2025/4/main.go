package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

type Removed struct {
	Row int
	Col int
}

//go:embed input.txt
var s string

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkTopLeft(rows []string, i, j int, removed []Removed) bool {
	if i == 0 || j == 0 {
		return false
	}
	row := i - 1
	col := j - 1
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkUp(rows []string, i, j int, removed []Removed) bool {
	if i == 0 {
		return false
	}
	row := i - 1
	col := j
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkTopRight(rows []string, i, j int, removed []Removed) bool {
	col := j + 1
	if j == len(rows[0])-1 || i == 0 {
		return false
	}
	row := i - 1
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkLeft(rows []string, i, j int, removed []Removed) bool {
	row := i
	col := j - 1
	if j == 0 {
		return false
	}
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkRight(rows []string, i, j int, removed []Removed) bool {
	row := i
	col := j + 1
	if j == len(rows)-1 {
		return false
	}
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkBottomLeft(rows []string, i, j int, removed []Removed) bool {
	row := i + 1
	if i == len(rows)-1 || j == 0 {
		return false
	}
	col := j - 1
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkBelow(rows []string, i, j int, removed []Removed) bool {
	row := i + 1
	if i == len(rows)-1 {
		return false
	}
	col := j
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

func checkBottomRight(rows []string, i, j int, removed []Removed) bool {
	row := i + 1
	if i == len(rows)-1 || j == len(rows[0])-1 {
		return false
	}
	col := j + 1
	return rows[row][col] == '@' && !slices.Contains(removed, Removed{row, col})
}

// Note for Pt1: doesn't have removed array
func main() {
	result := 0
	s = strings.TrimSpace(s)
	rows := strings.Split(s, "\n")
	removed := []Removed{}
	for {
		tmpRemoved := []Removed{}
		for i := range rows {
			rows[i] = strings.TrimSpace(rows[i])
			for j := 0; j < len(rows[0]); j++ {
				if rows[i][j] != '@' || slices.Contains(removed, Removed{i, j}) {
					continue
				}
				tmp := 0
				if checkTopLeft(rows, i, j, removed) {
					tmp += 1
				}
				if checkUp(rows, i, j, removed) {
					tmp += 1
				}
				if checkTopRight(rows, i, j, removed) {
					tmp += 1
				}
				if checkLeft(rows, i, j, removed) {
					tmp += 1
				}
				if checkRight(rows, i, j, removed) {
					tmp += 1
				}
				if checkBottomLeft(rows, i, j, removed) {
					tmp += 1
				}
				if checkBelow(rows, i, j, removed) {
					tmp += 1
				}
				if checkBottomRight(rows, i, j, removed) {
					tmp += 1
				}
				if tmp < 4 {
					result++
					tmpRemoved = append(tmpRemoved, Removed{
						i,
						j,
					})
				}
			}
		}
		// fmt.Println(len(tmpRemoved))
		if len(tmpRemoved) == 0 {
			break
		}
		removed = append(removed, tmpRemoved...)
	}

	fmt.Println(len(removed))
}
