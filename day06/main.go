package day06

import (
	"github.com/jscaltreto/aoc-2021/lib"
)

func Simulate(filename string, days int) int {
	spawnDays := make([]int, 8)
	for _, day := range lib.LisOfNumbers(filename) {
		spawnDays[day]++
	}

	nextCycle := 0
	for i := 0; i <= days; i++ {
		cur := i % 8
		lastCycle := nextCycle
		nextCycle = spawnDays[cur]
		spawnDays[(cur+7)%8] += spawnDays[cur] // Respawn
		spawnDays[cur] = lastCycle // New fish
	}

	total := 0
	for _, day := range spawnDays {
		total += day
	}
	return total
}

func PartA(filename string) int {
	return Simulate(filename, 80)
}

func PartB(filename string) int {
	return Simulate(filename, 256)
}
