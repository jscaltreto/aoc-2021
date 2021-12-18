package day17

import (
	"math"
	"regexp"

	"github.com/jscaltreto/aoc-2021/lib"
)

const (
	X1 = iota
	X2
	Y2
	Y1
	Regexp = `^target area: x=(\d+)..(\d+), y=-(\d+)..-(\d+)$`
)

type Bounds [4]int

func ReadInput(input string) Bounds {
	r := regexp.MustCompile(Regexp)
	m := r.FindStringSubmatch(input)
	b := Bounds{}
	for i := range b {
		b[i] = lib.StrToInt(m[i+1])
	}
	b[Y1] *= -1
	b[Y2] *= -1

	return b
}

func (b Bounds) FindValidStarts() int {
	intY := make(map[int][]int)
	maxSteps := 0
	for y := 0; y <= -b[Y2]; y++ {
		dist := y
		for i := 1; dist <= -b[Y2]; i++ {
			if dist >= -b[Y1] {
				if dist <= -b[Y2] {
					intY[i] = append(intY[i], -y)
				}
				if y-1 > 0 {
					posi := ((y - 1) * 2) + 1 + i
					intY[posi] = append(intY[posi], y-1)
					if posi > maxSteps {
						maxSteps = posi
					}
				}
			}
			dist += (y + i)
		}
	}

	s := int(math.Sqrt(float64(1 - (4 * -((b[X1]) * 2)))))
	minX := ((-1 + s) / 2) + 1
	values := make(map[[2]int]bool)
	for x := minX; x <= b[X2]; x++ {
		momentum, dist, step := x, x, 1
		for step = 1; momentum > 0 && dist <= b[X2]; step++ {
			if dist >= b[X1] {
				for _, y := range intY[step] {
					values[[2]int{x, y}] = true
				}
			}
			momentum -= 1
			dist += momentum
		}
		if momentum == 0 {
			for ; step <= maxSteps; step++ {
				for _, y := range intY[step] {
					values[[2]int{x, y}] = true
				}
			}
		}
	}
	return len(values)
}

func PartA(filename string) int {
	data := lib.SlurpFile(filename)[0]
	b := ReadInput(data)
	m := -b[Y2]
	return ((m - 1) * m) / 2
}

func PartB(filename string) int {
	data := lib.SlurpFile(filename)[0]
	b := ReadInput(data)
	return b.FindValidStarts()
}
