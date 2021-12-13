package day04

import (
	"log"
	"strconv"
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

type BoardPos struct {
	board *Board
	row   int
	col   int
}

func (b *BoardPos) CheckRow() bool {
	for _, num := range b.board.numbers[b.row] {
		if !num.called {
			return false
		}
	}
	return true
}

func (b *BoardPos) CheckCol() bool {
	for _, row := range b.board.numbers {
		if !row[b.col].called {
			return false
		}
	}
	return true
}

func (b *BoardPos) Check() bool {
	return b.CheckRow() || b.CheckCol()
}

type Number struct {
	number int
	called bool
	boards []*BoardPos
}

var numMap map[int]*Number

func addNumberToBoard(num int, boardPos *BoardPos) *Number {
	number := &Number{number: num}
	if n, found := numMap[num]; found {
		number = n
	} else {
		numMap[num] = number
	}
	number.boards = append(number.boards, boardPos)
	return number
}

func (n *Number) Call() *Board {
	n.called = true
	var winner *Board
	for _, b := range n.boards {
		if !b.board.won && b.Check() {
			b.board.won = true
			winner = b.board
		}
	}
	return winner
}

type Board struct {
	numbers [5][5]*Number
	won     bool
}

func loadBoard(data []string) {
	board := &Board{}
	for rid, row := range data {
		for cid, numStr := range strings.Fields(row) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("Invalid number %s", numStr)
			}
			board.numbers[rid][cid] = addNumberToBoard(
				num,
				&BoardPos{board: board, row: rid, col: cid},
			)
		}
	}
}

func (b *Board) Score() int {
	score := 0
	for _, row := range b.numbers {
		for _, num := range row {
			if !num.called {
				score += num.number
			}
		}
	}
	return score
}

func FindWinner(filename string, lastWinner bool) int {
	numMap = make(map[int]*Number)
	input := lib.SlurpFile(filename)

	for i := 2; i < len(input); i += 6 {
		loadBoard(input[i : i+5])
	}

	score := 0
	for _, callStr := range strings.Split(input[0], ",") {
		call, err := strconv.Atoi(callStr)
		if err != nil {
			log.Fatalf("Invalid number %s", callStr)
		}
		if num, found := numMap[call]; found {
			if bingo := num.Call(); bingo != nil {
				score = bingo.Score() * call
				if !lastWinner {
					return score
				}
			}
		}
	}
	return score
}

func PartA(filename string) int {
	return FindWinner(filename, false)
}

func PartB(filename string) int {
	return FindWinner(filename, true)
}
