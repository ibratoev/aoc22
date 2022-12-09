package day9

import (
	"bufio"
	"embed"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const (
	Right        = "R"
	Left  string = "L"
	Up           = "U"
	Down         = "D"
)

type Head struct {
	x int
	y int
}

func (this *Head) Move(direction string) {
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

type Tail Head

func (this *Tail) Move(head Head) {
	distanceX := Abs(this.x - head.x)
	distanceY := Abs(this.y - head.y)

	if distanceX <= 1 && distanceY <= 1 {
		return
	}

	if this.x < head.x {
		this.x++
	} else if this.x > head.x {
		this.x--
	}

	if this.y < head.y {
		this.y++
	} else if this.y > head.y {
		this.y--
	}
}

func Day8Part1() int {
	file, _ := f.Open("input.txt")
	defer file.Close()

	head := Head{100, 100}
	tail := Tail{100, 100}
	visited := make(map[Tail]bool)
	markVisited := func(t Tail) {
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
			tail.Move(head)
			markVisited(tail)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return len(visited)
}
