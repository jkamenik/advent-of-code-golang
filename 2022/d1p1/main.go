package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elves := Elves{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		elves.readLine(line)
	}

	fmt.Println(elves)
}
