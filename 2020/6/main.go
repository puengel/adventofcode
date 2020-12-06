package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strings"
)

func main() {

	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var resultA int
	var resultB int

	for _, group := range inputs {
		people := strings.Split(group, "\n")

		count := len(people)
		all := 0
		answers := make(map[string]int)
		for _, answer := range people {
			for _, char := range answer {
				answers[string(char)]++
				if answers[string(char)] == count {
					all++
				}

			}
		}

		resultA += len(answers)
		resultB += all

	}

	fmt.Printf("Part 1: %d\n", resultA)
	fmt.Printf("Part 2: %d\n", resultB)
}
