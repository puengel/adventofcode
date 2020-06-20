package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	vals, err := readValues()
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("Vals: %+v\n", vals)

	// part1
	// vals = programAlarm(vals)

	// fmt.Printf("Vals: %+v\n", vals)

	var noun int
	var verb int

	done := false

	for noun = 0; noun <= 99 && !done; noun++ {
		for verb = 0; verb <= 99 && !done; verb++ {

			current := make([]int, len(vals))
			copy(current, vals)

			current = nounVerb(noun, verb, current)

			current, err = operate(current)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if current[0] == 19690720 {
				copy(vals, current)
				done = true
			}

		}
	}

	noun--
	verb--

	fmt.Printf("Vals: %+v\n", vals)
	fmt.Printf("Noun: %d\nVerb: %d\n", noun, verb)
	fmt.Printf("Result: %d\n", 100*noun+verb)
}

func readValues() (vals []int, err error) {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	sVals := strings.Split(string(file), ",")

	for _, v := range sVals {
		iVal, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		vals = append(vals, iVal)
	}

	return vals, nil
}

func programAlarm(input []int) []int {
	input[1] = 12
	input[2] = 2

	return input
}

func nounVerb(noun, verb int, input []int) []int {
	input[1] = noun
	input[2] = verb

	return input
}

func operate(input []int) (output []int, err error) {
	output = make([]int, len(input))
	copy(output, input)

	fmt.Print(output)
	finished := false
	for i := 0; i < len(output) && !finished; i += 4 {

		switch output[i] {
		case 1:
			// add
			output[output[i+3]] = output[output[i+1]] + output[output[i+2]]
		case 2:
			// multiply
			output[output[i+3]] = output[output[i+1]] * output[output[i+2]]
		case 99:
			finished = true
		default:
			return nil, fmt.Errorf("Bad opcode: %d", input[i])
		}
	}

	return output, nil
}
