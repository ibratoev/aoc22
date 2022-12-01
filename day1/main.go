package day1

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Day1() int {
	var mostCal = 0

	var elfCal = 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		if strings.TrimSpace(line) == "" {
			elfCal = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elfCal += i

			if mostCal < elfCal {
				mostCal = elfCal
			}
		}
	}

	fmt.Println(mostCal)
	return mostCal
}
