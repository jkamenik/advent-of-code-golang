package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, ",")

		seed := make([]int, 0)
		for _, i := range s {
			x, err := strconv.ParseInt(i, 10, 32)
			if err != nil {
				log.Fatal(err)
			}

			seed = append(seed, int(x))
		}

		game := NewMemoryGame(seed)

		fmt.Printf("Starting: %+v\n", game)

		for i := len(seed); i < 2020; i++ {
			game.Turn()

			fmt.Printf("Turn %d: %+v\n", game.turn, game)
		}
	}
}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}