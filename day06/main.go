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

func getLanternfish() []int {
	data := `3,4,3,1,2`

	if !*flagTest {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	split := strings.Split(data, ",")
	fish := make([]int, len(split))
	for i, e := range split {
		fish[i], _ = strconv.Atoi(e)
	}

	return fish
}

func doFishThings(lanternfish []int, days int) int64 {
	// keep track of how many lanternfish are spawning each day,
	// instead of keeping track of when a given lanternfish spawns.
	fishReadyEachDay := make([]int64, days)
	for i := range lanternfish {
		fishReadyEachDay[lanternfish[i]]++
	}

	fishCount := int64(len(lanternfish))
	for i := 0; i < days; i++ {
		nextSpawnForFish := i + 7
		nextSpawnForNewFish := i + 9

		// number of lanternfish spawning today
		spawning := fishReadyEachDay[i]

		// the newly born lanternfish who will spawn in 9 days
		if nextSpawnForNewFish < len(fishReadyEachDay) {
			fishReadyEachDay[nextSpawnForNewFish] += spawning
		}

		// the lanternfish spawning today will spawn again in 7 days
		if nextSpawnForFish < len(fishReadyEachDay) {
			fishReadyEachDay[nextSpawnForFish] += spawning
		}

		// total so far
		fishCount += spawning
	}

	return fishCount
}

func main() {
	flag.Parse()

	data := getLanternfish()
	start := time.Now()
	p1 := doFishThings(data, 80)
	p1Time := time.Since(start)

	start = time.Now()
	p2 := doFishThings(data, 256)
	p2Time := time.Since(start)

	fmt.Printf("Part 1:\t\t%d\t\t%s\nPart 2:\t\t%d\t%s\nTotal Time:\t\t\t%s\n", p1, p1Time, p2, p2Time, p1Time+p2Time)
}
