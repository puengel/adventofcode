package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {

	inputs, err := input.ReadInts("input.txt", "\n")
	if err != nil {
		return
	}

	// Part 1
	prev := inputs[0]
	increases := 0
	for _, i := range inputs[1:] {
		if i > prev {
			increases++
		}
		prev = i
	}

	fmt.Printf("Part 1 Increases: %d\n", increases)

	// Part 2
	prevWindow := inputs[0] + inputs[1] + inputs[2]
	increases = 0
	for idx := range inputs {
		if idx < 2 {
			continue
		}

		newWindow := inputs[idx] + inputs[idx-1] + inputs[idx-2]
		if newWindow > prevWindow {
			increases++
		}
		prevWindow = newWindow
	}

	fmt.Printf("Part 2 Increases: %d\n", increases)
}
