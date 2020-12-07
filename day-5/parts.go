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