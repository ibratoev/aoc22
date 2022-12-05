package day5

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var cratesRegex = regexp.MustCompile(`^(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})\s(?:\[([A-Z])\]|\s{3})$`)
var moveRegex = regexp.MustCompile(`^move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)$`)

func atoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func Day5Part1() string {
	stacks := make([][]string, 9)
	for i := range stacks {
		stacks[i] = make([]string, 0, 50)
	}

	var parseCrates = true
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		if line == "" {
			parseCrates = false
			continue
		}
		if parseCrates {
			for i, crate := range cratesRegex.FindStringSubmatch(line) {
				if i > 0 && crate != "" {
					stacks[i-1] = append([]string{crate}, stacks[i-1]...)
				}
			}
		} else {
			move := moveRegex.FindStringSubmatch(line)
			count := atoi(move[1])
			from := atoi(move[2]) - 1
			to := atoi(move[3]) - 1
			for i := 0; i < count; i++ {
				crate := stacks[from][len(stacks[from])-1]
				stacks[from] = stacks[from][:len(stacks[from])-1]
				stacks[to] = append(stacks[to], crate)
			}
		}

	}

	var result = ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}

	return result
}

func Day5Part2() string {
	stacks := make([][]string, 9)
	for i := range stacks {
		stacks[i] = make([]string, 0, 50)
	}

	var parseCrates = true
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		if line == "" {
			parseCrates = false
			continue
		}
		if parseCrates {
			for i, crate := range cratesRegex.FindStringSubmatch(line) {
				if i > 0 && crate != "" {
					stacks[i-1] = append([]string{crate}, stacks[i-1]...)
				}
			}
		} else {
			move := moveRegex.FindStringSubmatch(line)
			count := atoi(move[1])
			from := atoi(move[2]) - 1
			to := atoi(move[3]) - 1

			crates := stacks[from][len(stacks[from])-count:]
			stacks[from] = stacks[from][:len(stacks[from])-count]
			stacks[to] = append(stacks[to], crates...)
		}

	}

	var result = ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}

	return result
}
