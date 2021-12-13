package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SlurpFile(filename string) []string {
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func LisOfNumbers(filename string) []int {
	contents := SlurpFile(filename)
	nums := []int{}
	for _, str := range strings.Split(contents[0], ",") {
		nums = append(nums, StrToInt(str))
	}
	return nums
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Not a number %s", s)
	}
	return i
}
