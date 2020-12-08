package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type InstructionType string

type Instruction struct {
	iType InstructionType
	offset int
}

type Registers struct {
	acc int
	pc int

	loop bool

	visited map[int]bool
}

func CreatRegisters() Registers {
	return Registers{ acc :0, pc : 0, loop: false, visited: make(map[int]bool)}
}

func (r *Registers) RunProgram(instructions []Instruction) {
	for r.pc < len(instructions) {
		inst := instructions[r.pc]

		if r.visited[r.pc] {
			r.loop = true
			break
		}
		r.visited[r.pc] = true
		switch inst.iType {
		case "nop": {
			r.pc += 1
		}
		case "jmp": {
			r.pc += inst.offset
		}
		case "acc": {
			r.pc += 1
			r.acc += inst.offset
		}
		}
	}
}

func ParseInstruction(instruction string) Instruction {
	split := strings.Split(instruction, " ")
	var inst = split[0]
	offset, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return Instruction{ InstructionType(inst), offset}
}

func ParseProgram(program string) []Instruction {
	instructions := strings.Split(program, "\n")
	instVec := make([]Instruction, 0)

	for _, i := range instructions {
		i = strings.TrimSpace(i)
		if i == "" {
			continue
		}
		instVec = append(instVec, ParseInstruction(i))
	}
	return instVec
}

func RunProgramImpl(program []Instruction) (bool, int) {
	registers := CreatRegisters()
	//fmt.Println(program)

	registers.RunProgram(program)
	return registers.loop, registers.acc
}

func RunProgram(program string) (bool, int) {
	parseProgram := ParseProgram(program)
	return RunProgramImpl(parseProgram)
}

func FixProgram(program string) int {
	parseProgram := ParseProgram(program)

	loop, acc := RunProgram(program)
	if !loop {
		return acc
	}

	for i:=0 ; i < len(parseProgram); i++ {
		oldInst := parseProgram[i]
		newInst := oldInst
		switch oldInst.iType {
		case "jmp": {
			newInst.iType = "nop"
		}
		case "nop": {
			newInst.iType = "jmp"
		}
		default: {
			continue
		}
		}
		parseProgram[i] = newInst
		loop, acc := RunProgramImpl(parseProgram)
		if !loop {
			return acc
		}
		parseProgram[i] = oldInst
	}

	fmt.Errorf("Failed to fix the program")
	return 0
}

func main() {
	content, err := ioutil.ReadFile("./input8.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, acc := RunProgram(string(content))
	fmt.Println("Solution day 8 part a ", acc)

	acc = FixProgram(string(content))
	fmt.Println("Solution day 8 part b ", acc)

}
