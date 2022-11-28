package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 1")

	if len(os.Args) <= 1 {
		fmt.Println("Error: No input file specified")
		os.Exit(1)
	}

	fileName := os.Args[1] // 0 is the command, 1 is the first argument

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Printf("Error: file '%s' does not exist\n", fileName)
		os.Exit(1)
	}

	nums, err := scanFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	tuples, err := findTuples(nums)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answers:")
	for i := 0; i < len(tuples); i++ {
		t := tuples[i]
		fmt.Printf("  %d. %d / %d: %d\n", i+1, t.A, t.B, t.Multiple())
	}
}

func scanFile(fileName string) ([]uint64, error) {
	nums := make([]uint64, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Invalid number %s", line)
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func findTuples(nums []uint64) ([]tuple, error) {
	length := len(nums)
	tuples := make([]tuple, 0)

	for i := 0; i < length; i++ {
		// First loop to find factors
		t := tuple{nums[i], 0, 2020}
		factor := t.OtherFactor()

		for j := i; j < length; j++ {
			// inner loop see if the other factor is found
			if nums[j] == factor {
				// found, complet the tuple
				t.B = factor
				tuples = append(tuples, t)
			}
		}
	}

	return tuples, nil
}
