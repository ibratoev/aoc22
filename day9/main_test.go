package day9

import "testing"

func TestDay9Part1(t *testing.T) {
	got := Day9Part1()
	if got != 5883 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay9Part2(t *testing.T) {
	got := Day9Part2()
	if got != 2367 {
		t.Errorf("Unexpected result: %v", got)
	}
}
