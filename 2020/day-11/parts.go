package main

import (
	"fmt"
)

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

	pre := NewSeatMap(filename)
	fmt.Printf("%v\n\n", pre)

	for true {
		next := pre.IteratePart1()
		fmt.Printf("%v\n\n", next)

		if pre.Same(next) {
			break
		}

		pre = next
	}

	occupied := countOccupied(pre.String())
	fmt.Printf("Occupied: %d\n", occupied)
}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}