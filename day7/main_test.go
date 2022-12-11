package day6

import "testing"

func TestDay7Part1(t *testing.T) {
	got := Day7Part1()
	if got != 1182909 {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay7Part2(t *testing.T) {
	got := Day7Part2()
	if got != 2832508 {
		t.Errorf("Unexpected result: %v", got)
	}
}
