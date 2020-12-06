package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	doubles := 0
	triples := 0

	for _, input := range inputs {

		seen := make(map[string]int, len(input))

		for _, c := range input {
			seen[string(c)]++
		}

		has2 := false
		has3 := false
		for _, b := range seen {
			if b == 2 {
				has2 = true
			}
			if b == 3 {
				has3 = true
			}
		}

		if has2 {
			doubles++
		}
		if has3 {
			triples++
		}

	}

	fmt.Printf("Part 1: %d\n", doubles*triples)
	var result string

	for idx, a := range inputs {
		for _, b := range inputs[idx:] {

			one := false
			pos := -1

			for i := 0; i < len(a); i++ {
				if a[i] == b[i] {
					continue
				}

				if !one {
					one = true
					pos = i
				} else {
					one = false
					break
				}
			}

			if one {
				fmt.Printf("A: %s\nB: %s\nPos: %d\n\n", a, b, pos)
				result = a[:pos] + a[pos+1:]
				break
			}
		}
	}

	fmt.Printf("Part 2: %s\n", result)
}
