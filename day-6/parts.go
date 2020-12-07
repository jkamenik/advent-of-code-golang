package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(filename string) {
	items := readInputFile(filename)

	count := 0
	c := &count

	for _, group := range items {
		*c = count + len(group)
	}

	fmt.Printf("%+v\n", items)

	fmt.Printf("Sum: %d\n", count)
}

func part2(filename string) {
}

func readInputFile(fileName string) []map[string]bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	groups := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			// collect more input
			lines = append(lines, line)
			continue
		}

		if len(lines) == 0 {
			// bail until there are lines
			continue
		}

		groups = append(groups, lines)
		lines = make([]string, 0)
	}

	if len(lines) != 0 {
		groups = append(groups, lines)
	}

	return reduce(groups)
}

func reduce(groups [][]string) []map[string]bool {
	items := []map[string]bool{}

	for _, g := range groups {
		item := map[string]bool{}

		for _, line := range g {
			for _, letter := range line {
				item[string(letter)] = true
				// fmt.Println(string(letter))
			}
		}

		items = append(items, item)

		// fmt.Println(len(item))
	}

	return items
}