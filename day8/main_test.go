package day8

import "testing"

func TestDay8Part1(t *testing.T) {
	got := Day8Part1()
	if got != 1832 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay8Part2(t *testing.T) {
	got := Day8Part2()
	if got != 2401 {
		t.Errorf("Unexpected result: %v", got)
	}
}
