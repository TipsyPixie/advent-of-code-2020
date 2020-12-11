package aoc

import "testing"

func TestFromFile(t *testing.T) {
	input, err := FromFile("./README.md")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer func() { _ = input.Close() }()

	for _, ok, err := input.ReadLine(); ok || err != nil; _, ok, err = input.ReadLine() {
		if err != nil {
			t.Errorf("ReadLine() failed: %s", err)
			t.FailNow()
		}
	}

	_, err = input.ReadAll()
	if err != nil {
		t.Errorf("ReadAll() failed: %s", err)
		t.FailNow()
	}
}
