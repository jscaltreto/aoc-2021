package day11

import (
	"github.com/jscaltreto/aoc-2021/lib"
)

type Octopus struct {
	x int
	y int
	power int
	flashed bool
	octopuses *Octopuses
}

func (o *Octopus) Neighbors() []*Octopus {
	ops := []*Octopus{}
	for yi:= o.y -1; yi <= o.y+1; yi++ {
		if yi < 0 || yi >= len(o.octopuses.os) {
			continue
		}
		for xi := o.x - 1; xi <= o.x + 1; xi++ {
			if xi < 0 || xi >= len(o.octopuses.os[yi]) {
				continue
			}
			ops = append(ops, o.octopuses.os[yi][xi])
		}
	}
	return ops
}

func (o *Octopus) PowerUp(ch chan *Octopus) {
	o.power++
	if o.power == 10 {
		ch <- o
		for _, n := range o.Neighbors() {
			n.PowerUp(ch)
		}
	}
}

func (o *Octopus) Flash() {
	o.octopuses.flashes++
	o.power = 0
}

type Octopuses struct {
	os [10][10]*Octopus
	flashes int
	allFlashed bool
}

func (o *Octopuses) Step() {
	flashchan := make(chan *Octopus, 100)
	for _, os := range o.os {
		for _, o := range os {
			o.PowerUp(flashchan)
		}
	}
	close(flashchan)
	if len(flashchan) == 100 {
		o.allFlashed = true
	}
	for flasher := range flashchan {
		flasher.Flash()
	}
}

func Load(filename string) *Octopuses {
	octopuses := &Octopuses{}
	for y, l := range lib.SlurpFile(filename) {
		for x, o := range l {
			octopuses.os[y][x] = &Octopus{
				x:x,
				y:y,
				power:lib.StrToInt(string(o)),
				octopuses: octopuses,
			}
		}
	}
	return octopuses
}

func PartA(filename string) int {
	octopuses := Load(filename)
	for i := 0; i < 100; i++ {
		octopuses.Step()
	}
	return octopuses.flashes
}

func PartB(filename string) int {
	octopuses := Load(filename)
	step := 0
	for {
		step++
		octopuses.Step()
		if octopuses.allFlashed {
			return step
		}
	}
}
