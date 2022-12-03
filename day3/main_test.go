package day3

import "testing"

func TestGetItemPriority(t *testing.T) {
	var data = []struct {
		itemType rune
		priority int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}

	for _, tt := range data {
		priority := getItemPriority(tt.itemType)
		if priority != tt.priority {
			t.Errorf("Unexpected priority: %d,", tt)
		}
	}
}

func TestGetDuplicateItem(t *testing.T) {
	var data = []struct {
		items    string
		expected rune
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p'},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L'},
		{"PmmdzqPrVvPwwTWBwg", 'P'},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v'},
		{"ttgJtRGJQctTZtZT", 't'},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 's'},
	}

	for _, tt := range data {
		item := getDuplicateItem(tt.items)
		if item != tt.expected {
			t.Errorf("Unexpected item: %+v,", tt)
		}
	}
}

func TestDay3(t *testing.T) {
	got := Day3()
	if got != 8139 {
		t.Errorf("Unexpected result: %d", got)
	}
}
