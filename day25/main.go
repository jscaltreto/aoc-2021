package day25

import (
	"github.com/jscaltreto/aoc-2021/lib"
)

func MoveCuces(cukes [][]byte) ([][]byte, int) {
	moved := 0
	new := make([][]byte, len(cukes))
	for y := range cukes {
		newRow := make([]byte, len(cukes[y]))
		copy(newRow, cukes[y])
		for x := range newRow {
			dest := (x + 1) % len(newRow)
			if cukes[y][x] == '>' && cukes[y][dest] == '.' {
				newRow[x] = '.'
				newRow[dest] = '>'
				moved++
			}
		}
		cukes[y] = newRow
		new[y] = make([]byte, len(newRow))
		copy(new[y], newRow)
	}
	for x := range cukes[0] {
		for y := range cukes {
			dest := (y + 1) % len(cukes)
			if cukes[y][x] == 'v' && cukes[dest][x] == '.' {
				new[y][x] = '.'
				new[dest][x] = 'v'
				moved++
			}
		}
	}

	return new, moved
}

func PartA(filename string) int {
	ct, step := 1, 0
	raw := lib.SlurpFile(filename)
	cukes := make([][]byte, len(raw))
	for i, l := range raw {
		cukes[i] = []byte(l)
	}

	for ; ct != 0; step++ {
		cukes, ct = MoveCuces(cukes)
	}
	return step
}
