package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"sort"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	// for i, rune := range "({[<>]})" {
	// 	fmt.Printf("%d: %c\n", i, rune)
	// }
	incomepletes := []int{}

	score := 0
	for _, line := range inputs {
		puffer := []string{}
		corrupted := false
		for _, r := range line {

			rs := string(r)

			switch rs {
			case "(":
				fallthrough
			case "{":
				fallthrough
			case "[":
				fallthrough
			case "<":
				puffer = append(puffer, rs)
			case ")":
				fallthrough
			case "}":
				fallthrough
			case "]":
				fallthrough
			case ">":
				if len(puffer) > 0 && isCounterpart(puffer[len(puffer)-1], rs) {
					puffer = puffer[:len(puffer)-1]
				} else {
					// fmt.Println("Bad line", line[:idx], puffer, rs)
					score += errorScore(rs)
					corrupted = true
				}
			}

			if corrupted {
				break
			}
		}

		if !corrupted && len(puffer) > 0 {
			completer := []string{}
			for len(puffer) > 0 {
				switch puffer[len(puffer)-1] {
				case "(":
					completer = append(completer, ")")
				case "{":
					completer = append(completer, "}")
				case "[":
					completer = append(completer, "]")
				case "<":
					completer = append(completer, ">")
				}
				puffer = puffer[:len(puffer)-1]
			}

			cScore := completionScore(completer)

			incomepletes = append(incomepletes, cScore)
		}
	}

	sort.Ints(incomepletes)

	fmt.Printf("Part 1: Errorscore %d\n", score)
	fmt.Printf("Part 2: Completeion middle score %d\n", incomepletes[len(incomepletes)/2])
}

func isCounterpart(open, close string) bool {
	switch open {
	case "(":
		return close == ")"
	case "{":
		return close == "}"
	case "[":
		return close == "]"
	case "<":
		return close == ">"
	}

	return false
}

func errorScore(err string) int {
	switch err {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}

	return 0
}

func completionScore(completer []string) int {
	score := 0

	for _, val := range completer {
		score *= 5
		switch val {
		case ")":
			score += 1
		case "]":
			score += 2
		case "}":
			score += 3
		case ">":
			score += 4
		}
	}

	return score
}
