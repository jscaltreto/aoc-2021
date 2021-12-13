package day02

import (
	"log"
	"strconv"
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

const (
	F = "forward"
	D = "down"
	U = "up"
)

func parseCommand(cmdStr string) (string, int) {
		cmd := strings.Fields(cmdStr)
		dist, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatalf("Invalid command %s", cmd)
		}
		return cmd[0], dist
}

func PartA(filename string) int {
	x, z := 0, 0
	for _, cmdStr := range lib.SlurpFile(filename) {
		inst, dist := parseCommand(cmdStr)

		switch inst {
		case F:
			x += dist
		case D:
			z += dist
		case U:
			z -= dist
		}
	}
	return x * z
}

func PartB(filename string) int {
	x, z, aim := 0, 0, 0
	for _, cmdStr := range lib.SlurpFile(filename) {
		inst, dist := parseCommand(cmdStr)

		switch inst {
		case F:
			x += dist
			z += (dist * aim)
		case D:
			aim += dist
		case U:
			aim -= dist
		}
	}
	return x * z
}
