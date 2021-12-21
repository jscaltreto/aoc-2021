package day20

import (
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

var neighbors [9][2]int = [9][2]int{
	{1, 1}, {0, 1}, {-1, 1},
	{1, 0}, {0, 0}, {-1, 0},
	{1, -1}, {0, -1}, {-1, -1},
}

func ProcessImage(filename string, steps int) int {
	// Read the input image
	data := lib.SlurpFile(filename)
	algo := data[0]
	image := data[2:]

	// Enhance!
	for i := 0; i < steps; i++ {
		enhanced := make([]string, len(image)+2)
		for y := range enhanced {
			line := make([]byte, len(image[0])+2)
			for x := range line {
				eidx := 0
				for nidx, nt := range neighbors {
					ny, nx := y+nt[1]-1, x+nt[0]-1
					outOfBounds := (ny < 0 || nx < 0 || ny >= len(image) || nx >= len(image[ny]))
					if (!outOfBounds && image[ny][nx] == '#') || (outOfBounds && algo[0] == '#' && i%2 == 1) {
						eidx |= 1 << nidx
					}
				}
				line[x] = algo[eidx]
			}
			enhanced[y] = string(line)
		}
		image = enhanced
	}

	// Count the lit pixels
	ct := 0
	for _, line := range image {
		ct += strings.Count(line, "#")
	}
	return ct
}

func PartA(filename string) int {
	return ProcessImage(filename, 2)
}

func PartB(filename string) int {
	return ProcessImage(filename, 50)
}
