package day10

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Instruction interface {
	runCycle(registerX *int)
	needsMoreCycles() bool
	complete(registerX *int)
}

type Noop struct {
	cycles int
}

func makeNoop() Noop {
	return Noop{1}
}

func (noop *Noop) runCycle(registerX *int) {
	noop.cycles--
}

func (noop *Noop) needsMoreCycles() bool {
	return noop.cycles > 0
}

func (noop *Noop) complete(registerX *int) {
}

type AddX struct {
	cycles int
	value  int
}

func makeAddX(value int) AddX {
	return AddX{cycles: 2, value: value}
}

func (addX *AddX) runCycle(registerX *int) {
	addX.cycles--
}

func (addX *AddX) needsMoreCycles() bool {
	return addX.cycles > 0
}

func (addX *AddX) complete(registerX *int) {
	*registerX += addX.value
}

type Cpu struct {
	X                  int
	CyclesCount        int
	currentInstruction Instruction
}

func makeCpu() Cpu {
	return Cpu{X: 1, CyclesCount: 0, currentInstruction: nil}
}

func (cpu *Cpu) HasWork() bool {
	return cpu.currentInstruction != nil && cpu.currentInstruction.needsMoreCycles()
}

func (cpu *Cpu) pushInstruction(instruction Instruction) {
	if cpu.HasWork() {
		panic("Cannot push another instruction if already working on one.")
	}

	cpu.currentInstruction = instruction
}

func (cpu *Cpu) StartCycle() {
	if cpu.HasWork() {
		cpu.currentInstruction.runCycle(&cpu.X)
		cpu.CyclesCount++
	}
}

func (cpu *Cpu) CompleteCycle() {
	if cpu.currentInstruction != nil && !cpu.currentInstruction.needsMoreCycles() {
		cpu.currentInstruction.complete(&cpu.X)
	}
}

func Day10Part1() int {
	cpu := makeCpu()
	result := 0

	file, _ := f.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for {
		cpu.StartCycle()

		if cpu.CyclesCount == 20 ||
			cpu.CyclesCount == 60 ||
			cpu.CyclesCount == 100 ||
			cpu.CyclesCount == 140 ||
			cpu.CyclesCount == 180 ||
			cpu.CyclesCount == 220 {
			result += (cpu.CyclesCount * cpu.X)
		}

		cpu.CompleteCycle()
		if !cpu.HasWork() {
			if !scanner.Scan() {
				break // no more instructions
			}

			line := scanner.Text()
			parts := strings.Split(line, " ")
			if parts[0] == "noop" {
				noop := makeNoop()
				cpu.pushInstruction(&noop)
			} else if parts[0] == "addx" {
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				addx := makeAddX(value)
				cpu.pushInstruction(&addx)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func Day10Part2() {
	cpu := makeCpu()

	file, _ := f.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for {
		cpu.StartCycle()
		position := (cpu.CyclesCount - 1) % 40
		if position == 0 {
			fmt.Println()
		}
		if position >= cpu.X-1 && position <= cpu.X+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		cpu.CompleteCycle()
		if !cpu.HasWork() {
			if !scanner.Scan() {
				break // no more instructions
			}

			line := scanner.Text()
			parts := strings.Split(line, " ")
			if parts[0] == "noop" {
				noop := makeNoop()
				cpu.pushInstruction(&noop)
			} else if parts[0] == "addx" {
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				addx := makeAddX(value)
				cpu.pushInstruction(&addx)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
