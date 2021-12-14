package day14

import (
	"math"

	"github.com/jscaltreto/aoc-2021/lib"
)

type Rules map[[2]byte]byte
type Pairs map[[2]byte]int
type ElemCt map[byte]int

type Poly struct {
	rules Rules
	pairs Pairs
	count ElemCt
}

func LoadElems(filename string) *Poly {
	data := lib.SlurpFile(filename)
	poly := &Poly{
		pairs: make(Pairs),
		count: ElemCt{data[0][len(data[0])-1]: 1},
		rules: make(Rules),
	}
	for i := 0; i < len(data[0])-1; i++ {
		poly.pairs[[2]byte{data[0][i], data[0][i+1]}]++
		poly.count[data[0][i]]++
	}

	for _, l := range data[2:] {
		poly.rules[[2]byte{l[0], l[1]}] = l[6]
	}
	return poly
}

func (p *Poly) Step() {
	nextPairs := make(Pairs)
	for pair, ct := range p.pairs {
		if ct > 0 {
			if e, f := p.rules[pair]; f {
				p.count[e] += ct
				nextPairs[[2]byte{pair[0], e}] += ct
				nextPairs[[2]byte{e, pair[1]}] += ct
			} else {
				nextPairs[pair] += ct
			}
		}
	}
	p.pairs = nextPairs
}

func GetDiff(filename string, iters int) int {
	poly := LoadElems(filename)
	for i := 0; i < iters; i++ {
		poly.Step()
	}
	min, max := math.MaxInt, 0
	for _, ct := range poly.count {
		if ct > max {
			max = ct
		}
		if ct < min {
			min = ct
		}
	}
	return max - min
}

func PartA(filename string) int {
	return GetDiff(filename, 10)
}

func PartB(filename string) int {
	return GetDiff(filename, 40)
}
