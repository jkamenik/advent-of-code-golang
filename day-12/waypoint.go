package main

import (
	"fmt"
	"log"
	"strconv"
)

type Waypoint struct {
	Ferry *Ferry
	north int
	east  int
}

func NewWaypoint() *Waypoint {
	x := Waypoint{Ferry: NewFerry()}

	return &x
}

func (w Waypoint) String() string {
	return fmt.Sprintf("Waypoint{North: %d, East: %d, Ferry: %v", w.north, w.east, w.Ferry)
}

func (w *Waypoint) Move(dir string) error {
	// fmt.Printf("--- %s ---\n", dir)
	d := dir[0]
	u := dir[1:]

	unit, err := strconv.ParseInt(u, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	unit32 := int(unit)

	switch d {
	case 'N':
		w.North(unit32)
	case 'S':
		w.South(unit32)
	case 'E':
		w.East(unit32)
	case 'W':
		w.West(unit32)
	case 'F':
		w.Forward(unit32)
	case 'L':
		w.Left(unit32)
	case 'R':
		w.Right(unit32)
	}

	return nil
}

func (w *Waypoint) South(unit int) {
	w.north -= unit
}

func (w *Waypoint) North(unit int) {
	w.north += unit
}

func (w *Waypoint) East(unit int) {
	w.east += unit
}

func (w *Waypoint) West(unit int) {
	w.east -= unit
}

// Left turns left by unit degrees
func (w *Waypoint) Left(unit int) {
	w.Right(unit * -1)
}

// Right turns to the right by unit degrees
func (w *Waypoint) Right(unit int) {
}

// Forward moves the ferry forward N times
func (w *Waypoint) Forward(unit int) {
}