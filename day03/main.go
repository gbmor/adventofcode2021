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

var partTwoByteMap = map[bool]map[bool]byte{
	true: {
		true:  byte('1'),
		false: byte('0'),
	},
	false: {
		true:  byte('0'),
		false: byte('1'),
	},
}

func getReport() []string {
	data := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

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

func powerConsumption(input []string) int64 {
	gamma := strings.Builder{}
	epsilon := strings.Builder{}
	for j := 0; j < len(input[0]); j++ {
		one := 0
		zero := 0
		for i := 0; i < len(input); i++ {
			if string(input[i][j]) == "1" {
				one++
			} else {
				zero++
			}
		}
		if one > zero {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		} else {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		}
	}
	gammaD, err := strconv.ParseInt(gamma.String(), 2, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	epsilonD, err := strconv.ParseInt(epsilon.String(), 2, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return gammaD * epsilonD
}

func lifeSupportRating(oxygen bool, arr []string, current int) []string {
	arrLenHalf := float64(len(arr)) / 2
	ones := make([]int, 0)
	for i := range arr[0] {
		ones = append(ones, 0)
		for _, j := range arr {
			if j[i] == byte('1') {
				ones[i]++
			}
		}
	}
	pick := float64(ones[current])
	common := partTwoByteMap[pick >= arrLenHalf][oxygen]
	out := make([]string, 0)
	for _, e := range arr {
		if e[current] == common {
			out = append(out, e)
		}
	}
	return out
}

func main() {
	flag.Parse()

	commands := getReport()
	start := time.Now()
	p1 := powerConsumption(commands)
	p1time := time.Since(start)
	start = time.Now()

	for i := range commands[0] {
		commands = lifeSupportRating(true, commands, i)
		if len(commands) == 1 {
			break
		}
	}
	OxygenGeneratorRating, err := strconv.ParseInt(commands[0], 2, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	commands = getReport()
	for i := range commands[0] {
		commands = lifeSupportRating(false, commands, i)
		if len(commands) == 1 {
			break
		}
	}
	CO2ScrubberRating, err := strconv.ParseInt(commands[0], 2, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	p2 := OxygenGeneratorRating * CO2ScrubberRating
	p2time := time.Since(start)

	fmt.Printf("Part 1:\t\t%d\t(%s)\nPart 2:\t\t%d\t(%s)\n", p1, p1time, p2, p2time)
	fmt.Printf("Total Time:\t%s\n", p1time+p2time)
}
