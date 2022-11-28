package main

import "fmt"

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

	stack := NewStackFromFile(filename)
	err := stack.Run()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("%+v\n", stack)

	fmt.Printf("Current Instruction: %d\nAccumulator: %d\n", stack.current, stack.accumulator)

}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")

	stack := NewStackFromFile(filename)

	for i, op := range stack.instructions {
		fmt.Printf("---- Checking op %d: %v\n", i, op)
		orig := op.op

		switch op.op {
		case "jmp":
			op.op = "nop"
		case "nop":
			op.op = "jmp"
		default:
			fmt.Println("Acc is ignored")
			continue
		}

		err := stack.Run()
		fmt.Printf("%+v\n", stack)
		if err != nil && err.Error() == "Infinite Loop" {
			fmt.Println(err)
			fmt.Printf("Incorrect swap instruction: %v, switching back to %s\n", op, orig)
			op.op = orig
		} else {
			fmt.Printf("Done at %d: %v\n", i, op)
			break
		}
	}

	fmt.Printf("Acc:%d\n", stack.accumulator)
}