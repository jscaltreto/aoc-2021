package day24

import (
	"strconv"
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetSerial(filename string, low bool) int {
	monad := lib.SlurpFile(filename)
	checks := [14][2]int{}
	for l := 0; l < len(monad); l += 18 {
		checks[l/18] = [2]int{
			lib.StrToInt(strings.Fields(monad[l+5])[2]),
			lib.StrToInt(strings.Fields(monad[l+15])[2]),
		}
	}
	stack := [][2]int{}
	digits := [2][]byte{make([]byte, 14), make([]byte, 14)}
	for idx, c := range checks {
		if c[0] > 0 {
			stack = append(stack, [2]int{c[1], idx})
		} else {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			o := i[0] + c[0]
			digits[0][idx] = strconv.Itoa(Min(9, 9+o))[0]
			digits[0][i[1]] = strconv.Itoa(Min(9, 9-o))[0]
			digits[1][idx] = strconv.Itoa(Max(1, 1+o))[0]
			digits[1][i[1]] = strconv.Itoa(Max(1, 1-o))[0]
		}
	}
	if low {
		return lib.StrToInt(string(digits[1]))
	}
	return lib.StrToInt(string(digits[0]))
}

func PartA(filename string) int {
	return GetSerial(filename, false)
}

func PartB(filename string) int {
	return GetSerial(filename, true)
}
