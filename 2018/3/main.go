package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

type Claim struct {
	left   int
	top    int
	width  int
	height int
	id     string
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, in := range inputs {

	}
}

func NewClaim(in string) *Claim {

}
