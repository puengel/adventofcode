package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	// make map
	ventmap := make([][]int, 1000)
	for idx := range ventmap {
		ventmap[idx] = make([]int, 1000)
	}

	for _, vent := range inputs {
		lr := strings.Split(vent, " -> ")
		start := strings.Split(lr[0], ",")
		x1, err := strconv.Atoi(start[0])
		if err != nil {
			return
		}
		y1, err := strconv.Atoi(start[1])
		if err != nil {
			return
		}
		end := strings.Split(lr[1], ",")
		x2, err := strconv.Atoi(end[0])
		if err != nil {
			return
		}
		y2, err := strconv.Atoi(end[1])
		if err != nil {
			return
		}

		// diagonal
		if x1 != x2 && y1 != y2 {
			p1, p2 := path(x1, x2), path(y1, y2)

			for idx := range p1 {
				ventmap[p1[idx]][p2[idx]] = ventmap[p1[idx]][p2[idx]] + 1
			}
			// fmt.Println("not a horizontal line")
			continue
		}

		// horizontal
		if x1 != x2 {
			for _, dx := range path(x1, x2) {
				ventmap[dx][y1] = ventmap[dx][y1] + 1
			}
		}

		// vertical
		if y1 != y2 {
			for _, dy := range path(y1, y2) {
				ventmap[x1][dy] = ventmap[x1][dy] + 1
			}
		}
	}

	sum := 0
	for i := range ventmap {
		for j := range ventmap[i] {
			if ventmap[i][j] >= 2 {
				sum++
			}
		}
	}

	fmt.Printf("Part 1: intersections: %d\n", sum)
	// fmt.Println(ventmap)

}

func path(start, end int) (res []int) {
	// max := math.Max(float64(start), float64(end))
	// min := math.Min(float64(start), float64(end))
	// fmt.Printf("Path: [%d, %d] min(%d) max(%d)\n", start, end, int(min), int(max))
	// res := []int{}
	// for i := min; i <= max; i++ {
	// 	// fmt.Println(int(i))
	// 	res = append(res, int(i))
	// }

	if start < end {
		for i := start; i <= end; i++ {
			res = append(res, i)
		}
	} else {
		for i := start; i >= end; i-- {
			res = append(res, i)
		}
	}
	return res
}
