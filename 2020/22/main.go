package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

var game int

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	pStr := strings.Split(inputs[0], "\n")[1:]
	p1 := make([]int, len(pStr))
	_p1 := make([]int, len(pStr))
	for i, v := range pStr {
		a, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		p1[i] = a
	}
	copy(_p1, p1)
	pStr = strings.Split(inputs[1], "\n")[1:]
	p2 := make([]int, len(pStr))
	_p2 := make([]int, len(pStr))
	for i, v := range pStr {
		a, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		p2[i] = a
	}
	copy(_p2, p2)

	var a int
	var b int
	for 0 < len(p1) && 0 < len(p2) {
		a, p1 = p1[0], p1[1:]
		b, p2 = p2[0], p2[1:]
		if a < b {
			p2 = append(p2, b, a)
		} else {
			p1 = append(p1, a, b)
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)

	winner := p1
	if len(p2) > 0 {
		winner = p2
	}

	res := 0
	for i, val := range winner {
		res += val * (len(winner) - i)
	}

	fmt.Printf("Part 1: %d\n", res)

	game = 0
	res1, res2 := RecursiveCombat(_p1, _p2)

	winner = res2
	if len(res1) > 0 {
		winner = res1
	}

	res = 0
	for i, val := range winner {
		res += val * (len(winner) - i)
	}

	fmt.Printf("Part 2: %d\n", res)
}

func RecursiveCombat(p1deck, p2deck []int) (resP1, resP2 []int) {

	p1 := make([]int, len(p1deck))
	copy(p1, p1deck)
	p2 := make([]int, len(p2deck))
	copy(p2, p2deck)

	orders := make(map[string]bool)

	var a int
	var b int

	// game++
	// gameInner := game
	// round := 0

	for 0 < len(p1) && 0 < len(p2) {
		// round++
		// fmt.Printf("Game\t%d\tRound\t%d:\n", gameInner, round)

		// Player 1 wins if order has been there already
		oP1 := fmt.Sprintf("p1:%v", p1)
		oP2 := fmt.Sprintf("p2:%v", p2)
		if orders[oP1] || orders[oP2] {
			// fmt.Printf("Had that order already. P1 wins game.\n")
			return p1, p2
		}
		orders[oP1] = true
		orders[oP2] = true

		a, p1 = p1[0], p1[1:]
		b, p2 = p2[0], p2[1:]
		if len(p1) < a || len(p2) < b {
			if a < b {
				p2 = append(p2, b, a)
			} else {
				p1 = append(p1, a, b)
			}
		} else {
			// fmt.Printf("P1:\tcard[%d]\tDeck%v\tLen[%d]\nP2:\tcard[%d]\tDeck%v\tLen[%d]\nStart recursive\n\n", a, p1, len(p1), b, p2, len(p2))
			r1, _ := RecursiveCombat(p1, p2)
			if len(r1) > 0 {
				p1 = append(p1, a, b)
			} else {
				p2 = append(p2, b, a)
			}
		}

		// fmt.Printf("%v\n%v\n\n", p1, p2)
	}

	return p1, p2
}
