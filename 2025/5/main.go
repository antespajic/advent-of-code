package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func mmax(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	scanningDone := false
	validIds := map[uint64]uint64{}
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanningDone = true
			fmt.Println(" Found empty lien ")
			continue
		}
		if !scanningDone {
			// fmt.Println(line)
			interval := strings.Split(line, "-")
			start, err := strconv.ParseUint(interval[0], 10, 64)
			checkErr(err)
			end, err := strconv.ParseUint(interval[1], 10, 64)
			checkErr(err)
			validIds[start] = mmax(validIds[start], end)
		} else {
			// fmt.Println(line)
			ingredient, err := strconv.ParseUint(line, 10, 64)
			checkErr(err)
			for k, v := range validIds {
				if k <= ingredient && ingredient <= v {
					result++
					break
				}
			}

		}
	}
	fmt.Println(result)
}
