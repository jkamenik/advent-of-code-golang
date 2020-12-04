package main

import "fmt"

func part1(passports []passport) {
	v := 0
	valid := &v
	for i, item := range passports {
		valdity := "invalid"
		if item.ValidNorthPoleCred() {
			valdity = "valid"
			*valid = v + 1
		}

		fmt.Printf("%d. %s is %s\n", i+1, item.String(), valdity)
	}

	fmt.Printf("\nTotal:       %d\nTotal Valid: %d\n", len(passports), v)
}

func part2(passports []passport) {

}