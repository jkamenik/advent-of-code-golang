package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(filename string) {
	fmt.Println("--- Part 1 ---")
	numbers := readInputFile(filename)
	fmt.Printf("%+v\n", numbers)
}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}

func readInputFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rtn := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		rtn = append(rtn, int(num))
	}

	return rtn
}