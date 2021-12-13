package day05

import (
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

type Coord [2]int

func getWalkFunc(d int) func(i int) int {
	if d == 0 {
		return func(i int) int { return i }
	} else if d < 0 {
		return func(i int) int { return i - 1 }
	} else {
		return func(i int) int { return i + 1 }
	}
}

func Walk(start, end Coord, diag bool) []Coord {
	x1, x2, y1, y2 := start[0], end[0], start[1], end[1]
	dx, dy := x2-x1, y2-y1
	if !diag && !(dx == 0 || dy == 0) {
		return nil
	}
	fx, fy := getWalkFunc(dx), getWalkFunc(dy)
	cur := Coord{x1, y1}
	coords := []Coord{cur}
	for cur != end {
		c := Coord{fx(cur[0]), fy(cur[1])}
		coords = append(coords, c)
		cur = c
	}
	return coords
}

func ReadCoords(cStr string) (Coord, Coord) {
	xys := [2]Coord{}
	for i, c := range strings.Split(cStr, " -> ") {
		xy := strings.Split(c, ",")
		xys[i][0], xys[i][1] = lib.StrToInt(xy[0]), lib.StrToInt(xy[1])
	}
	return xys[0], xys[1]
}

func FindVents(filename string, diag bool) int {
	input := lib.SlurpFile(filename)
	ventMap := make(map[Coord]int)
	for _, line := range input {
		xy1, xy2 := ReadCoords(line)
		for _, c := range Walk(xy1, xy2, diag) {
			if _, found := ventMap[c]; found {
				ventMap[c]++
			} else {
				ventMap[c] = 1
			}
		}
	}
	count := 0
	for _, vents := range ventMap {
		if vents >= 2 {
			count++
		}
	}
	return count
}

func PartA(filename string) int {
	return FindVents(filename, false)
}

func PartB(filename string) int {
	return FindVents(filename, true)
}
