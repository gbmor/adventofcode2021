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

type Point struct {
	x, y int
}

func getData() [][]Point {
	data := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	if !*flagTest {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	rows := strings.Split(data, "\n")
	points := make([][]Point, len(rows))
	for i, e := range rows {
		fields := strings.Fields(e)
		left := strings.Split(fields[0], ",")
		right := strings.Split(fields[2], ",")
		x1, err := strconv.Atoi(left[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		y1, err := strconv.Atoi(left[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		x2, err := strconv.Atoi(right[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		y2, err := strconv.Atoi(right[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		points[i] = []Point{
			{
				x: x1,
				y: y1,
			},
			{
				x: x2,
				y: y2,
			},
		}
	}

	return points
}

func exec(input [][]Point, includeDiagonals bool) int {
	pointMap := make(map[int]map[int]int)
	for _, l := range input {
		if !includeDiagonals && l[0].x != l[1].x && l[0].y != l[1].y {
			continue
		}

		dx := 0
		dy := 0

		if l[0].x < l[1].x {
			dx = 1
		} else if l[0].x > l[1].x {
			dx = -1
		}
		if l[0].y < l[1].y {
			dy = 1
		} else if l[0].y > l[1].y {
			dy = -1
		}

		done := false
		for !done {
			if pointMap[l[0].x] == nil {
				pointMap[l[0].x] = make(map[int]int)
			}
			pointMap[l[0].x][l[0].y]++
			if l[0].x == l[1].x && l[0].y == l[1].y {
				done = true
			}
			l[0].x += dx
			l[0].y += dy
		}
	}

	out := 0
	for _, point := range pointMap {
		for _, count := range point {
			if count >= 2 {
				out++
			}
		}
	}

	return out
}

func main() {
	flag.Parse()

	data := getData()
	start := time.Now()
	p1 := exec(data, false)
	p1Time := time.Since(start)

	data = getData()
	start = time.Now()
	p2 := exec(data, true)
	p2Time := time.Since(start)

	fmt.Printf("Part 1:\t\t%d\t(%s)\nPart 2:\t\t%d\t(%s)\nTotal Time:\t\t%s\n", p1, p1Time, p2, p2Time, p1Time+p2Time)
}
