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

	for pos, item := range items {
		if pos < length/2 {
			types[item] = true
		} else {
			if types[item] {
				return item
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

func findBadge(group []string) rune {
	var badges = make(map[rune]int)
	for _, rucksack := range group {
		var types = make(map[rune]bool)
		for _, item := range rucksack {
			if !types[item] {
				badges[item]++
				if badges[item] == 3 {
					return item
				}
			}
			types[item] = true
		}
	}

	panic("unexpected")
}

func Day3Part2() int {
	var sum = 0
	var lines = strings.Split(strings.TrimRight(input, "\n"), "\n")
	for i := 3; i <= len(lines); i += 3 {
		sum += getItemPriority(findBadge(lines[i-3 : i]))
	}

	return sum
}
