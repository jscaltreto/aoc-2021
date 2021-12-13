package day07

import (
	"math"
	"sort"

	"github.com/jscaltreto/aoc-2021/lib"
)

func FindMedian(crabs []int) int {
	sort.Slice(crabs, func(i, j int) bool {return crabs[i] < crabs[j]})
	numCrabs := len(crabs)
	m := numCrabs / 2
	avg := 0
	if numCrabs % 2 != 0 {
		avg = (crabs[m-1] + crabs[m] + crabs[m+1])/3
	} else {
		avg = (crabs[m] + crabs[m-1])/2
	}
	return avg
}

func PartA(filename string) int {
	crabs := lib.LisOfNumbers(filename)
	m := FindMedian(crabs)
	fuel := 0
	for _, c := range crabs {
		fuel += int(math.Abs(float64(m - c)))
	}
	return fuel
}

func GetMinimumFuel(crabs []int, start, currentCost, delta int) int {
	checkPos := []int{start, start + delta}
	minFuel := math.MaxInt
	for len(checkPos) > 0 {
		pos := checkPos[0]
		fuel := 0
		for _, c := range crabs {
			fuel += fuelMap[int(math.Abs(float64(pos - c)))]
		}
		if fuel < minFuel {
			minFuel = fuel
			checkPos = append(checkPos, pos + delta)
		}
		checkPos = checkPos[1:]
	}
	return minFuel
}

var fuelMap []int

func MakeFuelMap(crabs []int) {
	maxCrab := crabs[len(crabs) - 1]
	fuelMap = make([]int, maxCrab)
	for i := 1; i < maxCrab; i++ {
		fuelMap[i] = fuelMap[i-1] + i
	}
}

func PartB(filename string) int {
	crabs := lib.LisOfNumbers(filename)
	m := FindMedian(crabs)
	MakeFuelMap(crabs)
	right := GetMinimumFuel(crabs, m, 0, 1)
	left := GetMinimumFuel(crabs, m, 0, -1)
	if left < right {
		return left
	}
	return right
}
