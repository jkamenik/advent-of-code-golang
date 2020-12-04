package main

import (
	"fmt"
	"strings"
)

type passport struct {
	originalLines []string
	fieldMap      map[string]string
}

func newPassport(fields []string) (passport, error) {
	pass := passport{originalLines: fields, fieldMap: make(map[string]string, 0)}
	err := pass.parse()

	return pass, err
}

func (p *passport) parse() error {
	for _, line := range p.originalLines {
		fields := strings.Split(line, " ")
		for _, field := range fields {
			items := strings.Split(field, ":")
			if len(items) < 2 {
				return fmt.Errorf("%s is not a key:value field", field)
			}

			p.fieldMap[items[0]] = items[1]
		}
	}

	return nil
}

func (p *passport) ValidNorthPoleCred() bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validity := true
	v := &validity

	for _, key := range required {
		_, hasKey := p.fieldMap[key]

		// if !hasKey {
		// 	fmt.Printf("Missing key %s\n", key)
		// }
		//
		// fmt.Printf("key: %s, validity before: %t, hasKey: %t (%t)\n", key, validity, hasKey, validity && hasKey)
		*v = validity && hasKey
	}

	return validity
}

func (p passport) String() string {
	if len(p.fieldMap) == 0 {
		return fmt.Sprintf("%v (unparsed)", p.originalLines)
	}

	return fmt.Sprintf("%+v", p.fieldMap)
}