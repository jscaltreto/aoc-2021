package day09

import (
	"sort"

	"github.com/jscaltreto/aoc-2021/lib"
)

type Coord struct {
	x int
	y int
}

type Map []string

func (m Map) Up(p *Coord) *Coord {
	if p.x > 0 {
		return &Coord{p.x - 1, p.y}
	}
	return nil
}
func (m Map) Down(p *Coord) *Coord {
	if p.x < len(m)-1 {
		return &Coord{p.x + 1, p.y}
	}
	return nil
}
func (m Map) Left(p *Coord) *Coord {
	if p.y > 0 {
		return &Coord{p.x, p.y - 1}
	}
	return nil
}
func (m Map) Right(p *Coord) *Coord {
	if p.y < len(m[p.x])-1 {
		return &Coord{p.x, p.y + 1}
	}
	return nil
}
func (m Map) Adj(p *Coord) [4]*Coord {
	return [4]*Coord{m.Up(p), m.Down(p), m.Left(p), m.Right(p)}
}
func (m Map) Get(c *Coord) byte {
	return m[c.x][c.y]
}

func FindLowPoints(filename string) ([][2]int, []string) {
	input := lib.SlurpFile(filename)
	lowPoints := [][2]int{}
	for ri, row := range input {
		for ci, cr := range row {
			c := byte(cr)
			if (ci > 0 && row[ci-1] <= c) ||
				(ci < len(row)-1 && row[ci+1] <= c) ||
				(ri > 0 && input[ri-1][ci] <= c) ||
				(ri < len(input)-1 && input[ri+1][ci] <= c) {
				continue
			}
			lowPoints = append(lowPoints, [2]int{ri, ci})
		}
	}
	return lowPoints, input
}

func PartA(filename string) int {
	lowPoints, input := FindLowPoints(filename)
	risk := 0
	for _, lp := range lowPoints {
		risk += lib.StrToInt(string(input[lp[0]][lp[1]])) + 1
	}
	return risk
}

func MapBasin(input []string, lp [2]int) int {
	m := Map(input)
	toCheck := []*Coord{{lp[0], lp[1]}}
	checked := map[Coord]bool{}
	basinSize := 1
	for len(toCheck) > 0 {
		c := m.Get(toCheck[0])
		for _, bc := range m.Adj(toCheck[0]) {
			if bc != nil {
				if _, f := checked[*bc]; !f {
					b := m.Get(bc)
					if b != '9' && b >= c {
						basinSize++
						toCheck = append(toCheck, bc)
						checked[*bc] = true
					}
				}
			}
		}
		toCheck = toCheck[1:]
	}
	return basinSize
}

func PartB(filename string) int {
	lowPoints, input := FindLowPoints(filename)
	basins := []int{}
	for _, lp := range lowPoints {
		basins = append(basins, MapBasin(input, lp))
	}
	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j]})
	return basins[0] * basins[1] * basins[2]
}
