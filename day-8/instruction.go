package main

import (
	"fmt"
	"strconv"
)

// Instruction represents a single instruction in a total stack of instructions
type Instruction struct {
	op    string
	input int

	index   int
	visited uint // The number of times I have been visited
}

// NewInstruction returns a new instruction based on string
// <instruction> <dir><value>
//
// instruction is a 3 char code
//   - "nop" - do nothing, but advance the instruction pointer
//   - "acc" - increase or decrese the accumulator
//   - "jmp" - move instruction pointer
func NewInstruction(human string, index int) (Instruction, error) {
	i := Instruction{index: index}

	i.op = human[0:3]

	pos := string(human[4])
	num, err := strconv.ParseInt(human[5:], 10, 32)
	if err != nil {
		return i, err
	}

	if pos == "-" {
		num = num * -1
	}

	i.input = int(num)

	return i, nil
}

func (i Instruction) String() string {
	return fmt.Sprintf("{op:%s input:%d index:%d visited:%d}", i.op, i.input, i.index, i.visited)
}