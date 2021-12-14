package day13

import "testing"

const (
	ExpA = 17
	ExpB = "\n#####\n#   #\n#   #\n#   #\n#####\n"
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
		t.Fatalf("Answer should be %s! Got %s", ExpB, ans)
	}
	t.Log(PartB("data/input"))
}
