package day03

import "testing"

func TestPartA(t *testing.T) {
	dist := PartA("data/test")
	if dist != 198 {
		t.Fatalf("Position should be 198! Got %d", dist)
	}
	t.Log(PartA("data/input"))
}

func TestB(t *testing.T) {
	dist := PartB("data/test")
	if dist != 230 {
		t.Fatalf("Position should be 230! Got %d", dist)
	}
	t.Log(PartB("data/input"))
}
