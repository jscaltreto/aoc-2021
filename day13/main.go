package day13

import (
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

type Dot [2]int
type Dots map[Dot]bool

func LoadDots(data []string) (*Dots, int) {
	dots := Dots{}
	for idx, l := range data {
		if l == "" {
			return &dots, idx + 1
		}
		coord := strings.Split(l, ",")
		cx, cy := lib.StrToInt(coord[0]), lib.StrToInt(coord[1])
		dots[Dot{cx, cy}] = true
	}
	return nil, 0
}

func (d Dots) DoFold(inst string) {
	axis := inst[11] - 'x'
	fold := lib.StrToInt(inst[13:])
	for dot := range d {
		if dot[axis] > fold {
			delete(d, dot)
			dot[axis] = fold - (dot[axis] - fold)
			d[dot] = true
		}
	}
}

func (d Dots) Print() string {
	maxX, maxY := 0, 0
	for dot := range d {
		if dot[0] > maxX {
			maxX = dot[0]
		}
		if dot[1] > maxY {
			maxY = dot[1]
		}
	}
	out := "\n"
	for iy := 0; iy <= maxY; iy++ {
		for ix := 0; ix <= maxX; ix++ {
			if _, f := d[Dot{ix, iy}]; f {
				out += "#"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

func PartA(filename string) int {
	data := lib.SlurpFile(filename)
	dots, idx := LoadDots(data)
	dots.DoFold(data[idx])
	return len(*dots)
}

func PartB(filename string) string {
	data := lib.SlurpFile(filename)
	dots, idx := LoadDots(data)
	for _, inst := range data[idx:] {
		dots.DoFold(inst)
	}
	return dots.Print()
}
