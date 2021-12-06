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

func doFishThings(lanternfish []int) (int64, int64) {
	maxDays := 256

	// keep track of how many lanternfish are spawning each day,
	// instead of keeping track of when a given lanternfish spawns.
	fishReadyEachDay := make([]int64, maxDays)
	for i := range lanternfish {
		fishReadyEachDay[lanternfish[i]]++
	}

	fishCountDay256 := int64(len(lanternfish))
	fishCountDay80 := int64(len(lanternfish))
	for i := 0; i < maxDays; i++ {
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
		fishCountDay256 += spawning
		if i < 80 {
			fishCountDay80 += spawning
		}
	}

	return fishCountDay80, fishCountDay256
}

func main() {
	flag.Parse()

	data := getLanternfish()
	start := time.Now()
	p1, p2 := doFishThings(data)
	elapsedTime := time.Since(start)

	fmt.Printf("Lanternfish Population\n80 Days:\t%d\n256 Days:\t%d\nTime:\t\t%s\n", p1, p2, elapsedTime)
}
