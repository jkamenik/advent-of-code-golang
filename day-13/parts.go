package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	timeS := ""
	busesS := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if timeS == "" {
			timeS = line
		} else {
			busesS = line
			break
		}
	}

	t, err := strconv.ParseInt(timeS, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	time := int(t)

	fmt.Printf("Earliest Leave Time: %d\n", time)

	buses := buses{}
	for _, b := range strings.Split(busesS, ",") {
		if b == "x" {
			continue
		}

		b1, err := strconv.ParseInt(string(b), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		bus := NewBus(int(b1), time)

		buses = append(buses, bus)

		fmt.Printf("Bus: %+v\n", bus)
	}
	sort.Sort(buses)

	fmt.Printf("Buses: %+v\n", buses)

	earliest := buses[0]
	fmt.Printf("Earliest Bus: %+v\n", earliest)
	fmt.Printf("answer: %d\n", earliest.id*earliest.waitTime)
}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}