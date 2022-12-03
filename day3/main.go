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
			if types[char] == true {
				return char
			}
		}
	}

	return ' '
}

func Day3() int {
	var sum = 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		var item = getDuplicateItem(line)
		sum += getItemPriority(item)
	}

	return sum
}
