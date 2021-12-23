package day22

import (
	"regexp"

	"github.com/jscaltreto/aoc-2021/lib"
)

const re = `^(on|off) x=([-\d]+)..([-\d]+),y=([-\d]+)..([-\d]+),z=([-\d]+)..([-\d]+)$`

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

type Block struct {
	dims [2][3]int
	on   bool
}

func NewBlock(dims [2][3]int, on bool) *Block {
	return &Block{dims: dims, on: on}
}

func (b *Block) Vol() int {
	return (b.dims[1][0] - b.dims[0][0] + 1) * (b.dims[1][1] - b.dims[0][1] + 1) * (b.dims[1][2] - b.dims[0][2] + 1)
}

func (b *Block) Intersect(new *Block) *Block {
	i := [2][3]int{
		{Max(b.dims[0][0], new.dims[0][0]), Max(b.dims[0][1], new.dims[0][1]), Max(b.dims[0][2], new.dims[0][2])},
		{Min(b.dims[1][0], new.dims[1][0]), Min(b.dims[1][1], new.dims[1][1]), Min(b.dims[1][2], new.dims[1][2])},
	}
	if i[0][0] > i[1][0] || i[0][1] > i[1][1] || i[0][2] > i[1][2] {
		return nil
	}
	// If we intersect an existing "off" block, we need the intersection to be
	// "on" (and vice versa) to prevent double-dipping
	return NewBlock(i, !b.on)
}

func LoadBlocks(filename string, limit int) int {
	r := regexp.MustCompile(re)
	blocks := []*Block{}
	for _, ins := range lib.SlurpFile(filename) {
		if res := r.FindStringSubmatch(ins); len(res) > 0 {
			nb := NewBlock([2][3]int{
				{lib.StrToInt(res[2]), lib.StrToInt(res[4]), lib.StrToInt(res[6])},
				{lib.StrToInt(res[3]), lib.StrToInt(res[5]), lib.StrToInt(res[7])},
			}, res[1] == "on")
			if limit > 0 && (nb.dims[0][0] < -limit || nb.dims[0][0] > limit ||
				nb.dims[1][0] < -limit || nb.dims[1][0] > limit ||
				nb.dims[0][1] < -limit || nb.dims[0][1] > limit ||
				nb.dims[1][1] < -limit || nb.dims[1][1] > limit ||
				nb.dims[0][2] < -limit || nb.dims[0][2] > limit ||
				nb.dims[1][2] < -limit || nb.dims[1][2] > limit) {
				continue
			}
			for _, b := range blocks {
				if ib := b.Intersect(nb); ib != nil {
					blocks = append(blocks, ib)
				}
			}
			if nb.on {
				blocks = append(blocks, nb)
			}
		}
	}

	t := 0
	for _, b := range blocks {
		if b.on {
			t += b.Vol()
		} else {
			t -= b.Vol()
		}
	}
	return t
}

func PartA(filename string) int {
	return LoadBlocks(filename, 50)
}

func PartB(filename string) int {
	return LoadBlocks(filename, 0)
}
