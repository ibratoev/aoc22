package day8

import "testing"

func TestDay8Part1(t *testing.T) {
	got := Day8Part1()
	if got != 1150 {
		t.Errorf("Unexpected result: %v", got)
	}
}
