package day2

import "testing"

func TestPart1(t *testing.T) {
	answer, err := solvePart1("./input1.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(answer)
}

func TestPart2(t *testing.T) {
	answer, err := solvePart2("./input1.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(answer)
}
