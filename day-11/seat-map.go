package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	floor = "."
	empty = "L"
	full  = "#"
)

// SeatMap represents a single instance seatmap
// Each location is a single Rune
//  - "." = empty, and is never filed
//  - "L" = an empty seat
//  - "#" = a full seat
type SeatMap []string

// NewSeatMap creates a new SeatMap from a file
func NewSeatMap(filename string) SeatMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	return SeatMap(lines)
}

func (s SeatMap) String() string {
	return strings.Join(s, "\n")
}

// Same return true of one SeatMap is the same as the other
func (s SeatMap) Same(other SeatMap) bool {
	if len(s) != len(other) {
		return false
	}

	// now loop
	for i, line := range s {
		if line != other[i] {
			return false
		}
	}

	return true
}

// Iterate generates a new SeatMap using the following rules:
// 1. Empty areas (.) are ignored
// 2. Empty Seats (L) w/ no adjacent occupied seats (#) become occupied
// 3. Occupied Seats (#) w/ 4 or more adjacent occupied seats become empty
// 4. Otherwise the seat doesn't change.
func (s SeatMap) Iterate() SeatMap {
	next := []string{}

	for i, line := range s {
		nextLine := ""

		for j, char := range line {
			nextLine = nextLine + s.nextStringFor(i, j, string(char))
		}

		next = append(next, nextLine)
	}

	return next
}

func (s SeatMap) nextStringFor(i, j int, char string) string {
	if char == floor {
		return floor
	}

	above := s.above(i, j)
	below := s.below(i, j)
	besides := s.besides(i, j)
	occupied := countOccupied(above + below + besides)

	// 	fmt.Printf(`
	// ---
	// %dx%d=%d
	// %s
	// %s%s%s
	// %s
	// ---`, i, j, occupied, above, string(besides[0]), char, string(besides[1]), below)

	if char == empty {
		if occupied == 0 {
			return full
		}
		return char
	}

	// full seats
	if occupied >= 4 {
		return empty
	}

	return char
}

func (s SeatMap) above(i, j int) string {
	above := i - 1
	if above < 0 {
		// Assume unlisted row is floors
		return floor + floor + floor
	}

	aboveLine := s[above]
	rtn := ""

	before := j - 1
	if before < 0 {
		rtn = floor
	} else {
		rtn = string(aboveLine[before])
	}

	// there is no case where direct above j is out of bounds
	rtn = rtn + string(aboveLine[j])

	after := j + 1
	if after > len(aboveLine)-1 {
		rtn = rtn + floor
	} else {
		rtn = rtn + string(aboveLine[after])
	}

	return rtn
}

func (s SeatMap) below(i, j int) string {
	below := i + 1
	if below > len(s)-1 {
		// Assume unlisted row is floors
		return floor + floor + floor
	}

	belowLine := s[below]
	rtn := ""

	before := j - 1
	if before < 0 {
		rtn = floor
	} else {
		rtn = string(belowLine[before])
	}

	// there is no case where direct above j is out of bounds
	rtn = rtn + string(belowLine[j])

	after := j + 1
	if after > len(belowLine)-1 {
		rtn = rtn + floor
	} else {
		rtn = rtn + string(belowLine[after])
	}

	return rtn
}

func (s SeatMap) besides(i, j int) string {
	line := s[i]
	rtn := ""

	before := j - 1
	if before < 0 {
		rtn = floor
	} else {
		rtn = string(line[before])
	}

	after := j + 1
	if after > len(line)-1 {
		rtn = rtn + floor
	} else {
		rtn = rtn + string(line[after])
	}

	return rtn
}

func countOccupied(line string) int {
	count := 0
	for _, char := range line {
		if string(char) == full {
			count++
		}
	}

	return count
}