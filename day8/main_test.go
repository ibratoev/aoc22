package day8

import "testing"

func BenchmarkDay8Part1(t *testing.B) {
	Day8Part1()
}

func BenchmarkDay8Part2(t *testing.B) {
	Day8Part2()
}

func TestDay8Part1(t *testing.T) {
	got := Day8Part1()
	if got != 58681 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay8Part2(t *testing.T) {
	got := Day8Part2()
	if got != 3632850 {
		t.Errorf("Unexpected result: %v", got)
	}
}
