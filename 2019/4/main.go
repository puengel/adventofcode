package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("2019 #4")

	lower, upper := getRange()

	matches := findMatches(lower, upper)

	fmt.Printf("Sum matching: %d\n", matches)
}

func getRange() (lower, upper int) {
	return 245182, 790572
}

func findMatches(lower, upper int) (result int) {
	result = 0
	for num := lower; num <= upper; num++ {

		current := iToArray(num)

		// Rule 1: six digit number (by input)
		if len(current) != 6 {
			log.Println("What the fuck?")
			continue
		}
		// Rule 2: value within input (by for loop)

		// Rule 3: Two adjacent must be the same #Part 1
		// r3 := false
		// for c := 0; c < len(current)-1; c++ {
		// 	if current[c] == current[c+1] {
		// 		r3 = true
		// 		break
		// 	}
		// }

		// Rule 3: Two adjacent must be the same #Part 2
		r3 := false

		// borders
		if current[0] == current[1] && current[1] != current[2] {
			// fmt.Println(1)
			r3 = true
		}

		if current[5] == current[4] && current[4] != current[3] {
			// fmt.Println(2)
			r3 = true
		}

		// middle part
		for i, j := 1, 2; !r3 && j < 5; i, j = i+1, j+1 {
			if current[i-1] != current[i] &&
				current[i] == current[j] &&
				current[j] != current[j+1] {
				// fmt.Println(3)
				r3 = true
			}
		}

		if !r3 {
			// log.Println("dafuq?")
			continue
		}

		// Rule 4: never decrease
		r4 := true
		for c := 0; c < len(current)-1; c++ {
			if current[c] > current[c+1] {
				r4 = false
				break
			}
		}
		if !r4 {
			// log.Println("Comeon")
			continue
		}

		// fmt.Println(num)
		result++
	}

	return
}

func iToArray(input int) (result []int) {
	result = make([]int, 0)

	for {
		digit := input % 10
		result = append(result, digit)
		input = input / 10
		if input == 0 {
			break
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return
}
