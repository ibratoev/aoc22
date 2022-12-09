package day9

import "testing"

func TestDay9Part1(t *testing.T) {
	got := Day8Part1()
	if got != 5883 {
		t.Errorf("Unexpected result: %v", got)
	}
}
