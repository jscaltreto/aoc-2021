package day02

import "testing"

func TestA(t *testing.T) {
	dist := PartA("data/test")
	if dist != 150 {
		t.Fatalf("Position should be 150! Got %d", dist)
	}
	t.Log(PartA("data/input"))
}

func TestB(t *testing.T) {
	dist := PartB("data/test")
	if dist != 900 {
		t.Fatalf("Position should be 900! Got %d", dist)
	}
	t.Log(PartB("data/input"))
}
