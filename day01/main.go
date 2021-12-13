package day01

import (
	"log"
	"strconv"

	"github.com/jscaltreto/aoc-2021/lib"
)

func PartA(filename string) uint {
	var increases uint
	last := 0
	for idx, depthStr := range lib.SlurpFile(filename) {
		depth, err := strconv.Atoi(depthStr)
		if err != nil {
			log.Fatal(err)
		}
		if idx > 0 && depth > last {
			increases++
		}
		last = depth
	}
	return increases
}

func PartB(filename string) uint {
	var increases uint
	lastSum := 0

	depthStrs := lib.SlurpFile(filename)
	d := make([]int, len(depthStrs))
	for i, depthStr := range depthStrs {
		depth, err := strconv.Atoi(depthStr)
		if err != nil {
			log.Fatal(err)
		}
		d[i] = depth

		if i > 1 {
			sum := d[i-2] + d[i-1] + d[i]
			if i > 2 && sum > lastSum {
				increases++
			}
			lastSum = sum
		}
	}
	return increases

}
