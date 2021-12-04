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

type BingoBoard struct {
	Nums        [][]int
	BingoNumber int

	// Answer is the sum of the uncalled numbers multiplied by
	// the called number that earned this board bingo.
	Answer int
}

func makeBoards(input []string) []BingoBoard {
	out := make([]BingoBoard, len(input))

	for i, e := range input {
		out[i].Nums = make([][]int, 5)
		rows := strings.Split(e, "\n")

		for j, k := range rows {
			out[i].Nums[j] = make([]int, 5)
			nums := strings.Fields(k)

			for jfc, stop := range nums {
				n, err := strconv.Atoi(stop)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				out[i].Nums[j][jfc] = n
			}
		}
	}

	return out
}

func (b *BingoBoard) bingoThisRound(currentNumber int) bool {
	for i := 0; i < 5; i++ {
		colCount := 0
		rowCount := 0
		for j := 0; j < 5; j++ {
			if b.Nums[j][i] == -1 {
				colCount++
			}
			if b.Nums[i][j] == -1 {
				rowCount++
			}
		}
		if colCount == 5 || rowCount == 5 {
			b.BingoNumber = currentNumber
			return true
		}
	}

	return false
}

func (b *BingoBoard) alreadyHasBingo() bool {
	return b.BingoNumber > 0
}

func getBingoData() ([]int, []BingoBoard) {
	data := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

	if !*flagTest {
		b, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = string(b)
	}

	splitSec := strings.Split(data, "\n\n")
	numStr := strings.Split(splitSec[0], ",")
	numbers := make([]int, len(numStr))

	for i, e := range numStr {
		n, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		numbers[i] = n
	}
	return numbers, makeBoards(splitSec[1:])
}

// This solution is a tragedy, but it works.
// First returned int is the first winning board's non-picked-numbers summed and multiplied by
// the picked number that caused it to win. Second is the same but for the last winning board.
func cheatAtBingo(picks []int, boards []BingoBoard) (int, int) {
	boardsInOrder := make([]BingoBoard, 0)

	for _, e := range picks {
		for j, k := range boards {
			if k.alreadyHasBingo() {
				continue
			}
			for _, dearGod := range k.Nums {
				numFound := false
				for please, forgiveMe := range dearGod {
					if forgiveMe == e {
						dearGod[please] = -1
						numFound = true
						break
					}
				}
				if numFound {
					break
				}
			}
			if boards[j].bingoThisRound(e) {
				boardsInOrder = append(boardsInOrder, boards[j])
			}
		}
		if len(boardsInOrder) == len(boards) {
			break
		}
	}

	theTwo := []BingoBoard{
		boardsInOrder[0],
		boardsInOrder[len(boardsInOrder)-1],
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := range theTwo {
				pos := theTwo[k].Nums[i][j]
				if pos != -1 {
					theTwo[k].Answer += pos
				}
			}
		}
	}

	theTwo[0].Answer *= theTwo[0].BingoNumber
	theTwo[1].Answer *= theTwo[1].BingoNumber

	return theTwo[0].Answer, theTwo[1].Answer
}

func main() {
	flag.Parse()

	picks, boards := getBingoData()
	start := time.Now()
	winningBoard, losingBoard := cheatAtBingo(picks, boards)
	execTime := time.Since(start)

	fmt.Printf("Part 1:\t%d\nPart 2:\t%d\nTime:\t%s\n", winningBoard, losingBoard, execTime)
}
