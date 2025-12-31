package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(line string) int {
	lineLen := len(line)
	runedLine := []rune(line)
	firstPtr := 0
	firstNum := runedLine[0]
	for i, c := range line[:lineLen-1] {
		if c > firstNum {
			firstNum = c
			firstPtr = i
		}
		if firstNum == '9' {
			break
		}
	}
	secondNum := runedLine[firstPtr+1]
	for i := firstPtr + 1; i < lineLen; i++ {
		if runedLine[i] > secondNum {
			secondNum = runedLine[i]
		}
		if secondNum == '9' {
			break
		}
	}
	highestNum, err := strconv.Atoi(string(firstNum) + string(secondNum))
	checkErr(err)
	return highestNum
}

func part2(line string) int {
	start := 0
	result := 0
	for i := range 12 {
		currentMax := 0
		for j := start; j <= len(line)-(12-i); j++ {
			d := int(line[j] - '0')
			if d > currentMax {
				currentMax = d
				start = j + 1
			}
		}
		result = result*10 + currentMax
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		// result += part1(line)
		result += part2(line)
	}

	fmt.Println(result)
}
