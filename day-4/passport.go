package main

import (
	"fmt"
	"strconv"
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

func (p *passport) ValidData() bool {
	// if not a valid then not valid
	if !p.ValidNorthPoleCred() {
		fmt.Print("Missing keys\n")
		return false
	}

	return p.ValidBirthYear() && p.ValidIssueYear() &&
		p.ValidExpirationYear() &&
		p.ValidHeight() && p.ValidHairColor() && p.ValidEyeColor() &&
		p.ValidPassportID()
}

func (p passport) checkYear(key string, minYearInclusive, maxYearInclusive int64) bool {
	val, _ := p.fieldMap[key]
	if len(val) < 4 {
		fmt.Printf("Invalid %s (%s), too few digits", key, val)
		return false
	}

	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Printf("Invalid %s (%s), %v", key, val, err)
		return false
	}

	minimal := num >= minYearInclusive
	maximal := num <= maxYearInclusive

	if !(minimal && maximal) {
		fmt.Printf("%s: Year %d, greater then %d? %t, less then %d? %t, valid? %t\n", key, num, minYearInclusive, minimal, maxYearInclusive, maximal, minimal && maximal)
	}

	return minimal && maximal
}

func (p *passport) ValidBirthYear() bool      { return p.checkYear("byr", 1920, 2002) }
func (p *passport) ValidIssueYear() bool      { return p.checkYear("iyr", 2010, 2020) }
func (p *passport) ValidExpirationYear() bool { return p.checkYear("eyr", 2020, 2030) }

func (p *passport) ValidHeight() bool {
	val, _ := p.fieldMap["hgt"]
	if len(val) < 3 {
		fmt.Printf("hgt: %s, invalid length\n", val)
		return false
	}

	num := val[0 : len(val)-2]
	unit := val[len(val)-2:]

	realNumber, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		fmt.Printf("hcl: %s error: %v", val, err)
	}

	if unit == "in" {
		if realNumber < 59 || realNumber > 76 {
			fmt.Printf("hcl: %s out of range\n", val)
			return false
		}
	} else if unit == "cm" {
		if realNumber < 150 || realNumber > 193 {
			fmt.Printf("hcl: %s out of range\n", val)
			return false
		}
	} else {
		fmt.Printf("hcl: %s invalid unit %s\n", val, unit)
	}

	return true
}

func (p *passport) ValidHairColor() bool {
	val, _ := p.fieldMap["hcl"]
	if len(val) != 7 {
		fmt.Printf("hcl: %s invalid number of characters\n", val)
		return false
	}

	if val[0] != "#"[0] {
		fmt.Printf("hcl: %s invalid format\n", val)
		return false
	}

	for _, char := range val[1:] {
		isDigit := (char >= '0' && char <= '9')
		isHexDigit := (char >= 'a' && char <= 'f')

		if !(isDigit || isHexDigit) {
			fmt.Printf("hcl: %s; %s is number? %t, is hex digit? %t\n", val, string(char), isDigit, isHexDigit)

			return false
		}
	}

	// fmt.Printf("hcl: valid\n")

	return true
}

func (p *passport) ValidEyeColor() bool {
	val, _ := p.fieldMap["ecl"]

	valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range valid {
		if val == v {
			return true
		}
	}

	fmt.Printf("ecl: %s not in list %v\n", val, valid)

	return false
}

func (p *passport) ValidPassportID() bool {
	val, _ := p.fieldMap["pid"]
	if len(val) != 9 {
		fmt.Printf("pid: %s invalid number of characters\n", val)
		return false
	}
	return true
}

func (p passport) String() string {
	if len(p.fieldMap) == 0 {
		return fmt.Sprintf("%v (unparsed)", p.originalLines)
	}

	return fmt.Sprintf("%+v", p.fieldMap)
}