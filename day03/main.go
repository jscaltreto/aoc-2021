package day03

import (
	"log"
	"strconv"

	"github.com/jscaltreto/aoc-2021/lib"
)

func PartA(filename string) int {
	input := lib.SlurpFile(filename)
	counts := make([]int, len(input[0]))
	for _, bits := range input {
		for i, b := range bits {
			if b == '1' {
				counts[i]++
			}
		}
	}
	var gamma, epsilon uint
	for i, ct := range counts {
		if ct*2 >= len(input) {
			gamma |= 1 << (len(counts) - i - 1)
		} else {
			epsilon |= 1 << (len(counts) - i - 1)
		}
	}
	return int(epsilon * gamma)
}

func findReading(readings []string, keepMost bool) int {
	pos := 0
	for len(readings) > 1 {
		buckets := map[bool][]string{
			false: {}, // Zeroes
			true:  {}, // Ones
		}
		for _, n := range readings {
			buckets[n[pos] == '1'] = append(buckets[n[pos] == '1'], n)
		}
		readings = buckets[(len(buckets[true])*2 >= len(readings)) == keepMost]
		pos++
	}
	i, err := strconv.ParseInt(readings[0], 2, 64)
	if err != nil {
		log.Fatalf("Invalid binary number %s", readings[0])
	}
	return int(i)
}

func PartB(filename string) int {
	input := lib.SlurpFile(filename)
	o2 := findReading(input, true)
	co2 := findReading(input, false)

	return o2 * co2
}
