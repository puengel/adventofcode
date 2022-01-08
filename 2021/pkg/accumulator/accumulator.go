package accumulator

import (
	"fmt"
	"strconv"
	"strings"
)

// Operation enum for Program code
type Operation string

// Values for Operations
const (
	ACC Operation = "acc"
	JMP Operation = "jmp"
	NOP Operation = "nop"
)

// Instruction of a program
type Instruction struct {
	Operation
	Value int
}

// Program - Code with internal ptr and accumulator
type Program struct {
	Code        []*Instruction
	ptr         int
	accumulator int
}

// Reset internal variables
func (p *Program) Reset() {
	p.ptr = 0
	p.accumulator = 0
}

// RunUntilLoop - returns value before an operation would have run twice
func (p *Program) RunUntilLoop() int {
	visited := make(map[int]bool)
	p.Reset()

	for !visited[p.ptr] {
		visited[p.ptr] = true

		op := p.Code[p.ptr]
		switch op.Operation {
		case ACC:
			p.accumulator += op.Value
			p.ptr++
		case JMP:
			p.ptr += op.Value
		case NOP:
			p.ptr++
		}

	}

	return p.accumulator
}

// RunUntilLast - returns value after last opteration in code array is executed. Error if ran into loop
func (p *Program) RunUntilLast() (int, error) {
	visited := make(map[int]bool)
	p.Reset()
	last := false

	for !last && !visited[p.ptr] {
		visited[p.ptr] = true
		if p.ptr == len(p.Code)-1 {
			last = true
		}

		op := p.Code[p.ptr]
		switch op.Operation {
		case ACC:
			p.accumulator += op.Value
			p.ptr++
		case JMP:
			p.ptr += op.Value
		case NOP:
			p.ptr++
		}
	}

	if !last {
		return p.accumulator, fmt.Errorf("Run into loop")
	}

	return p.accumulator, nil
}

// ParseProgram - parses input strings into a program
func ParseProgram(inputs []string) *Program {
	code := make([]*Instruction, len(inputs))

	for idx, in := range inputs {
		code[idx] = NewInstruction(in)
	}

	return NewProgram(code)
}

// NewProgram - return new *Program with given code
func NewProgram(code []*Instruction) *Program {
	return &Program{
		Code:        code,
		ptr:         0,
		accumulator: 0,
	}
}

// ToggleJMPNOP - switches JMP to NOP or NOP to JMP at given index
func ToggleJMPNOP(idx int, code *[]*Instruction) {
	switch (*code)[idx].Operation {
	case ACC:
		return
	case JMP:
		(*code)[idx].Operation = NOP
	case NOP:
		(*code)[idx].Operation = JMP
	}
}

// NewInstruction - create Instruction from string
func NewInstruction(in string) *Instruction {
	fields := strings.Split(in, " ")

	num, err := strconv.Atoi(fields[1])
	if err != nil {
		fmt.Println("atoi", err)
		return nil
	}

	return &Instruction{
		Operation: Operation(fields[0]),
		Value:     num,
	}
}
