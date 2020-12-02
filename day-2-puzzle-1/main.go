package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 2")

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

	passwords, err := scanFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Answers:")
	count := 0
	loopCount := &count
	for i := 0; i < len(passwords); i++ {
		p := passwords[i]
		if !p.IsValid() {
			*loopCount = *loopCount + 1
		}
		fmt.Printf("  %s\n", p)
	}

	fmt.Printf("Total Invalid: %d\n", count)
}

func scanFile(fileName string) ([]*password, error) {
	passwords := make([]*password, 0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pass, err := NewPassword(line)
		if err != nil {
			return nil, err
		}

		passwords = append(passwords, pass)
	}

	return passwords, nil
}
