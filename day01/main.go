package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func getDepths() ([]int, error) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(b, []byte("\n"))
	depths := make([]int, len(lines))
	for i, e := range lines {
		d, err := strconv.Atoi(string(e))
		if err != nil {
			return nil, err
		}
		depths[i] = d
	}
	return depths, nil
}

func part1(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			count++
		}
	}
	return count
}

func part2(depths []int) int {
	count := 0
	for i := 3; i < len(depths); i++ {
		first := depths[i-3] + depths[i-2] + depths[i-1]
		second := depths[i-2] + depths[i-1] + depths[i]
		if second > first {
			count++
		}
	}
	return count
}

func main() {
	depths, err := getDepths()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	start := time.Now()
	p1 := part1(depths)
	p1time := time.Since(start)
	fmt.Printf("Part 1: %d\t(%s)\n", p1, p1time)

	start = time.Now()
	p2 := part2(depths)
	p2time := time.Since(start)
	fmt.Printf("Part 2: %d\t(%s)\n", p2, p2time)
}
