package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"sort"
)

func main() {
	inputs, err := input.ReadInts("input.txt", ",")
	if err != nil {
		return
	}

	// test median
	sort.Ints(inputs)
	median := inputs[len(inputs)/2]

	fuel := 0
	for _, val := range inputs {
		diff := median - val
		if diff > 0 {
			fuel += diff
		} else {
			fuel -= diff
		}
	}

	fmt.Printf("Part 1: fuel %d\n", fuel)

	// inputs = []int{
	// 	16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	// }
	// sort.Ints(inputs)

	cheapest := 9223372036854775807 // maxint
	for pos := inputs[0]; pos <= inputs[len(inputs)-1]; pos++ {
		fuel = 0
		for _, val := range inputs {
			diff := pos - val
			if diff < 0 {
				diff = -diff
			}
			fuel += (diff * (diff + 1)) / 2
		}

		if fuel < cheapest {
			cheapest = fuel
		}
	}

	fmt.Printf("Part 2: fuel %d\n", cheapest)
}
