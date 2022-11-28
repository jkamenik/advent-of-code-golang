package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

	numbers := readInputFile(filename)

	fmt.Printf("%v\n", numbers)

	oneJolt := 0
	threeJolt := 1 // always 3 joints between adator and computer

	for i, num := range numbers[0 : len(numbers)-1] {
		next := numbers[i+1]
		diff := next - num

		fmt.Printf("%d jolt(s) between %d and %d\n", diff, num, next)

		if diff == 1 {
			oneJolt++
		} else if diff == 3 {
			threeJolt++
		}
	}

	fmt.Printf("1 jolt count: %d\n", oneJolt)
	fmt.Printf("3 jolt count: %d\n", threeJolt)
	fmt.Printf("Answer: %d\n", oneJolt*threeJolt)

}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}

func readInputFile(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int64{0} // always assume a leading zero
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	// Sort them
	sort.Sort(IntSlice(numbers))

	max := numbers[len(numbers)-1]

	// add 3
	numbers = append(numbers, max+3)

	return numbers
}