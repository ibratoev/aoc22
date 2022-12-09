package day8

import (
	"embed"
	_ "embed"
	"io/fs"
)

//go:embed input5000.txt
var f embed.FS

const LINES = 5000
const COLUMNS = 5000

func readFile(file fs.File, a []byte) int {
	i, err := file.Read(a)
	if err != nil {
		panic("Error reading from file")
	}
	return i
}

func Day8Part1() int {
	file, _ := f.Open("input5000.txt")
	defer file.Close()

	var trees [LINES][COLUMNS]byte

	visible := make(map[int]bool)
	markVisible := func(l int, c int) {
		visible[l*10000+c] = true
	}

	var maxColumns [COLUMNS]byte

	for l := 0; l < LINES; l++ {
		line := trees[l][:]
		readFile(file, line)
		var maxColumn byte = 0
		for c := 0; c < COLUMNS; c++ {
			// CHECK line ->
			if line[c] > maxColumn {
				maxColumn = line[c]
				markVisible(l, c)
			}

			// CHECK column ->
			if line[c] > maxColumns[c] {
				maxColumns[c] = line[c]
				markVisible(l, c)
			}
		}

		var maxX2 byte = 0
		// CHECK  line <-
		for c2 := COLUMNS - 1; c2 > -1; c2-- {
			if line[c2] > maxX2 {
				maxX2 = line[c2]
				markVisible(l, c2)
			}
			if maxX2 == '9' {
				break
			}
		}

		if l != COLUMNS-1 {
			// READ /n to remove it from the stream
			var newLine [1]byte
			readFile(file, newLine[:])
			if newLine[0] != '\n' {
				panic("EXPECTED NEW LINE!")
			}
		}
	}

	for c := 0; c < COLUMNS; c++ {
		var maxC byte = 0
		// CHECK column <-
		for l := LINES - 1; l > -1; l-- {
			if trees[l][c] > maxC {
				maxC = trees[l][c]
				markVisible(l, c)
			}
			if maxC == '9' {
				break
			}
		}
	}

	return len(visible)
}

func Day8Part2() int {
	file, _ := f.Open("input5000.txt")
	defer file.Close()

	var trees [LINES][COLUMNS]byte

	scores := make(map[int]int)
	setScore := func(l int, c int, score int) {
		val, ok := scores[l*10000+c]
		if ok {
			scores[l*10000+c] = val * score
		} else {
			scores[l*10000+c] = score
		}
	}

	for l := 0; l < LINES; l++ {
		line := trees[l][:]
		readFile(file, line)

		for c := 0; c < COLUMNS; c++ {
			// COUNT line right score ->
			var rightScore = 0
			for cRight := c + 1; cRight < COLUMNS; cRight++ {
				rightScore++
				if line[c] <= line[cRight] {
					break
				}
			}

			// COUNT line left score <-
			var leftScore = 0
			for cLeft := c - 1; cLeft > -1; cLeft-- {
				leftScore++
				if line[c] <= line[cLeft] {
					break
				}
			}

			setScore(l, c, leftScore*rightScore)
		}

		if l != LINES-1 {
			// READ /n to remove it from the stream
			var newLine [1]byte
			readFile(file, newLine[:])
			if newLine[0] != '\n' {
				panic("EXPECTED NEW LINE!")
			}
		}
	}

	for c := 0; c < COLUMNS; c++ {
		for l := 0; l < LINES; l++ {
			// COUNT column down score <-
			var downScore = 0
			for lDown := l + 1; lDown < LINES; lDown++ {
				downScore++
				if trees[l][c] <= trees[lDown][c] {
					break
				}
			}

			// COUNT column up score ->
			var upScore = 0
			for lUp := l - 1; lUp > -1; lUp-- {
				upScore++
				if trees[l][c] <= trees[lUp][c] {
					break
				}
			}

			setScore(l, c, upScore*downScore)
		}
	}

	max := 0
	for _, val := range scores {
		if val > max {
			max = val
		}
	}

	return max
}
