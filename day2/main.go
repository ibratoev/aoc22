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
)

var stringToShape = map[byte]Shape{
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
	'A': Rock,
	'B': Paper,
	'C': Scissors,
}

func ScoreRound(opponentShape Shape, myShape Shape) int {
	if opponentShape == myShape {
		return 3
	}

	if myShape == Rock && opponentShape == Scissors {
		return 6
	}

	if myShape == Paper && opponentShape == Rock {
		return 6
	}

	if myShape == Scissors && opponentShape == Paper {
		return 6
	}

	return 0
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
