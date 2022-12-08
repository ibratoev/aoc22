package day5

import "testing"

func TestDay6Part1(t *testing.T) {
	got := Day6Part1()
	if got != 1582 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay6Part2(t *testing.T) {
	got := Day6Part2()
	if got != 3588 {
		t.Errorf("Unexpected result: %v", got)
	}
}
