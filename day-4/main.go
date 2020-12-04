package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Printf("%s <file>\n\n", os.Args[0])
	fmt.Println(`
		Reads password files and prints valid count`)
}

func main() {
	fmt.Println("Day 4")

	if len(os.Args) <= 1 {
		fmt.Println("Error: Missing required inputs")
		usage()
		os.Exit(1)
	}

	fileName := os.Args[1] // 0 is the command, 1 is the first argument

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Printf("Error: file '%s' does not exist\n", fileName)
		os.Exit(1)
	}

	part1(readPassports(fileName))
	part2(readPassports(fileName))
}

func readPassports(fileName string) []passport {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	passports := make([]passport, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			// collect more input
			lines = append(lines, line)
			continue
		}

		passport, err := newPassport(lines)
		if err != nil {
			log.Fatal(err)
		}

		passports = append(passports, passport)

		// Finish by making a new array
		lines = make([]string, 0)
	}

	if len(lines) != 0 {
		passport, err := newPassport(lines)
		if err != nil {
			log.Fatal(err)
		}
		passports = append(passports, passport)
	}

	return passports
}