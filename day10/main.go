package day10

import (
	"sort"

	"github.com/jscaltreto/aoc-2021/lib"
)

var match map[byte]byte = map[byte]byte{
	'[': ']',
	'(': ')',
	'<': '>',
	'{': '}',
}

var pts map[byte]int = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func scoreLine(l string) (int, []byte) {
	chunks := []byte{}
	for _, c := range l {
		if c == '[' || c == '{' || c == '(' || c == '<' {
			chunks = append(chunks, byte(c))
		} else if match[chunks[len(chunks)-1]] != byte(c) {
			return pts[byte(c)], chunks
		} else {
			chunks = chunks[:len(chunks)-1]
		}
	}
	return 0, chunks
}

func PartA(filename string) int {
	score := 0
	for _, l := range lib.SlurpFile(filename) {
		s, _ := scoreLine(l)
		score += s
	}
	return score
}

func PartB(filename string) int {
	scores := []int{}
	for _, l := range lib.SlurpFile(filename) {
		s, o := scoreLine(l)
		if s != 0 {
			continue
		}
		score := 0
		for i := len(o) - 1; i >= 0; i-- {
			score *= 5
			score += pts[o[i]]
		}
		scores = append(scores, score)
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	return scores[len(scores)/2]
}
