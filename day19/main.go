package day19

import (
	"math"
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

var scannerDist map[int][3]int

func makeRotations() [][]int {
	rotations := [][]int{}
	order := [][3]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}
	rot := [2]int{1, -1}
	for _, xr := range rot {
		for _, yr := range rot {
			for _, zr := range rot {
				for _, o := range order {
					rotations = append(rotations, []int{
						xr,
						yr,
						zr,
						o[0],
						o[1],
						o[2],
					})
				}
			}
		}
	}
	return rotations
}

var rotations [][]int = makeRotations()

type Beacon struct {
	s *Scanner
	d [3]int
	r map[[3]int]bool
	o int
}

type Beacons map[[3]int]*Beacon

func (b *Beacon) AddRelation(b2 *Beacon) {
	r := [3]int{}
	for i, d := range b.d {
		r[i] = d - b2.d[i]
	}
	b.r[r] = true
}

type Scanner struct {
	id        int
	b         Beacons
	rotations map[int]*Scanner
}

func (s *Scanner) BuildRelations() {
	for _, b1 := range s.b {
		for _, b2 := range s.b {
			if b2 != b1 {
				b1.AddRelation(b2)
				b2.AddRelation(b1)

			}
		}
	}
}

func (s *Scanner) AddBeacon(dists [3]int) {
	s.b[dists] = &Beacon{s: s, d: dists, r: make(map[[3]int]bool)}
}

func (s *Scanner) Rotate(rot []int) *Scanner {
	ns := &Scanner{id: s.id, b: make(Beacons)}
	for _, r := range s.b {
		rt := [3]int{
			r.d[0] * rot[0],
			r.d[1] * rot[1],
			r.d[2] * rot[2],
		}
		dists := [3]int{
			rt[rot[3]],
			rt[rot[4]],
			rt[rot[5]],
		}
		ns.AddBeacon(dists)
	}
	ns.BuildRelations()
	return ns
}

func (s *Scanner) FindOverlap(s2 *Scanner) bool {
	for i := range rotations {
		ns := s2.rotations[i]
		for _, b1 := range s.b {
			for _, b2 := range ns.b {
				unmatched := [][3]int{}
				matched := 0
				for br := range b2.r {
					if _, f := b1.r[br]; f {
						matched++
					} else {
						unmatched = append(unmatched, br)
					}
				}
				if matched >= 11 {
					for _, d := range unmatched {
						dists := [3]int{
							b1.d[0] - d[0],
							b1.d[1] - d[1],
							b1.d[2] - d[2],
						}
						s.AddBeacon(dists)
					}
					s.BuildRelations()
					scannerDist[s2.id] = [3]int{
						b1.d[0] - b2.d[0],
						b1.d[1] - b2.d[1],
						b1.d[2] - b2.d[2],
					}
					return true
				}
				break
			}
		}
	}
	return false
}

func ReadData(filename string) []*Scanner {
	var s *Scanner
	sid := 0
	scanners := []*Scanner{}
	for _, l := range lib.SlurpFile(filename) {
		if l == "" {
			s.BuildRelations()
		} else if l[:3] == "---" {
			s = &Scanner{id: sid, rotations: make(map[int]*Scanner), b: make(Beacons)}
			scanners = append(scanners, s)
			sid++
		} else {
			dists := [3]int{}
			for i, ds := range strings.Split(l, ",") {
				dists[i] = lib.StrToInt(ds)
			}
			s.AddBeacon(dists)
		}
	}
	for _, sc := range scanners {
		for i, r := range rotations {
			sc.rotations[i] = sc.Rotate(r)
		}
	}
	s.BuildRelations()
	return scanners
}

func FindOverlaps(scanners []*Scanner) *Scanner {
	s1 := scanners[0]
	toCheck := make([]*Scanner, len(scanners)-1)
	copy(toCheck, scanners[1:])
	for len(toCheck) > 0 {
		s2 := toCheck[0]
		toCheck = toCheck[1:]
		if !s1.FindOverlap(s2) {
			toCheck = append(toCheck, s2)
		}
	}
	return s1

}

func PartA(filename string) int {
	scannerDist = map[int][3]int{0: {0, 0, 0}}
	scanners := ReadData(filename)
	s := FindOverlaps(scanners)
	return len(s.b)
}

func PartB(filename string) int {
	scannerDist = map[int][3]int{0: {0, 0, 0}}
	scanners := ReadData(filename)
	FindOverlaps(scanners)
	maxDist := 0
	for _, s1 := range scannerDist {
		for _, s2 := range scannerDist {
			d := int(math.Abs(float64(s1[0]-s2[0])) +
				math.Abs(float64(s1[1]-s2[1])) +
				math.Abs(float64(s1[2]-s2[2])))
			if d > maxDist {
				maxDist = d
			}
		}
	}
	return maxDist
}
