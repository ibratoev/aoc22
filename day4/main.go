package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Range struct {
	start int
	end   int
}

func (r *Range) Contains(other *Range) bool {
	return r.start <= other.start && r.end >= other.end
}

func (r *Range) Overlaps(other *Range) bool {
	return r.Contains(&Range{other.start, other.start}) ||
		other.Contains(&Range{r.start, r.start})
}

func parseRange(rangeString string) Range {
	var parsed = strings.Split(rangeString, "-")
	if start, errStart := strconv.Atoi(parsed[0]); errStart == nil {
		if end, errEnd := strconv.Atoi(parsed[1]); errEnd == nil {
			return Range{start, end}
		}
	}

	panic(fmt.Sprintf("Error parsing range: %v", rangeString))
}

func parseRanges(line string) (Range, Range) {
	var ranges = strings.Split(line, ",")
	return parseRange(ranges[0]), parseRange(ranges[1])
}

func Day4Part1() int {
	var sum = 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		first, second := parseRanges(line)
		if first.Contains(&second) || second.Contains(&first) {
			sum++
		}
	}

	return sum
}

func Day4Part2() int {
	var sum = 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		first, second := parseRanges(line)
		if first.Overlaps(&second) {
			sum++
		}
	}

	return sum
}
