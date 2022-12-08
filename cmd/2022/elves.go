package twentyTwentyTwo

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
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

	log.Info().Msgf("Adding %v\n", num)
	e.calories = append(e.calories, int(num))
	log.Debug().Msgf("%v: %v\n", len(e.calories), e.calories)

	return nil
}

func (e Elf) Total() int {
	var sum int = 0

	for _, val := range e.calories {
		sum = sum + val
	}

	return sum
}

func (e Elf) String() string {
	return fmt.Sprintf("%v", e.calories)
}

type Elves struct {
	all []Elf
}

func NewElves() Elves {
	e := Elves{make([]Elf, 0)}
	return e
}

func (e *Elves) ReadLine(line string) error {
	var elf Elf

	if e.all == nil {
		log.Info().Msg("No Elves array, creating")
		e.all = make([]Elf, 0)
	}

	if len(e.all) < 1 || line == "" {
		log.Info().Str("line", line).Msg("No Elf, creating")
		elf = NewElf()
		e.all = append(e.all, elf)
	} else {
		elf = e.all[len(e.all)-1]
	}

	log.Info().Str("line",line).Msgf("AddCalorie to %v", elf)
	return elf.AddCalories(line)
}
