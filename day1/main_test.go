package day1

import (
	"testing"
)

func TestDay1(t *testing.T) {
	got := Day1()
	if got == 0 {
		t.Errorf("Unexpected result: %d", got)
	}
}
