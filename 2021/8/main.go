package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"strings"
)

func main() {
	lines, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	sum := 0
	for _, line := range lines {
		outputs := strings.Split(line, " | ")
		for _, output := range strings.Split(outputs[1], " ") {
			if len(output) <= 4 || len(output) == 7 {
				sum++
			}
		}
	}

	fmt.Printf("Part 1: sum(1,4,7,8): %d\n", sum)

	sumAllOuts := 0
	for _, line := range lines {
		// find 1, 4, 7
		inout := strings.Split(line, " | ")
		ins := strings.Split(inout[0], " ")
		outs := strings.Split(inout[1], " ")

		one := ""
		four := ""
		seven := ""

		for _, val := range ins {
			switch len(val) {
			case 2:
				one = val
			case 4:
				four = val
			case 3:
				seven = val
			}
		}

		right := []rune{}      // from 1
		top := []rune{}        // from 7 - 1
		centerLeft := []rune{} // from 4 - 1

		for _, oneRune := range one {
			right = append(right, oneRune)
		}

		for _, sevenRune := range seven {
			for _, val := range right {
				if val != sevenRune {
					top = append(top, sevenRune)
				}
			}
		}

		for _, fourRune := range four {
			inRight := false
			for _, val := range right {
				if val == fourRune {
					inRight = true
				}
			}
			if !inRight {
				centerLeft = append(centerLeft, fourRune)
			}
		}

		// get output
		mult := 1000
		combinedOut := 0
		for _, out := range outs {
			outNum := -1
			switch len(out) {
			case 2:
				outNum = 1
			case 4:
				outNum = 4
			case 3:
				outNum = 7
			case 7:
				outNum = 8
			case 5:
				inRight := allIn(right, []rune(out))
				inCenter := allIn(centerLeft, []rune(out))
				if !inRight && !inCenter {
					outNum = 2
				} else if !inRight && inCenter {
					outNum = 5
				} else if inRight && !inCenter {
					outNum = 3
				}
			case 6:
				inRight := allIn(right, []rune(out))
				inCenter := allIn(centerLeft, []rune(out))
				if inRight && inCenter {
					outNum = 9
				} else if !inRight && inCenter {
					outNum = 6
				} else if inRight && !inCenter {
					outNum = 0
				}
			}

			if outNum == -1 {
				fmt.Println("not found", out)
			}

			combinedOut += outNum * mult
			mult = mult / 10
		}

		sumAllOuts += combinedOut
	}

	fmt.Printf("Part 2: %d\n", sumAllOuts)
}

func allIn(target, test []rune) bool {
	for _, targetVal := range target {
		found := false
		for _, testVal := range test {
			if testVal == targetVal {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
