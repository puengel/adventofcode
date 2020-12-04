package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	sumValid := 0
	sumValidData := 0

	for _, line := range inputs {
		p := NewPassport(line)
		if p.Validate() {
			sumValid++

			if p.ValidateData() {
				sumValidData++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sumValid)
	fmt.Printf("Part 2: %d\n", sumValidData)
}

func NewPassport(input string) *Passport {
	vals := strings.Fields(input)
	pass := Passport{}
	for _, val := range vals {
		entry := strings.Split(val, ":")

		switch entry[0] {
		case "byr":
			pass.byr = entry[1]
		case "iyr":
			pass.iyr = entry[1]
		case "eyr":
			pass.eyr = entry[1]
		case "hgt":
			pass.hgt = entry[1]
		case "hcl":
			pass.hcl = entry[1]
		case "ecl":
			pass.ecl = entry[1]
		case "pid":
			pass.pid = entry[1]
		case "cid":
			pass.cid = entry[1]
		default:
			fmt.Println("damn")
		}
	}
	return &pass
}

func (p *Passport) Validate() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p *Passport) ValidateData() bool {
	//byr
	byr, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt
	cm := strings.HasSuffix(p.hgt, "cm")
	in := strings.HasSuffix(p.hgt, "in")
	if !cm && !in {
		return false
	}
	hgt, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
	if err != nil {
		return false
	}
	if cm && (hgt < 150 || hgt > 193) {
		return false
	}
	if in && (hgt < 59 || hgt > 76) {
		return false
	}

	// hcl
	regexHcl := "#[a-f|0-9]{6}"
	match, err := regexp.MatchString(regexHcl, p.hcl)
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	// ecl
	regexEcl := "(amb|blu|brn|gry|grn|hzl|oth)"
	match, err = regexp.MatchString(regexEcl, p.ecl)
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	// pid
	_, err = strconv.Atoi(p.pid)
	if err != nil {
		return false
	}
	if len(p.pid) != 9 {
		return false
	}

	return true

}
