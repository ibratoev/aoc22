package day3

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func getItemPriority(x rune) int {
	if x >= 'a' && x <= 'z' {
		return int(x) - 'a' + 1
	}
	return int(x) - 'A' + 27
}

func getDuplicateItem(items string) rune {
	var length = len([]rune(items))
	var types = make(map[rune]bool)

	for pos, char := range items {
		if pos < length/2 {
			types[char] = true
		} else {
			if types[char] {
				return char
			}
		}
	}

	panic("unexpected")
}

func Day3() int {
	var sum = 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		var item = getDuplicateItem(line)
		sum += getItemPriority(item)
	}

	return sum
}

func findBadge(group [3]string) rune {
	var badges = make(map[rune]int)
	for _, rucksack := range group {
		var types = make(map[rune]bool)
		for _, item := range rucksack {
			if !types[item] {
				badges[item]++
			}
			types[item] = true
		}
	}

	for key, value := range badges {
		if value == 3 {
			return key
		}
	}

	panic("unexpected")
}

func Day3Part2() int {
	var sum = 0
	var group [3]string
	for pos, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		group[pos%3] = line
		if pos%3 == 2 {
			sum += getItemPriority(findBadge(group))
		}
	}

	return sum
}
