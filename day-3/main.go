package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func usage() {
	fmt.Printf("%s <file> <right> <down>\n\n", os.Args[0])
	fmt.Println(`
		First reads the input file and determines if it will be wide enough given the inputs.  If not it will generate a wider map by repeating pattern.

		Then it will walk the array counting the number of trees ("#") it would hit as it walked.

		Then it will write its results to walk.txt, and print out the number of trees it hit.`)
}

func main() {
	fmt.Println("Day 3")

	if len(os.Args) <= 3 {
		fmt.Println("Error: Missing required inputs")
		usage()
		os.Exit(1)
	}

	fileName := os.Args[1] // 0 is the command, 1 is the first argument
	right := os.Args[2]
	down := os.Args[3]

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Printf("Error: file '%s' does not exist\n", fileName)
		os.Exit(1)
	}

	walkRight, err := strconv.ParseInt(right, 10, 32)
	walkDown, err1 := strconv.ParseInt(down, 10, 32)
	if err != nil || err1 != nil {
		fmt.Printf("Error: right or down value not an int\n")
		usage()
		os.Exit(1)
	}

	fmt.Printf("Walking %d x %d\n", walkRight, walkDown)

	field, err := newField(fileName)
	if err != nil {
		log.Fatal(err)
	}

	steps, hits, misses, err := field.walk(int(walkRight), int(walkDown))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`
-------------
%s
-------------
`, field.String())
	fmt.Printf(`
Hits:   %d
Misses: %d
Steps:  %d
`, hits, misses, steps)
}