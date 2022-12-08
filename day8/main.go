package day8

import (
	"embed"
	_ "embed"
	"io/fs"
)

//go:embed input.txt
var f embed.FS

func readFile(file fs.File, a []byte) int {
	i, err := file.Read(a)
	if err != nil {
		panic("Error reading from file")
	}
	return i
}

func Day8Part1() int {
	file, _ := f.Open("input.txt")
	defer file.Close()

	var trees [99][99]byte

	visible := make(map[int]bool)
	markVisible := func(l int, c int) {
		visible[l*100+c] = true
	}

	var maxColumns [99]byte

	for l := 0; l < 99; l++ {
		line := trees[l][:]
		readFile(file, line)
		var maxColumn byte = 0
		for c := 0; c < 99; c++ {
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
		for c2 := 98; c2 > -1; c2-- {
			if line[c2] > maxX2 {
				maxX2 = line[c2]
				markVisible(l, c2)
			}
		}

		if l != 98 {
			// READ /n to remove it from the stream
			var newLine [1]byte
			readFile(file, newLine[:])
			if newLine[0] != '\n' {
				panic("EXPECTED NEW LINE!")
			}
		}
	}

	for c := 0; c < 99; c++ {
		var maxC byte = 0
		// CHECK column <-
		for l := 98; l > -1; l-- {
			if trees[l][c] > maxC {
				maxC = trees[l][c]
				markVisible(l, c)
			}
		}
	}

	return len(visible)
}
