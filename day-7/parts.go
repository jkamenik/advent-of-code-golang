package main

import "fmt"

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

	stack := NewStackFromFile(filename)
	err := stack.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", stack)

	fmt.Printf("Current Instruction: %d\nAccumulator: %d\n", stack.current, stack.accumulator)

}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}