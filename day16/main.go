package day16

import (
	"math"
	"strconv"

	"github.com/jscaltreto/aoc-2021/lib"
)

const (
	EndBitMask uint64 = 0b10000
	DataMask   uint64 = 0b1111
)

type Packet struct {
	data       string
	version    int
	value      uint64
	cur        int
}

func (p *Packet) Literal() uint64 {
	var result uint64
	for {
		g := p.ReadBits(5)
		result = (result << 4) | g&DataMask
		if g&EndBitMask == 0 {
			return result
		}
	}
}

func (p *Packet) ReadBits(n int) uint64 {
	sbyte := p.cur / 4
	p.cur += n
	ebyte := (p.cur - 1) / 4
	bytes, _ := strconv.ParseUint(p.data[sbyte:ebyte+1], 16, 64)

	// Strip trailing bytes
	if s := (p.cur % 4); s != 0 {
		bytes >>= 4 - s
	}

	// Strip leading bytes
	shift := 64 - n
	bytes = (bytes << shift) >> shift
	return bytes
}

func (p *Packet) Read() {
	p.version = int(p.ReadBits(3))
	packettype := byte(p.ReadBits(3))
	if packettype == 4 {
		p.value = p.Literal()
	} else {
		sps := []*Packet{}
		if p.ReadBits(1) == 1 {
			subPackets := p.ReadBits(11)
			for i := 0; i < int(subPackets); i++ {
				sps = append(sps, p.ReadSubPacket())
			}
		} else {
			t := int(p.ReadBits(15))
			s := p.cur
			for p.cur < t+s {
				sps = append(sps, p.ReadSubPacket())
			}
		}
		switch packettype {
		case 0:
			for _, sp := range sps {
				p.value += sp.value
			}
		case 1:
			p.value = 1
			for _, sp := range sps {
				p.value *= sp.value
			}
		case 2:
			p.value = math.MaxUint64
			for _, sp := range sps {
				if sp.value < p.value {
					p.value = sp.value
				}
			}
		case 3:
			for _, sp := range sps {
				if sp.value > p.value {
					p.value = sp.value
				}
			}
		case 5:
			if sps[0].value > sps[1].value {
				p.value = 1
			}
		case 6:
			if sps[0].value < sps[1].value {
				p.value = 1
			}
		case 7:
			if sps[0].value == sps[1].value {
				p.value = 1
			}
		}
	}
}

func (p *Packet) ReadSubPacket() *Packet {
	var leadIn, pre int
	rem := p.data[p.cur/4:]

	// If the packet doesn't start on a 4-bit boundary,
	// strip the leading bits.
	if pre = p.cur % 4; pre != 0 {
		leadIn = 4 - pre
		rem = strconv.FormatInt(int64(p.ReadBits(leadIn)), 16) + rem[1:]
		p.cur -= leadIn
	}
	sp := ReadPacket(rem, pre)
	p.cur += sp.cur - pre
	p.version += sp.version
	return sp
}

func ReadPacket(packet string, cur int) *Packet {
	p := &Packet{
		data: packet,
		cur:  cur,
	}
	p.Read()
	return p
}

func PartA(filename string) int {
	data := lib.SlurpFile(filename)[0]
	p := ReadPacket(data, 0)
	return p.version
}

func PartB(filename string) int {
	data := lib.SlurpFile(filename)[0]
	p := ReadPacket(data, 0)
	return int(p.value)
}
