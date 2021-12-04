package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var flagTest = flag.Bool("test", false, "use the test input")

func getDepthReadings() []int {
	data := `199
200
208
210
200
207
240
269
260
263`

	if !*flagTest {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	lines := strings.Split(data, "\n")
	depths := make([]int, len(lines))
	for i, e := range lines {
		d, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		depths[i] = d
	}

	return depths
}

func countDepthIncreases(depths []int) (int, int) {
	perReadingCount := 0
	slidingWindowCount := 0
	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			perReadingCount++
		}
		if i > 2 && depths[i-3] < depths[i] {
			slidingWindowCount++
		}
	}
	return perReadingCount, slidingWindowCount
}

func main() {
	flag.Parse()

	depths := getDepthReadings()
	start := time.Now()
	p1, p2 := countDepthIncreases(depths)
	done := time.Since(start)

	fmt.Printf("Part 1: %d\nPart 2: %d\nTime: %s\n", p1, p2, done)
}
