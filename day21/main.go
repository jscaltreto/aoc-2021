package day21

import (
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

func PartA(filename string) int {
	input := lib.SlurpFile(filename)
	players := [2][2]int{}
	for i := range players {
		players[i] = [2]int{lib.StrToInt(strings.Split(input[i], ": ")[1]) - 1, 0}
	}

	d := 1
	for p := 0; ; p = (p + 1) % 2 {
		players[p][0] = (players[p][0] + (3 * d) + 3) % 10
		players[p][1] += players[p][0] + 1
		d += 3
		if players[p][1] >= 1000 {
			return players[(p+1)%2][1] * (d - 1)
		}
	}
}

// Pre-computed the number of universes created from
// the results of any 3 dice rolls.
var resultUniverses [7][2]int = [7][2]int{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}

func Turn(players [2][3]int, prevUniverses int) [2]int {
	won := [2]int{}

	for _, resultP1 := range resultUniverses {
		nextPlayers := [2][3]int{}
		nextPos := (players[0][0] + resultP1[0]) % 10
		nextPlayers[0] = [3]int{nextPos, players[0][1] + nextPos + 1, players[0][2] * resultP1[1] * prevUniverses}
		if nextPlayers[0][1] >= 21 {
			won[0] += nextPlayers[0][2]
			continue
		}
		for _, resultP2 := range resultUniverses {
			nextPos := (players[1][0] + resultP2[0]) % 10
			nextPlayers[1] = [3]int{nextPos, players[1][1] + nextPos + 1, players[1][2] * (resultP2[1] * resultP1[1])}
			if nextPlayers[1][1] >= 21 {
				won[1] += nextPlayers[1][2]
			} else {
				w := Turn(nextPlayers, resultP2[1])
				won[0] += w[0]
				won[1] += w[1]
			}
		}
	}
	return won
}

func PartB(filename string) int {
	input := lib.SlurpFile(filename)
	players := [2][3]int{}
	for i := range players {
		players[i] = [3]int{lib.StrToInt(strings.Split(input[i], ": ")[1]) - 1, 0, 1}
	}

	w := Turn(players, 1)
	if w[0] > w[1] {
		return w[0]
	}
	return w[1]
}
