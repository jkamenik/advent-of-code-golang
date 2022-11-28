package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func part1(filename string) {
	seats := readInputFile(filename)
	sort.Sort(seats)

	fmt.Printf("%+v\n", seats)

	fmt.Printf("\nLarges id: %d\n", seats.items[len(seats.items)-1].id)
}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
	// gaps := []int

	seats := readInputFile(filename)
	sort.Sort(seats)

	lastSeat := -1
	x := &lastSeat
	for _, seat := range seats.items {
		if seat.id-1 != lastSeat {
			fmt.Printf("Missing seat before: %s\n Last seat ID: %d\n", seat.String(), lastSeat)
		}

		*x = seat.id
	}
}

func readInputFile(fileName string) *seats {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	Seats := newSeats()
	for scanner.Scan() {
		line := scanner.Text()

		Seats.Append(line)
	}

	return Seats
}