package main

import (
	"adventofcode/2020/pkg/accumulator"
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	program := accumulator.ParseProgram(inputs)

	result := program.RunUntilLoop()

	fmt.Printf("Part 1: %d\n", result)

	for idx := range program.Code {
		accumulator.ToggleJMPNOP(idx, &program.Code)

		result = 0

		result, err = program.RunUntilLast()
		if err == nil {
			break
		}

		accumulator.ToggleJMPNOP(idx, &program.Code)
	}

	fmt.Printf("Part 2: %d\n", result)

}
