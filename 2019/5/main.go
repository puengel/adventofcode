package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("2019 #5")

	in, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("in: %+v\n", in)

}

func readInput() ([]int, error) {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	sVars := strings.Split(string(file), ",")

	iVars := make([]int, len(sVars))

	for idx, val := range sVars {
		iVars[idx], err = strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
	}

	return iVars, err
}



func operate(input []int) (output []int, err error) {
	output = make([]int, len(input))
	copy(output, input)

	fmt.Print(output)
	finished := false
	for i := 0; !finished {

		switch output[i] {
		case 1:
			// add
			output[output[i+3]] = output[output[i+1]] + output[output[i+2]]
			i += 4
		case 2:
			// multiply
			output[output[i+3]] = output[output[i+1]] * output[output[i+2]]
			i += 4
		case 3:

		case 4:

		case 99:
			finished = true
		default:
			return nil, fmt.Errorf("Bad opcode: %d", input[i])
		}
	}

	return output, nil
}
