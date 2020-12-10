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
	sort.Sort(IntSlice(numbers))

	fmt.Printf("%v\n", numbers)

	oneJolt := 0
	threeJolt := 1 // always 3 joints between adator and computer

	for i, num := range numbers[0 : len(numbers)-1] {
		next := numbers[i+1]
		diff := next - num

		if i == 0 {
			// Check the first step against an assumed 0
			if num == 1 {
				fmt.Printf("1 jolt(s) between 0 and %d\n", num)
				// first is a 1 step
				oneJolt++
			} else if num == 3 {
				fmt.Printf("3 jolt(s) between 0 and %d\n", num)
				threeJolt++
			}
		}

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
	rtn := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		rtn = append(rtn, num)
	}

	return rtn
}