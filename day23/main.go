package day23

import (
	"container/heap"
	"fmt"

	"github.com/jscaltreto/aoc-2021/lib"
)

var MoveCost map[Amph][2]int = map[Amph][2]int{
	'A': {1, 0},
	'B': {10, 1},
	'C': {100, 2},
	'D': {1000, 3},
}

type Amph byte

func (t Amph) Cost() int { return MoveCost[t][0] }
func (t Amph) TargetRoom() int { return MoveCost[t][1] }

var X Amph = Amph('x')
var Clear [11]*Amph = [11]*Amph{nil, nil, &X, nil, &X, nil, &X, nil, &X, nil, nil}

type Room struct {
	affinity Amph
	slots    []*Amph
	hidx     int
}

func (r *Room) Top() int {
	for i, a := range r.slots {
		if a != nil {
			return i
		}
	}
	return len(r.slots)
}

func (r *Room) HasUnwanted() bool {
	for _, a := range r.slots {
		if a != nil && *a != r.affinity {
			return true
		}
	}
	return false
}

func (r *Room) CanEnter() bool {
	return !r.HasUnwanted() && r.Top() > 0
}

type Move interface {
	Execute()
	SetIndex(int)
	GetSub() *Sub
	H() int
}

type RoomToHall struct {
	sub    *Sub
	roomid int
	slotid int
	hallid int
	i      int
}

func (m *RoomToHall) Execute() {
	ns := m.sub.Clone()
	room := ns.rooms[m.roomid]
	a := room.slots[m.slotid]
	ns.energy += (1 + m.slotid + lib.Abs(m.hallid-room.hidx)) * a.Cost()
	ns.hall[m.hallid] = a
	room.slots[m.slotid] = nil
	m.sub = ns
}

func (m *RoomToHall) H() int         { return m.sub.energy + (m.sub.disorder * 2) }
func (m *RoomToHall) SetIndex(i int) { m.i = i }
func (m *RoomToHall) GetSub() *Sub   { return m.sub }

type HallToRoom struct {
	sub    *Sub
	roomid int
	hallid int
	i      int
}

func (m *HallToRoom) Execute() {
	ns := m.sub.Clone()
	room := ns.rooms[m.roomid]
	a := ns.hall[m.hallid]
	slotid := room.Top() -1
	ns.disorder -= a.Cost()
	ns.energy += (1 + slotid + lib.Abs(room.hidx-m.hallid)) * a.Cost()
	room.slots[slotid] = a
	ns.hall[m.hallid] = nil
	m.sub = ns
}

func (m *HallToRoom) H() int         { return m.sub.energy + (m.sub.disorder * 2) }
func (m *HallToRoom) SetIndex(i int) { m.i = i }
func (m *HallToRoom) GetSub() *Sub   { return m.sub }

type Sub struct {
	hall     [11]*Amph
	rooms    [4]*Room
	roomsize int
	disorder int
	energy   int
}

func NewSub(size int) *Sub {
	sub := &Sub{hall: [11]*Amph{}, rooms: [4]*Room{}, roomsize: size}

	for i := range sub.rooms {
		sub.rooms[i] = &Room{
			hidx:     2 + (2 * i),
			affinity: Amph('A' + i),
			slots:    make([]*Amph, size),
		}
	}
	return sub
}

func (s *Sub) Clone() *Sub {
	ns := NewSub(s.roomsize)
	copy(ns.hall[:], s.hall[:])
	for i, r := range s.rooms {
		copy(ns.rooms[i].slots[:], r.slots[:])
	}
	ns.energy = s.energy
	ns.disorder = s.disorder
	return ns
}

func (s *Sub) Disorder() int {
	disorder := 0
	for ri, r := range s.rooms {
		for sid, s := range r.slots {
			if s != nil && s.TargetRoom() != ri {
				disorder += s.Cost() * (sid + 1)
			}
		}
	}
	return disorder
}

