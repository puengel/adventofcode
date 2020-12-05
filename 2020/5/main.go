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

	highest := 0
	seats := make([]bool, 1016)

	for _, pass := range inputs {
		add := 64
		row := 0
		for i := 0; i < 7; i++ {
			if string(pass[i]) == "B" {
				row += add
			}
			add = add / 2
		}

		add = 4
		col := 0

		for i := 7; i < 10; i++ {
			if string(pass[i]) == "R" {
				col += add
			}
			add = add / 2
		}

		id := (row * 8) + col
		seats[id] = true
		if id > highest {
			highest = id
		}
	}

	for idx, seat := range seats {
		if idx == 0 || idx == len(seats)-1 {
			continue
		}

		if !seat && seats[idx-1] && seats[idx+1] {
			fmt.Printf("Part 2: %d\n", idx)
		}
	}

	fmt.Printf("Part 1: %d\n", highest)
}
