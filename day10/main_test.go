package day10

import "testing"

func TestDay10Part1(t *testing.T) {
	got := Day10Part1()
	if got != 11820 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay10Part2(t *testing.T) {
	Day10Part2()
}
