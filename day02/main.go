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

func getCommands() []string {
	data := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	if !*flagTest {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	return strings.Split(data, "\n")
}

func exec(input []string) (int, int) {
	horiz := 0
	depth := 0
	aim := 0
	horiz2 := 0
	depth2 := 0

	for _, e := range input {
		split := strings.Fields(e)
		chg, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		switch split[0] {
		case "forward":
			horiz += chg
			horiz2 += chg
			depth2 += aim * chg
		case "down":
			depth += chg
			aim += chg
		case "up":
			depth -= chg
			aim -= chg
		}
	}
	return horiz * depth, horiz2 * depth2
}

func main() {
	flag.Parse()

	commands := getCommands()
	start := time.Now()
	p1, p2 := exec(commands)
	finish := time.Since(start)

	fmt.Printf("Part 1: %d\nPart 2: %d\nTime: %s\n", p1, p2, finish)
}
