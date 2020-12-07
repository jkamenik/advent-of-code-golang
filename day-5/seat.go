package main

import (
	"fmt"
)

type seat struct {
	space string

	row    int
	column int
	id     int
}

// Represents the sortalbe interface
type seats struct {
	items []seat
}

func newFromLine(line string) seat {
	x := seat{space: line}

	x.row = binarySearch(line[0:len(line)-3], 127)
	x.column = binarySearch(line[len(line)-3:], 7)

	x.id = (x.row * 8) + x.column
	return x
}

func binarySearch(searchPath string, max int) int {
	fmt.Printf("Checking %s at max %d\n", searchPath, max)
	min := 0
	updatedMax := &max
	updatedMin := &min
	for _, char := range searchPath {
		distance := (((*updatedMax - *updatedMin) / 2) + 1)

		if char == 'F' || char == 'L' {
			// lower half, so move the max number back
			*updatedMax = *updatedMax - distance
			if *updatedMax < 0 {
				*updatedMax = 0
			}
		} else {
			// upper half, move the min number forward
			*updatedMin = *updatedMin + distance
			if *updatedMin > *updatedMax {
				*updatedMin = *updatedMax
			}
		}

		fmt.Printf("%s: %d x %d, dist:%d\n", string(char), *updatedMin, *updatedMax, distance)
	}

	// The converge so just return one of them
	return *updatedMax
}

func newSeats() *seats {
	x := seats{make([]seat, 0)}
	return &x
}

func (s seats) String() string {
	seats := ""
	for i, s := range s.items {
		if i > 0 {
			seats += ", "
		}

		seats += s.String()
	}
	return fmt.Sprintf("seats[%s]", seats)
}

func (s *seats) Append(line string) {
	s.items = append(s.items, newFromLine(line))
}

func (s seat) String() string {
	return fmt.Sprintf("seat{space:'%s', row:%d, column:%d, id:%d}", string(s.space), s.row, s.column, s.id)
}

// Sortable interface
func (s *seats) Len() int           { return len(s.items) }
func (s *seats) Swap(i, j int)      { s.items[i], s.items[j] = s.items[j], s.items[i] }
func (s *seats) Less(i, j int) bool { return s.items[i].id < s.items[j].id }