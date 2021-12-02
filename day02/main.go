package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getCommands(test bool) []string {
	if test {
		return []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}
	}
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	split := bytes.Split(b, []byte("\n"))
	out := make([]string, len(split))
	for i, e := range split {
		out[i] = string(e)
	}
	return out
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
	commands := getCommands(false)
	start := time.Now()
	p1, p2 := exec(commands)
	finish := time.Since(start)
	fmt.Printf("Part 1: %d\nPart 2: %d\nTime: %s\n", p1, p2, finish)
}
