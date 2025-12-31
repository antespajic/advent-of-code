package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	result := 0

	for element := range strings.SplitSeq(input, ",") {
		rng := strings.Split(element, "-")
		start, err := strconv.Atoi(rng[0])
		checkErr(err)
		end, err := strconv.Atoi(strings.TrimSpace(rng[1]))
		checkErr(err)
		for i := start; i <= end; i++ {
			// pt1
			// curr := strconv.Itoa(i)
			// strLen := len(curr)
			// halfLen := strLen / 2
			// if strLen%2 == 0 && curr[:halfLen] == curr[halfLen:] {
			// 	fmt.Println("Adding ", curr)
			// 	result += i
			// }

			// pt2
			curr := strconv.Itoa(i)
			// take chunk by chunk and split, if result of splitting is empty strings
			// then patterns repeat
			for j := 1; j <= len(curr)/2; j++ {
				chunks := strings.Split(curr, curr[:j])
				if len(chunks) > 1 && checkAllElemsEmpty(chunks) {
					fmt.Println(i)
					result += i
					break
				}
			}
		}
	}

	fmt.Println(result)
}

func checkAllElemsEmpty(arr []string) bool {
	for _, elem := range arr {
		if elem != "" {
			return false
		}
	}
	return true
}
