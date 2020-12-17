package main

import (
	"fmt"
	"log"
	"strconv"
)

type Ferry struct {
	facing int
	north  int
	east   int
}

func NewFerry() *Ferry {
	ferry := Ferry{facing: 90}

	return &ferry
}

func (f Ferry) String() string {
	facing := ""
	switch f.facing {
	case 0:
		facing = "North"
	case 90:
		facing = "East"
	case 180:
		facing = "South"
	case 270:
		facing = "West"
	}

	return fmt.Sprintf("Ferry{Facing: %s, North: %d, East: %d}", facing, f.north, f.east)
}

func (f *Ferry) Move(dir string) error {
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
		f.North(unit32)
	case 'S':
		f.South(unit32)
	case 'E':
		f.East(unit32)
	case 'W':
		f.West(unit32)
	case 'F':
		f.Forward(unit32)
	case 'L':
		f.Left(unit32)
	case 'R':
		f.Right(unit32)
	}

	return nil
}

func (f *Ferry) South(unit int) {
	f.north -= unit
}

func (f *Ferry) North(unit int) {
	f.north += unit
}

func (f *Ferry) East(unit int) {
	f.east += unit
}

func (f *Ferry) West(unit int) {
	f.east -= unit
}

// Left turns left by unit degrees
func (f *Ferry) Left(unit int) {
	f.Right(unit * -1)
}

// Right turns to the right by unit degrees
func (f *Ferry) Right(unit int) {
	newDir := f.facing + unit

	fmt.Printf("Facing %d, new Dir %d\n", f.facing, newDir)

	// For now do the brute force common dirs
	if newDir == 0 || newDir == 360 {
		// north
		f.facing = 0
	} else if newDir == 90 || newDir == -270 || newDir == 450 {
		// east
		f.facing = 90
	} else if newDir == 180 || newDir == -180 || newDir == 540 {
		// south
		f.facing = 180
	} else if newDir == 270 || newDir == -90 {
		// west
		f.facing = 270
	} else {
		log.Fatal(fmt.Sprintf("Unsupported turn (%d), needs to be 0, 90, 180, or 270\n", newDir))
	}

	fmt.Printf("New direction: %d, now facing %d\n", newDir, f.facing)
}

// Forward moves forward in what ever direction we are facing
func (f *Ferry) Forward(unit int) {
	switch f.facing {
	case 0:
		f.North(unit)
	case 180:
		f.South(unit)
	case 90:
		f.East(unit)
	case 270:
		f.West(unit)
	}
}