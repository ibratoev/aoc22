package day1

import (
	"testing"
)

func TestDay1(t *testing.T) {
	got := Day1()
	if got != 71502 {
		t.Errorf("Unexpected result: %d", got)
	}
}

func TestDay1Part2(t *testing.T) {
	got := Day1Part2()
	if got != 208191 {
		t.Errorf("Unexpected result: %d", got)
	}
}
