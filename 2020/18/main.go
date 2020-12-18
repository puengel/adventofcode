package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	r1 := regexp.MustCompile(`\([0-9|\s|\+|\*]+\)`)

	var sum uint64

	for _, line := range lines {

		for inner := r1.FindString(line); inner != ""; inner = r1.FindString(line) {
			trimmed := strings.TrimSuffix(strings.TrimPrefix(inner, "("), ")")
			res := eval(trimmed)
			line = strings.Replace(line, inner, fmt.Sprintf("%d", res), 1)
		}

		vLine := eval(line)
		sum += vLine
	}

	fmt.Printf("Part 1: %d\n", sum)

	lines, err = input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var sum2 uint64

	for _, line := range lines {

		for inner := r1.FindString(line); inner != ""; inner = r1.FindString(line) {
			trimmed := strings.TrimSuffix(strings.TrimPrefix(inner, "("), ")")
			res := eval2(trimmed)
			line = strings.Replace(line, inner, fmt.Sprintf("%d", res), 1)
		}

		vLine := eval2(line)
		sum2 += vLine
	}

	fmt.Printf("Part 2: %d\n", sum2)
}

func eval2(str string) (result uint64) {

	r2 := regexp.MustCompile(`[0-9]+\s\+\s[0-9]+`)

	for inner := r2.FindString(str); inner != ""; inner = r2.FindString(str) {
		sVals := strings.Split(inner, " + ")
		op1, err := strconv.ParseUint(sVals[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		op2, err := strconv.ParseUint(sVals[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		addRes := op1 + op2
		str = strings.Replace(str, inner, fmt.Sprintf("%d", addRes), 1)
	}

	f := strings.Split(str, " * ")

	result = 1
	for _, next := range f {
		v, err := strconv.ParseUint(next, 10, 64)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		result *= v
	}

	return
}

func eval(str string) (result uint64) {
	f := strings.Split(str, " ")

	mult := false
	for _, next := range f {
		switch next {
		case "*":
			mult = true
		case "+":
			mult = false
		default:
			v, err := strconv.ParseUint(next, 10, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			if mult {
				result *= v
			} else {
				result += v
			}
		}
	}

	return
}
