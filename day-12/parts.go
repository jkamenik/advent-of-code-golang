package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

	ferry := NewFerry()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Printf("** %s **\n", line)
		ferry.Move(line)
		fmt.Printf("  %v\n", ferry)
	}

	manDist := math.Abs(float64(ferry.north)) + math.Abs(float64(ferry.east))

	fmt.Printf("Final Position\n  %v\n  Manhatten Dist: %f\n", ferry, manDist)
}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}