package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Stack represents the full instruction list
type Stack struct {
	instructions []*Instruction

	accumulator int
	current     int // current instruction
}

// NewStackFromFile loads a full stack from files
func NewStackFromFile(filename string) *Stack {
	s := &Stack{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		instr, err := NewInstruction(line, len(s.instructions))
		if err != nil {
			log.Fatal(err)
		}

		s.instructions = append(s.instructions, &instr)
	}

	return s
}

func (s *Stack) String() string {
	str := ""
	for idx, op := range s.instructions {
		if idx >= 1 {
			str = str + ", "
		}

		str = str + fmt.Sprintf("%+v", op)
	}

	return fmt.Sprintf("{acc:%d cur:%d inst:%s }", s.accumulator, s.current, str)
}

// Run runs the stack
func (s *Stack) Run() error {
	// always start at 0
	s.current = -1
	for true {
		s.current++
		if s.current >= len(s.instructions) {
			return errors.New("No such instruction")
		}

		op := s.instructions[s.current]
		op.visited++

		// Bail on infinite loop
		if op.visited > 1 {
			return errors.New("Infinite Loop")
		}

		switch op.op {
		case "nop":
			// do nothing
		case "acc":
			s.accumulator = s.accumulator + op.input
		case "jmp":
			// Go to right before the op
			s.current = s.current + op.input - 1
		}

		fmt.Printf("Op: %v\n  acc: %d, cur: %d\n", op, s.accumulator, s.current)

	}

	return nil
}

// func Acc(adjustment int) (int, error) {
//
// }
//
// func Jump(adjustment int) (int, error) {
//
// }