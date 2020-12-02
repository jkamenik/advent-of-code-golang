package main

import (
	"fmt"
	"strconv"
	"strings"
)

type password struct {
	Least    int64
	Most     int64
	Char     string
	Password string

	Count int64
}

// NewPassword splits a string in the format of
// <least>-<most> <character>: <password>
// into it parts.
func NewPassword(input string) (*password, error) {
	pass := password{}

	splits := strings.Split(input, ": ")
	if len(splits) < 2 {
		return nil, fmt.Errorf("Input missing password part: '%s'", input)
	}

	pass.Password = splits[1]

	splits = strings.Split(splits[0], " ")
	if len(splits) < 2 {
		return nil, fmt.Errorf("Input missing character part: '%s'", input)
	}

	pass.Char = splits[1]

	splits = strings.Split(splits[0], "-")
	if len(splits) < 2 {
		return nil, fmt.Errorf("Input missing most and least count: '%s'", input)
	}

	least, err := strconv.ParseInt(splits[0], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Least count is not a digit: '%s'", input)
	}
	pass.Least = least

	most, err := strconv.ParseInt(splits[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Most count is not a digit: '%s'", input)
	}
	pass.Most = most

	pass.CountOccurances()

	return &pass, nil
}

func (p *password) CountOccurances() {
	splits := strings.Split(p.Password, p.Char)

	// The split slurps the split string leaving the other parts.
	// So there should be 1 more split item then we expect
	p.Count = int64(len(splits) - 1)
}

func (p password) IsValid() bool {
	return (p.Count >= p.Least) && (p.Count <= p.Most)
}

func (p *password) String() string {
	validity := "Invalid"
	if p.IsValid() {
		validity = "Valid"
	}

	return fmt.Sprintf("%d-%d %s: %s (%d / %s)", p.Least, p.Most, p.Char, p.Password, p.Count, validity)
}
