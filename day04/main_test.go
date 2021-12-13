package day04

import "testing"

func TestPartA(t *testing.T) {
	dist := PartA("data/test")
	if dist != 4512 {
		t.Fatalf("Position should be 4512! Got %d", dist)
	}
	t.Log(PartA("data/input"))
}

func TestB(t *testing.T) {
	dist := PartB("data/test")
	if dist != 1924 {
		t.Fatalf("Position should be 1924! Got %d", dist)
	}
	t.Log(PartB("data/input"))
}
