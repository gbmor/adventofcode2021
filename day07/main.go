package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var flagTest = flag.Bool("test", false, "use the test input")

func getCrabPositions() []int {
	data := `16,1,2,0,4,2,7,1,2,14`

	if !*flagTest {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	split := strings.Split(data, ",")
	crabs := make([]int, len(split))
	var err error
	for i, e := range split {
		crabs[i], err = strconv.Atoi(e)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	return crabs
}

func doCrabThings(crabs []int) (int, int) {
	high := math.MinInt32
	low := math.MaxInt32
	for _, c := range crabs {
		if c > high {
			high = c
		} else if c < low {
			low = c
		}
	}

	minimumFuelConsumed := math.MaxInt32
	minimumFuelConsumedRolling := math.MaxInt32
	for i := low; i < high; i++ {
		fuelConsumed := 0
		fuelConsumedRolling := 0
		for _, c := range crabs {
			fuelConsumedThisCrab := c - i
			if fuelConsumedThisCrab < 0 {
				fuelConsumedThisCrab *= -1
			}
			fuelConsumed += fuelConsumedThisCrab
			for k := 0; k <= fuelConsumedThisCrab; k++ {
				fuelConsumedRolling += k
			}
		}
		if fuelConsumed < minimumFuelConsumed {
			minimumFuelConsumed = fuelConsumed
		}
		if fuelConsumedRolling < minimumFuelConsumedRolling {
			minimumFuelConsumedRolling = fuelConsumedRolling
		}
	}

	return minimumFuelConsumed, minimumFuelConsumedRolling
}

func main() {
	flag.Parse()

	data := getCrabPositions()
	start := time.Now()
	p1, p2 := doCrabThings(data)
	elapsedTime := time.Since(start)

	fmt.Printf("Crab Submarine Fuel Consumption\nPart 1:\t%d\nPart 2:\t%d\nTime:\t%s\n", p1, p2, elapsedTime)
}
