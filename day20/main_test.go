package day20

import "testing"

const (
	ExpA = 35
	ExpB = 3351
)

func TestPartA(t *testing.T) {
	ans := PartA("data/test")
	if ans != ExpA {
		t.Fatalf("Answer should be %d! Got %d", ExpA, ans)
	}
	t.Log(PartA("data/input"))
}

func TestPartB(t *testing.T) {
	ans := PartB("data/test")
	if ans != ExpB {
		t.Fatalf("Answer should be %d! Got %d", ExpB, ans)
	}
	t.Log(PartB("data/input"))
}
