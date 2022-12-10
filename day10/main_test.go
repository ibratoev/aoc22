package day10

import "testing"

func TestDay9Part1(t *testing.T) {
	got := Day10Part1()
	if got != 11820 {
		t.Errorf("Unexpected result: %v", got)
	}
}
