package main

import (
	"fmt"
)

func part1(filename string) {
	fmt.Println("--- Part 1 ---")

}

func part2(filename string) {
	fmt.Println("--- Part 2 ---")
}

// func readInputFile(filename string) []int64 {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	scanner := bufio.NewScanner(file)
// 	numbers := []int64{0} // always assume a leading zero
// 	for scanner.Scan() {
// 		line := scanner.Text()
//
// 		num, err := strconv.ParseInt(line, 10, 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		numbers = append(numbers, num)
// 	}
//
// 	// Sort them
// 	sort.Sort(IntSlice(numbers))
//
// 	max := numbers[len(numbers)-1]
//
// 	// add 3
// 	numbers = append(numbers, max+3)
//
// 	return numbers
// }