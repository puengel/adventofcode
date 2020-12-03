package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {

	lines, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Print(err)
		return
	}

	modX := len(lines[0])

	x := 0

	trees := 0

	for _, line := range lines {
		if (string(line[x])) == "#" {
			trees++
		}

		x += 3
		x = x % modX
	}

	fmt.Printf("Part 1: %d\n", trees)

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	result := 1

	for _, slope := range slopes {
		trees = 0
		dx := slope[0]
		dy := slope[1]
		y := 0
		x := 0
		for y < len(lines) {
			if (string(lines[y][x])) == "#" {
				trees++
			}
			y += dy
			x += dx
			x = x % modX
		}
		result *= trees
	}

	fmt.Printf("Part 2: %d\n", result)
}
