package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"slices"
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
	type Range struct {
		From uint64
		To   uint64
	}
	scanner := bufio.NewScanner(file)
	ranges := []Range{}
	mergedIntervals := []Range{}
	result := uint64(0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println(" Found empty lien ")
			break
		}
		interval := strings.Split(line, "-")
		start, err := strconv.ParseUint(interval[0], 10, 64)
		checkErr(err)
		end, err := strconv.ParseUint(interval[1], 10, 64)
		checkErr(err)
		ranges = append(ranges, Range{
			From: start,
			To:   end,
		})
		// validIds[start] = mmax(validIds[start], end)
	}
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.From < b.From {
			return -1
		} else if a.From > b.From {
			return 1
		}
		return 0
	})
	// for _, slice := range ranges {
	// 	fmt.Println(slice.From, " ", slice.To)
	// }
	next := 0
	for i, r := range ranges {
		if i < next {
			// fmt.Println("Skipping ", r)
			continue
		}
		maxTo := r.To
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j].From >= r.From && ranges[j].From <= maxTo {
				// fmt.Println("Merging ", ranges[j], " into ", r)
				maxTo = mmax(maxTo, ranges[j].To)
			} else {
				// fmt.Println("Not merging ", ranges[j], " into ", r)
				next = j
				break
			}

		}
		mergedIntervals = append(mergedIntervals, Range{
			From: r.From,
			To:   maxTo,
		})
	}
	// fmt.Println("Initial ", len(ranges))
	// fmt.Println("Merged ", len(mergedIntervals))
	for _, slice := range mergedIntervals {
		fmt.Println(slice.From, " ", slice.To)
	}
	for _, r := range mergedIntervals {
		result += (r.To - r.From + 1)
	}
	fmt.Println(result)
}
