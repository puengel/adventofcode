package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"

	"github.com/ernestosuarez/itertools"
)

func main() {
	numbers, err := input.ReadInts("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int

	for idx, number := range numbers {
		// skip first 25
		if idx < 25 {
			continue
		}

		if !FindSum(number, numbers[idx-25:idx]) {
			result = number
			break
		}
	}

	fmt.Printf("Part 1: %d\n", result)

	p2 := FindContiguousSum(result, numbers)

	fmt.Printf("Part 2: %d\n", p2)
}

func FindSum(sum int, list []int) bool {
	// must be length 25
	if len(list) != 25 {
		fmt.Println("need 25 vals in array")
		return false
	}

	for v := range itertools.CombinationsInt(list, 2) {
		if (v[0] + v[1]) == sum {
			// fmt.Printf("%d + %d = %d\n", v[0], v[1], sum)
			return true
		}
	}

	return false
}

func FindContiguousSum(sum int, list []int) int {

	for idx, a := range list {
		check := a
		for i := idx + 1; i < len(list); i++ {
			check += list[i]
			if check > sum {
				break
			}
			if check == sum {
				min, max := MinMax(list[idx : i+1])
				return min + max
			}
		}
	}

	return -1
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
