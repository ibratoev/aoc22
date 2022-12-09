package day9

import (
	"bufio"
	"embed"
	_ "embed"
	mathUtils "ibratoev/aoc22/utils/math"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

const (
	Right string = "R"
	Left  string = "L"
	Up    string = "U"
	Down  string = "D"
)

type Knot struct {
	x int
	y int
}

func (this *Knot) Move(direction string) {
	switch direction {
	case Right:
		this.x++
	case Left:
		this.x--
	case Up:
		this.y++
	case Down:
		this.y--
	}
}

func (this *Knot) PullAfter(prev Knot) {
	distanceX := mathUtils.Abs(this.x - prev.x)
	distanceY := mathUtils.Abs(this.y - prev.y)

	if distanceX <= 1 && distanceY <= 1 {
		return
	}

	if this.x < prev.x {
		this.x++
	} else if this.x > prev.x {
		this.x--
	}

	if this.y < prev.y {
		this.y++
	} else if this.y > prev.y {
		this.y--
	}
}

func Day9Part1() int {
	file, _ := f.Open("input.txt")
	defer file.Close()

	head := Knot{}
	tail := Knot{}
	visited := make(map[Knot]bool)
	markVisited := func(t Knot) {
		visited[t] = true
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < steps; i++ {
			head.Move(direction)
			tail.PullAfter(head)
			markVisited(tail)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return len(visited)
}

const TAIL_COUNT = 9

func Day9Part2() int {
	file, _ := f.Open("input.txt")
	defer file.Close()

	head := Knot{}
	var tails [TAIL_COUNT]Knot
	for i := 0; i < TAIL_COUNT; i++ {
		tails[i] = Knot{}
	}

	visited := make(map[Knot]bool)
	markVisited := func(t Knot) {
		visited[t] = true
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < steps; i++ {
			head.Move(direction)

			last := head
			for t := 0; t < TAIL_COUNT; t++ {
				tails[t].PullAfter(last)
				last = tails[t]
			}
			markVisited(last)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return len(visited)
}
