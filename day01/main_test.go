package day01

import (
	"testing"
)

func TestPartA(t *testing.T) {
	increases := PartA("data/test")
	if increases != 7 {
		t.Fatalf("Increases should be 7! Got %d", increases)
	}
	t.Log(PartA("data/input"))
}

func TestPartB(t *testing.T) {
	increases := PartB("data/test")
	if increases != 5 {
		t.Fatalf("Increases should be 5! Got %d", increases)
	}
	t.Log(PartB("data/input"))
}
