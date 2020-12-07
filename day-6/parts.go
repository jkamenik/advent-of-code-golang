package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(filename string) {
	items := reduceAny(readInputFile(filename))

	count := 0
	c := &count

	for _, group := range items {
		*c = count + len(group)
	}

	fmt.Printf("%+v\n", items)

	fmt.Printf("Sum: %d\n", count)
}

func part2(filename string) {
	items := reduceAll(readInputFile(filename))

	count := 0
	c := &count

	for _, group := range items {
		count := group["count"]

		for key, answer := range group {
			if key == "count" {
				continue
			}

			if answer == count {
				fmt.Printf("all members agreed to %s\n", key)
				*c = *c + 1
			}
		}
	}

	fmt.Printf("%+v\n", items)

	fmt.Printf("Sum: %d\n", count)
}

func readInputFile(fileName string) [][]string {
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

	return groups
}

func reduceAny(groups [][]string) []map[string]bool {
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

func reduceAll(groups [][]string) []map[string]int {
	items := []map[string]int{}

	for _, g := range groups {
		item := map[string]int{}

		for _, line := range g {
			item["count"] = len(g)

			for _, letter := range line {
				if x, ok := item[string(letter)]; !ok {
					item[string(letter)] = 1
				} else {
					item[string(letter)] = x + 1
				}

				// fmt.Println(string(letter))
			}
		}

		items = append(items, item)

		// fmt.Println(len(item))
	}

	return items
}