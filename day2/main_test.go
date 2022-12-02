package day2

import "testing"

func TestScoreRound(t *testing.T) {
	var data = []struct {
		myShape       Shape
		opponentShape Shape
		expectedScore int
	}{
		{Rock, Rock, 3},
		{Rock, Scissors, 6},
		{Paper, Scissors, 0},
	}

	for _, tt := range data {
		score := ScoreRound(tt.opponentShape, tt.myShape)
		if score != tt.expectedScore {
			t.Errorf("Unexpected score: %d,", tt)
		}
	}
}

func TestDay2(t *testing.T) {
	got := Day2()
	if got != 12586 {
		t.Errorf("Unexpected result: %d", got)
	}
}

func TestDay2Part2(t *testing.T) {
	got := Day2Part2()
	if got != 13193 {
		t.Errorf("Unexpected result: %d", got)
	}
}
