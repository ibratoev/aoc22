package day6

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

func Day6Part1() int {
	file, _ := f.Open("input.txt")
	defer file.Close()
	var last4 [4]byte
	i := 0
	for {
		readFile(file, last4[i%4:i%4+1])
		if i > 3 && last4[0] != last4[1] && last4[0] != last4[2] && last4[0] != last4[3] &&
			last4[1] != last4[2] && last4[1] != last4[3] && last4[2] != last4[3] {
			return i + 1
		}
		i++
	}
}

func checkDuplicates(window *[14]byte) bool {
	for j := 0; j < 13; j++ {
		for k := j + 1; k < 14; k++ {
			if window[j] == window[k] {
				return true
			}
		}
	}

	return false
}

func Day6Part2() int {
	file, _ := f.Open("input.txt")
	defer file.Close()
	var window [14]byte
	i := 0
	for {
		readFile(file, window[i%14:i%14+1])
		if i > 13 {
			if !checkDuplicates(&window) {
				return i + 1
			}
		}
		i++
	}
}
