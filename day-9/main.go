package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func usage() {
	fmt.Printf("%s <file> <preample length>\n\n", os.Args[0])
}

func main() {
	fmt.Println("Day 8")

	if len(os.Args) <= 2 {
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

	preamble, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	part1(fileName, int(preamble))
	part2(fileName, int(preamble))
}