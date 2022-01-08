package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"strconv"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	var gamma, epsilon string

	numbers := len(inputs[0])

	for i := 0; i < numbers; i++ {
		zeros := 0
		ones := 1
		for _, val := range inputs {
			if string(val[i]) == "1" {
				ones++
			} else {
				zeros++
			}
		}

		if zeros > ones {
			gamma = fmt.Sprintf("%s%s", gamma, "0")
			epsilon = fmt.Sprintf("%s%s", epsilon, "1")
		} else {
			gamma = fmt.Sprintf("%s%s", gamma, "1")
			epsilon = fmt.Sprintf("%s%s", epsilon, "0")
		}
	}

	fmt.Println(gamma, epsilon)

	gNum, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return
	}

	eNum, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return
	}

	fmt.Printf("Part 1: gamma: %d\tepsilon: %d\tmult: %d\n", gNum, eNum, gNum*eNum)

	// PART 2
	oxyInputs := make([]string, len(inputs), cap(inputs))
	copy(oxyInputs, inputs)
	for i := 0; len(oxyInputs) > 1; i = (i + 1) % numbers {
		// find next most common
		zeros := 0
		ones := 0
		for _, val := range oxyInputs {
			if string(val[i]) == "1" {
				ones++
			} else {
				zeros++
			}
		}
		keep := "1"
		if zeros > ones {
			keep = "0"
		}

		newOxyInputs := []string{}
		for _, val := range oxyInputs {
			if string(val[i]) == keep {
				newOxyInputs = append(newOxyInputs, val)
			}
		}
		oxyInputs = newOxyInputs
	}

	co2Inputs := make([]string, len(inputs), cap(inputs))
	copy(co2Inputs, inputs)
	for i := 0; len(co2Inputs) > 1; i = (i + 1) % numbers {
		// find next most common
		zeros := 0
		ones := 0
		for _, val := range co2Inputs {
			if string(val[i]) == "1" {
				ones++
			} else {
				zeros++
			}
		}
		keep := "0"
		if zeros > ones {
			keep = "1"
		}

		newCo2Inputs := []string{}
		for _, val := range co2Inputs {
			if string(val[i]) == keep {
				newCo2Inputs = append(newCo2Inputs, val)
			}
		}
		co2Inputs = newCo2Inputs
	}

	oxyS := oxyInputs[0]
	co2S := co2Inputs[0]
	fmt.Println(oxyS, co2S)

	oxy, err := strconv.ParseInt(oxyS, 2, 64)
	if err != nil {
		return
	}

	co2, err := strconv.ParseInt(co2S, 2, 64)
	if err != nil {
		return
	}

	fmt.Printf("Part 2: oxygen: %d\tco2: %d\tmult: %d\n", oxy, co2, oxy*co2)
}
