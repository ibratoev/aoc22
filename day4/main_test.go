package day4

import "testing"

func TestParseRange(t *testing.T) {
	var data = []struct {
		rangeStr string
		expected Range
	}{
		{"2-4", Range{2, 4}},
		{"2-2", Range{2, 2}},
		{"0-6", Range{0, 6}},
	}

	for _, tt := range data {
		parsed := parseRange(tt.rangeStr)
		if parsed != tt.expected {
			t.Errorf("Unexpected range: %+v,", tt)
		}
	}
}

func TestRangeOverlap(t *testing.T) {
	var data = []struct {
		ranges   [2]Range
		overlaps bool
	}{
		{[2]Range{{5, 7}, {7, 9}}, true},
		{[2]Range{{2, 8}, {3, 7}}, true},
		{[2]Range{{2, 8}, {9, 10}}, false},
		{[2]Range{{4, 6}, {6, 6}}, true},
		{[2]Range{{2, 6}, {4, 8}}, true},
	}

	for _, tt := range data {
		overlaps := tt.ranges[0].Overlap(&tt.ranges[1])
		if overlaps != tt.overlaps {
			t.Errorf("Unexpected: %+v,", tt)
		}
	}
}

func TestDay4Part1(t *testing.T) {
	got := Day4Part1()
	if got != 651 {
		t.Errorf("Unexpected result: %d", got)
	}
}

func TestDay4Part2(t *testing.T) {
	got := Day4Part2()
	if got != 825 {
		t.Errorf("Unexpected result: %d", got)
	}
}
