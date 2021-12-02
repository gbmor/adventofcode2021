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

func calc(depths []int) (int, int) {
	count := 0
	count2 := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			count++
		}
		if i > 2 && depths[i-3] < depths[i] {
			count2++
		}
	}
	return count, count2
}

func main() {
	depths, err := getDepths()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	start := time.Now()
	p1, p2 := calc(depths)
	done := time.Since(start)
	fmt.Printf("Part 1: %d\nPart 2: %d\nTime: %s\n", p1, p2, done)
}
