package day12

import (
	"strings"

	"github.com/jscaltreto/aoc-2021/lib"
)

type Node struct {
	cave     *Cave
	parent   *Node
	vistwice bool // True if we've visited any small caves twice
}

type Cave struct {
	id    string
	small bool
	conns []*Cave
}

type CaveMap map[string]*Cave

func (m CaveMap) Get(id string) *Cave {
	if c, f := m[id]; f {
		return c
	}
	c := &Cave{id: id, small: strings.ToLower(id) == id}
	m[id] = c
	return c
}

func MapCaves(filename string) *Node {
	cm := &CaveMap{}
	for _, connStr := range lib.SlurpFile(filename) {
		conns := strings.Split(connStr, "-")
		c1, c2 := cm.Get(conns[0]), cm.Get(conns[1])
		c1.conns = append(c1.conns, c2)
		c2.conns = append(c2.conns, c1)
	}
	return &Node{cave: cm.Get("start")}
}

func CanVisit(head *Node, partB bool) bool {
	cur := head
	if head.cave.small {
		for cur.parent != nil {
			cur = cur.parent
			if cur.cave.id == head.cave.id {
				if !partB || head.vistwice {
					return false
				}
				head.vistwice = true
				return true
			}
		}
	}
	return true
}

func Recurse(head *Node, partB bool) []*Node {
	paths := []*Node{}
	for _, next := range head.cave.conns {
		if next.id != "start" {
			nx := &Node{cave: next, parent: head, vistwice: head.vistwice}
			if next.id == "end" {
				paths = append(paths, nx)
			} else if CanVisit(nx, partB) {
				paths = append(paths, Recurse(nx, partB)...)
			}
		}
	}
	return paths
}

func PartA(filename string) int {
	return len(Recurse(MapCaves(filename), false))
}

func PartB(filename string) int {
	return len(Recurse(MapCaves(filename), true))
}