func (s *Sub) HallAvail(idx int) []int {
	moves := []int{}
	for i := idx - 1; i >= 0; i-- {
		if s.hall[i] == &X {
			continue
		} else if s.hall[i] != nil {
			break
		}
		moves = append(moves, i)
	}
	for i := idx + 1; i < 11; i++ {
		if s.hall[i] == &X {
			continue
		} else if s.hall[i] != nil {
			break
		}
		moves = append(moves, i)
	}
	return moves
}

func (s *Sub) CanReach(i int, r *Room) bool {
	for hi := i + 1; hi < r.hidx; hi++ {
		if s.hall[hi] != nil && s.hall[hi] != &X {
			return false
		}
	}
	for hi := i - 1; hi > r.hidx; hi-- {
		if s.hall[hi] != nil && s.hall[hi] != &X {
			return false
		}
	}
	return true
}

func (s *Sub) FindValidMoves() []interface{} {
	var valid []interface{}
	idx := 0
	for ri, r := range s.rooms {
		if r.HasUnwanted() {
			si := r.Top()
			if si == len(r.slots) {
				continue
			}
			for _, hi := range s.HallAvail(r.hidx) {
				move := &RoomToHall{
					sub:    s,
					roomid: ri,
					slotid: si,
					hallid: hi,
					i:      idx,
				}
				move.Execute()
				valid = append(valid, move)
				idx++
			}
		}
	}
	for hi, a := range s.hall {
		if a != nil && a != &X {
			trid := a.TargetRoom()
			tr := s.rooms[trid]
			if tr.CanEnter() && s.CanReach(hi, tr) {
				move := &HallToRoom{
					sub:    s,
					hallid: hi,
					roomid: trid,
					i:      idx,
				}
				move.Execute()
				valid = append(valid, move)
				idx++
			}

		}
	}

	return valid
}

func (s *Sub) Print() {
	fmt.Print("#############\n#")
	for _, h := range s.hall {
		if h == nil || h == &X {
			fmt.Print(".")
		} else {
			fmt.Print(string(*h))
		}
	}
	fmt.Print("#\n")
	for i := range s.rooms[0].slots {
		fmt.Print("  #")
		for ri := range s.rooms {
			if s.rooms[ri].slots[i] == nil {
				fmt.Print(".#")
			} else {
				fmt.Print(string(*s.rooms[ri].slots[i]) + "#")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("  #########")
}

func LoadInput(data []string) int {
	sub := NewSub(len(data) - 3)
	sub.hall[2], sub.hall[4], sub.hall[6], sub.hall[8] = &X, &X, &X, &X
	for i, l := range data[2 : len(data)-1] {
		a0, a1, a2, a3 := Amph(l[3]), Amph(l[5]), Amph(l[7]), Amph(l[9])
		sub.rooms[0].slots[i] = &a0
		sub.rooms[1].slots[i] = &a1
		sub.rooms[2].slots[i] = &a2
		sub.rooms[3].slots[i] = &a3
	}
	sub.disorder = sub.Disorder()

	moves := sub.FindValidMoves()
	pq := PQ(moves)
	heap.Init(&pq)

	memo := make(map[[11]*Amph]int)
	for pq.Len() > 0 {
		move := heap.Pop(&pq).(Move)
		result := move.GetSub()
		nextMoves := result.FindValidMoves()
		if len(nextMoves) == 0 && result.hall == Clear {
			return result.energy
		}
		for _, m := range nextMoves {
			ms := m.(Move).GetSub()
			if s, f := memo[ms.hall]; f && s == ms.energy {
				continue
			}
			memo[ms.hall] = ms.energy
			heap.Push(&pq, m)
		}
	}

	fmt.Println(len(moves))
	return 0

}

func PartA(filename string) int {
	data := lib.SlurpFile(filename)
	return LoadInput(data)
}

func PartB(filename string) int {
	extra := []string{
		"  #D#C#B#A#",
		"  #D#B#A#C#",
	}
	data := lib.SlurpFile(filename)
	nd := make([]string, 3)
	copy(nd, data[:3])
	nd = append(nd, extra...)
	nd = append(nd, data[3:]...)
	return LoadInput(nd)
}
