package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {
	inputs, err := input.ReadInts("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := 0

	for _, val := range inputs {
		result += val
	}

	fmt.Printf("Part 1: %d\n", result)

	seen := make(map[int]bool, len(inputs))
	foundTwice := false
	result = 0

	for !foundTwice {

		for _, val := range inputs {
			result += val

			if seen[result] {
				foundTwice = true
				break
			}
			seen[result] = true
		}

	}

	fmt.Printf("Part 2: %d\n", result)

}
