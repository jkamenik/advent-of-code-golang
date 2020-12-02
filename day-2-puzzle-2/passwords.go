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

	Pos1Char string
	Pos2Char string
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

	// Set positional characters
	pass.Pos1Char = string(pass.Password[pass.Least-1])
	pass.Pos2Char = string(pass.Password[pass.Most-1])

	return &pass, nil
}

func (p password) IsValid() bool {
	// Xor
	return (p.Pos1Char == p.Char) != (p.Pos2Char == p.Char)
}

func (p *password) String() string {
	validity := "Invalid"
	if p.IsValid() {
		validity = "Valid"
	}

	return fmt.Sprintf("%d-%d %s: %s (%s: pos1-%s, pos2-%s)", p.Least, p.Most, p.Char, p.Password, validity, p.Pos1Char, p.Pos2Char)
}
