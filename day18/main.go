package day18

import (
	"fmt"
	"regexp"

	"github.com/jscaltreto/aoc-2021/lib"
)

var explodable regexp.Regexp = *regexp.MustCompile(`\[(\d+),(\d+)\]`)
var digits regexp.Regexp = *regexp.MustCompile(`\d+`)
var tenplus regexp.Regexp = *regexp.MustCompile(`\d{2,}`)

func Split(sn string) string {
	for _, n := range tenplus.FindAllStringIndex(sn, 1) {
		num := lib.StrToInt(sn[n[0]:n[1]])
		return fmt.Sprintf("%s[%d,%d]%s",
			sn[:n[0]],
			num/2,
			(num/2)+(num%2),
			sn[n[1]:],
		)
	}
	return ""
}

func Explode(sn string) string {
	for _, n := range explodable.FindAllStringSubmatchIndex(sn, -1) {
		n1 := lib.StrToInt(sn[n[2]:n[3]])
		n2 := lib.StrToInt(sn[n[4]:n[5]])

		levels := 0
		for _, c := range sn[:n[0]] {
			switch c {
			case '[':
				levels++
			case ']':
				levels--
			}
		}
		if levels >= 4 {
			newsn := ""
			if leftNums := digits.FindAllStringIndex(sn[:n[0]], -1); len(leftNums) > 0 {
				m := leftNums[len(leftNums)-1]
				num := lib.StrToInt(sn[m[0]:m[1]]) + n1
				newsn = fmt.Sprintf("%s%d%s0", sn[:m[0]], num, sn[m[1]:n[0]])
			} else {
				newsn = sn[:n[0]] + "0"
			}
			rem := sn[n[1]:]
			if m := digits.FindStringIndex(rem); len(m) > 0 {
				num := lib.StrToInt(rem[m[0]:m[1]]) + n2
				newsn += fmt.Sprintf("%s%d%s", rem[:m[0]], num, rem[m[1]:])
			} else {
				newsn += rem
			}
			return newsn
		}
	}
	return ""
}

func GetMagnitude(sn string) int {
	for {
		if n := explodable.FindStringSubmatchIndex(sn); len(n) > 0 {
			n1 := lib.StrToInt(sn[n[2]:n[3]])
			n2 := lib.StrToInt(sn[n[4]:n[5]])
			sn = fmt.Sprintf("%s%d%s", sn[:n[0]], (n1*3)+(n2*2), sn[n[1]:])
			continue
		}
		return lib.StrToInt(sn)
	}
}

func Reduce(sn string) string {
	for {
		if new := Explode(sn); new != "" {
			sn = new
			continue
		}
		if new := Split(sn); new != "" {
			sn = new
			continue
		}
		return sn
	}
}

func Add(a, b string) string {
	return Reduce(fmt.Sprintf("[%s,%s]", a, b))
}

func AddAll(numbers []string) string {
	a := numbers[0]
	for _, b := range numbers[1:] {
		a = Add(a, b)
		continue
	}
	return a
}

func FindMaxPair(numbers []string) int {
	max := 0
	for i, a := range numbers {
		for _, b := range numbers[i+1:] {
			for _, mag := range [2]int{GetMagnitude(Add(a, b)), GetMagnitude(Add(b, a))} {
				if mag > max {
					max = mag
				}
			}
		}
	}
	return max
}

func PartA(filename string) int {
	data := lib.SlurpFile(filename)
	return GetMagnitude(AddAll(data))
}

func PartB(filename string) int {
	data := lib.SlurpFile(filename)
	return FindMaxPair(data)
}
