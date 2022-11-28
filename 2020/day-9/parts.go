package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(filename string, preambleLength int) {
	return

	fmt.Println("--- Part 1 ---")
	numbers := readInputFile(filename)
	fmt.Printf("%+v\n", numbers)

	invalid := make([]int64, 0)
	for i, num := range numbers {
		if i < preambleLength {
			continue
		}

		preamble := numbers[i-preambleLength : i]
		fmt.Printf("%d. %d: %v\n", i, num, preamble)

		// Now loop for multiples
		valid := false
		for j, candidate := range preamble {
			if candidate >= num {
				// Number 0 and below aren't possible so
				// if the candidate is greater or equal to the number then
				// ignore
				continue
			}

			other := num - candidate
			fmt.Printf("  Checking c1 %d for %d\n", candidate, other)
			found := false
			for _, candidate2 := range preamble[j+1:] {
				fmt.Printf("    Checking c2 %d\n", candidate2)
				// scan of the other candidate
				if candidate2 == other {
					fmt.Printf("      %d and %d found for %d\n", candidate, candidate2, num)
					found = true
				}
			}

			if found {
				fmt.Printf("  %d is valid, found %d and %d\n", num, candidate, other)
				valid = true
				break
			}
		}

		if !valid {
			invalid = append(invalid, num)
		}
	}

	fmt.Printf("Invalid items: %v\n", invalid)
}

func part2(filename string, ign int) {
	fmt.Println("--- Part 2 ---")

	numbers := readInputFile(filename)
	fmt.Printf("%v\n", numbers)

	// Instead we find the minimum number of contigous sets that can add to the number

	target := int64(50047984) // Puzzle target
	// target := int64(127) // Sample target

	maxSetSize := len(numbers) - 1

	for i := 2; i < maxSetSize; i++ {
		fmt.Printf("Checking sets of %d\n", i)

		weakset := make([]int64, 0)
		for j := range numbers[0 : len(numbers)-i+1] {
			set := numbers[j : j+i]

			acc := int64(0)
			for _, x := range set {
				acc = acc + x
			}

			fmt.Printf("  Set %v - %d\n", set, acc)

			if acc == target {
				fmt.Println("    Weakness found")
				weakset = set
				break
			}
		}

		if len(weakset) != 0 {
			fmt.Printf("Weak set found: %v\n", weakset)

			smallest := int64(0)
			largest := int64(0)
			for _, num := range weakset {
				if smallest == 0 {
					smallest = num
				}

				if num < smallest {
					smallest = num
				}

				if num > largest {
					largest = num
				}
			}

			fmt.Printf("Weakness is: %d\n", smallest+largest)

			break
		}
	}

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