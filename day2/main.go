package day2

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

type Shape int

const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
	Won      int   = 6
	Draw     int   = 3
	Lost     int   = 0
)

var stringToShape = map[byte]Shape{
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
	'A': Rock,
	'B': Paper,
	'C': Scissors,
}

var loosingShape = map[Shape]Shape{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var winningShape = map[Shape]Shape{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

func ScoreRound(opponentShape Shape, myShape Shape) int {
	if opponentShape == myShape {
		return Draw
	}

	if winningShape[opponentShape] == myShape {
		return Won
	}

	return Lost
}

func Day2() int {
	var total = 0

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		var opponentShape = stringToShape[line[0]]
		var myShape = stringToShape[line[2]]
		total += ScoreRound(opponentShape, myShape) + int(myShape)
	}

	return total
}

func Day2Part2() int {
	var total = 0

	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		var opponentShape = stringToShape[line[0]]
		var result = line[2]
		if result == 'X' {
			total += (Lost + int(loosingShape[opponentShape]))
		} else if result == 'Y' {
			total += (Draw + int(opponentShape))
		} else {
			total += (Won + int(winningShape[opponentShape]))
		}
	}

	return total
}
