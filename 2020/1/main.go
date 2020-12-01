package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {

	values, err := input.ReadInts("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Part 1
	for idx, first := range values {
		for _, second := range values[idx:] {
			if first+second == 2020 {
				fmt.Printf("First: %d\nSecond: %d\nResult: %d\n", first, second, (first * second))
			}
		}
	}

	// Part 2
	for idxFirst, first := range values {
		for idxSecond, second := range values[idxFirst:] {
			for _, third := range values[idxSecond:] {
				if first+second+third == 2020 {
					fmt.Printf("First: %d\nSecond: %d\nThird: %d\nResult: %d\n", first, second, third, (first * second * third))
				}
			}
		}
	}
}
