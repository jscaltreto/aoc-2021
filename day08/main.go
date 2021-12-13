package day08

import (
	"sort"
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

const (
	A     = 1 << 0
	B     = 1 << 1
	C     = 1 << 2
	D     = 1 << 3
	E     = 1 << 4
	F     = 1 << 5
	G     = 1 << 6
	ALL   = (1 << 7) - 1
	ONE   = C | F
	FOUR  = B | C | D | F
	SEVEN = A | C | F
)

var SumMap map[int]byte = map[int]byte{
	A + (F | C) + (F | C) + (E | G) + (E | G) + (B | D): '0',
	(F | C) + (F | C):                                             '1',
	A + (F | C) + (B | D) + (E | G) + (E | G):                     '2',
	A + (F | C) + (F | C) + (E | G) + (B | D):                     '3',
	(B | D) + (B | D) + (F | C) + (F | C):                         '4',
	A + (B | D) + (B | D) + (F | C) + (E | G):                     '5',
	A + (F | C) + (E | G) + (E | G) + (B | D) + (B | D):           '6',
	A + (F | C) + (F | C):                                         '7',
	A + (F | C) + (F | C) + (E | G) + (E | G) + (B | D) + (B | D): '8',
	A + (F | C) + (F | C) + (E | G) + (B | D) + (B | D):           '9',
}

type Sample struct {
	signals []string
	outputs []string
}

// Once we remove the segmets only used by 1, 4, & 7 we can narrow
// down the possible segmets for each letter down to two possibities.
// Then, for each number, the sum of the value of those possibilities
// turns out to be unique, so it can be looked up in a table (see SumMap).
func (s *Sample) Analyze() map[rune]byte {
	sigs := s.signals
	sort.Slice(sigs, func(i, j int) bool {
		return len(sigs[i]) < len(sigs[j])
	})
	bitMap := map[rune]byte{
		'a': ALL, 'b': ALL, 'c': ALL, 'd': ALL, 'e': ALL, 'f': ALL, 'g': ALL,
	}
	for c := range bitMap {
		if strings.ContainsRune(sigs[0], c) {
			bitMap[c] &= ONE
		} else {
			bitMap[c] &^= ONE
		}
		if strings.ContainsRune(sigs[1], c) {
			bitMap[c] &= SEVEN
		} else {
			bitMap[c] &^= SEVEN
		}
		if strings.ContainsRune(sigs[2], c) {
			bitMap[c] &= FOUR
		} else {
			bitMap[c] &^= FOUR
		}
	}
	return bitMap
}

func (s *Sample) GetDigits() int {
	bitMap := s.Analyze()
	digits := []byte{}
	for _, output := range s.outputs {
		sum := 0
		for _, c := range output {
			sum += int(bitMap[c])
		}
		digits = append(digits, SumMap[sum])
	}
	return lib.StrToInt(string(digits))
}

func ParseSample(sample string) *Sample {
	sigout := strings.Split(sample, " | ")
	return &Sample{
		signals: strings.Fields(sigout[0]),
		outputs: strings.Fields(sigout[1]),
	}
}

func ReadDisplays(filename string) []*Sample {
	samples := []*Sample{}
	for _, sampleStr := range lib.SlurpFile(filename) {
		samples = append(samples, ParseSample(sampleStr))
	}
	return samples
}

func PartA(filename string) int {
	total := 0
	for _, sample := range ReadDisplays(filename) {
		for _, output := range sample.outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				total++
			}
		}
	}
	return total
}

func PartB(filename string) int {
	total := 0
	for _, sample := range ReadDisplays(filename) {
		total += sample.GetDigits()
	}
	return total
}
