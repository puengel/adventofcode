package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
)

func main() {
	initial, err := input.ReadInts("input.txt", ",")
	if err != nil {
		return
	}

	fishes := make(map[int]int)

	// set zeros
	for i := 0; i <= 8; i++ {
		fishes[i] = 0
	}

	// set initial
	for _, fish := range initial {
		fishes[fish] = fishes[fish] + 1
	}

	sum80 := 0
	// loop over
	for day := 0; day < 256; day++ {
		newFishes := make(map[int]int)
		for key, val := range fishes {
			if key == 0 {
				newFishes[6] += val
				newFishes[8] += val
			} else {
				newFishes[key-1] += val
			}
		}
		fishes = newFishes
		if day == 79 {
			for _, val := range fishes {
				sum80 += val
			}
		}
	}

	sum256 := 0
	for _, val := range fishes {
		sum256 += val
	}

	fmt.Printf("Part 1: fishes %d\n", sum80)
	fmt.Printf("Part 2: fishes %d\n", sum256)
}
