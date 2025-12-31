package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	currentSum := 50
	answer1 := 0
	answer2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		incrementStep := 1
		if strings.HasPrefix(line, "L") {
			incrementStep = -1
		}

		for range number {
			currentSum += incrementStep
			if currentSum%100 == 0 {
				answer2++
			}
		}

		currentSum = mmod(currentSum, 100)
		if currentSum == 0 {
			answer1++
		}
	}

	// fmt.Println(answer1)
	fmt.Println(answer2)
}

func mmod(x, d int) int {
	return (x%d + d) % d
}
