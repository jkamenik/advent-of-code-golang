package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("%s <file> <preample length>\n\n", os.Args[0])
}

func main() {
	fmt.Println("Day 13")

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

	part1(fileName)
	part2(fileName)
}