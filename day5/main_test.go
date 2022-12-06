package day5

import "testing"

func TestDay5Part1(t *testing.T) {
	got := Day5Part1()
	if got != "DHBJQJCCW" {
		t.Errorf("Unexpected result: %v", got)
	}
}

func TestDay5Part2(t *testing.T) {
	got := Day5Part2()
	if got != "WJVRLSJJT" {
		t.Errorf("Unexpected result: %v", got)
	}
}
