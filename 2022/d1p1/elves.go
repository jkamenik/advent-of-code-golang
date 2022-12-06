package main

import (
	"errors"
	"fmt"
	"strconv"
)

var Done = errors.New("Done")

type Elf struct {
	calories []int
}

func NewElf() Elf {
	e := Elf{}
	e.calories = make([]int, 0)

	return e
}

func (e *Elf) AddCalories(line string) error {
	num, err := strconv.ParseInt(line, 10, 32)
	if err != nil {
		return fmt.Errorf("Invalid number %s", line)
	}

	fmt.Printf("Adding %v\n", num)
	e.calories = append(e.calories, int(num))
	fmt.Printf("%v: %v\n",len(e.calories), e.calories)

	return nil
}

func (e Elf) String() string {
	return fmt.Sprintf("%v",e.calories)
}

type Elves struct {
	all []Elf
}

func (e *Elves) readLine(line string) error {
	var elf Elf

	if e.all == nil {
		fmt.Println("No Elves, adding the first one")
		elf = Elf{}
		e.all = make([]Elf, 1)
		e.all[0] = elf
	} else {
		elf = e.all[len(e.all)-1]
	}
	fmt.Printf("Elf: %v %v\n", &elf, elf)

	if line == "" {
		fmt.Println("Empty line means new elf")
		elf = NewElf()
		e.all = append(e.all, elf)
	}

	return elf.AddCalories(line)
}
