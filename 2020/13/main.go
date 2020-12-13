package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	ts, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var ids []int
	var Ax []int64

	for idx, in := range strings.Split(lines[1], ",") {
		if in == "x" {
			continue
		}
		num, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		ids = append(ids, num)

		Ax = append(Ax, int64(idx+1))
	}

	var wait int
	var shortest int

	for _, num := range ids {
		res := num
		for res < ts {
			res += num
		}

		diff := res - ts
		if diff < wait || shortest == 0 {
			wait = diff
			shortest = num
		}
	}

	fmt.Printf("Part 1: %d\n", wait*shortest)

	// Ids: Mod classes
	// Ax: Rest classes

	// (i)
	N := int64(1)
	for _, mX := range ids {
		N *= int64(mX)
	}

	// (ii)
	var Nx []int64
	for _, mX := range ids {
		Nx = append(Nx, N/int64(mX))
	}

	// (iii)
	var X []int64
	for idx := range Nx {
		for test := 0; test < ids[idx]; test++ {
			if (int64(test)*Nx[idx])%int64(ids[idx]) == int64(1%ids[idx]) {
				X = append(X, int64(test))
				break
			}
		}
	}
	// (iv)
	var res int64
	for idx := range ids {
		fmt.Printf("%d * %d * %d\n", Ax[idx], X[idx], Nx[idx])
		res += -Ax[idx] * X[idx] * Nx[idx]
		res = res % N
	}

	res = res%N + N

	fmt.Printf("Part 2: %d\n", res+1)
}
