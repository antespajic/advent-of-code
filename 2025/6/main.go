package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(scanner bufio.Scanner) {
	re := regexp.MustCompile(`\s+`)
	equations := map[int][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		nums := re.Split(line, -1)
		for i, n := range nums {
			equations[i] = append(equations[i], n)
		}
	}

	result := 0
	for _, v := range equations {
		operand := v[len(v)-1]
		nums := MapToInt(v[:len(v)-1])
		if operand == "*" {
			result += MulAll(nums)
		}
		if operand == "+" {
			result += AddAll(nums)
		}
	}
	fmt.Println(result)
}

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)
	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	nums := []int{}
	result := 0
	for j := len(grid[0]) - 1; j >= 0; j-- {
		colJ := []rune{}
		operator := ' '
		for i := range len(grid) {
			if grid[i][j] == '*' || grid[i][j] == '+' {
				operator = grid[i][j]
				break
			}
			if grid[i][j] != ' ' {
				colJ = append(colJ, grid[i][j])
			}
		}
		if len(colJ) == 0 {
			continue
		}
		num, err := strconv.Atoi(string(colJ))
		checkErr(err)
		nums = append(nums, num)
		if operator != ' ' {
			if operator == '+' {
				result += AddAll(nums)
			}
			if operator == '*' {
				result += MulAll(nums)
			}
			nums = []int{}
		}
	}
	// fmt.Println(len(grid))
	fmt.Println(result)
}

func MulAll(slice []int) int {
	res := 1
	for _, x := range slice {
		res *= x
	}
	return res
}

func AddAll(slice []int) int {
	res := 0
	for _, x := range slice {
		res += x
	}
	return res
}

func MapToInt(slice []string) []int {
	res := []int{}
	for _, x := range slice {
		n, err := strconv.Atoi(x)
		checkErr(err)
		res = append(res, n)
	}
	return res
}
