package day15

import (
	"container/heap"
	"math"

	"github.com/jscaltreto/aoc-2021/lib"
)

type Node struct {
	x, y, w, c, i int
	v             bool
}

func (n *Node) Id() [2]int {
	return [2]int{n.x, n.y}
}

func (n *Node) Update(baseCost int) bool {
	g := baseCost + n.w
	if g < n.c {
		n.c = g
		return true
	}
	return false
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].c < pq[j].c
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].i = i
	pq[j].i = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.i = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.i = -1
	*pq = old[0 : n-1]
	return item
}

type Map [][]*Node

var nx [4]int = [4]int{0, 0, -1, 1}
var ny [4]int = [4]int{-1, 1, 0, 0}

func (m Map) Neighbors(n *Node) []*Node {
	ns := []*Node{}
	for i, ox := range nx {
		x := n.x + ox
		y := n.y + ny[i]
		if !(x < 0 || x >= len(m) || y < 0 || y >= len(m)) {
			if n := m[y][x]; !n.v {
				ns = append(ns, n)
			}
		}
	}
	return ns
}

func (m Map) FindPath() int {
	root := m[0][0]
	root.c = 0
	pq := PQ{root}
	heap.Init(&pq)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if node.x+1 >= len(m[0]) && node.y+1 >= len(m) {
			return node.c
		}
		for _, n := range m.Neighbors(node) {
			if updated := n.Update(node.c); updated {
				heap.Push(&pq, n)
			}
		}
		node.v = true
	}
	return 0
}

func LoadMap(filename string, extend int) Map {
	data := lib.SlurpFile(filename)
	idx := 0
	sy, sx := len(data), len(data[0])
	m := make(Map, sy*extend)
	for iy := 0; iy < sy*extend; iy++ {
		m[iy] = make([]*Node, sx*extend)
		for ix := 0; ix < sx*extend; ix++ {
			w := ((lib.StrToInt(string(data[iy%sy][ix%sx])) + (iy / sy) + (ix / sx) - 1) % 9) + 1
			node := &Node{
				x: ix,
				y: iy,
				w: w,
				c: math.MaxInt,
				i: idx,
			}
			m[iy][ix] = node
			idx++
		}
	}
	return m
}

func PartA(filename string) int {
	m := LoadMap(filename, 1)
	p := m.FindPath()
	return p
}

func PartB(filename string) int {
	m := LoadMap(filename, 5)
	p := m.FindPath()
	return p
}
