package day8

import "testing"

func BenchmarkDay8Part1(t *testing.B) {
	got := Day8Part1()
	if got != 58681 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func BenchmarkDay8Part2(t *testing.B) {
	got := Day8Part2()
	if got != 3632850 {
		t.Errorf("Unexpected result: %v", got)
	}
}
