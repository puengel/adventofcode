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

	// fmt.Printf("in: %+v\n", in)
	operate(in)

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

	// fmt.Print(output)
	finished := false
	for i := 0; !finished; {

		opcode := output[i]
		DE := opcode % 100
		C := (opcode / 100) % 10
		B := (opcode / 1000) % 10
		A := (opcode / 10000) % 10
		// fmt.Println(opcode)

		switch DE {
		case 1:
			// add
			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}
			p2 := output[i+2]
			if B == 0 {
				p2 = output[output[i+2]]
			}
			if A == 0 {
				output[output[i+3]] = p1 + p2
			} else {
				output[i+3] = p1 + p2
			}
			i += 4
		case 2:
			// multiply
			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}
			p2 := output[i+2]
			if B == 0 {
				p2 = output[output[i+2]]
			}
			if A == 0 {
				output[output[i+3]] = p1 * p2
			} else {
				output[i+3] = p1 * p2
			}
			i += 4
		case 3:
			// Store input

			var input int
			fmt.Print("Input: ")
			_, err := fmt.Scanf("%d", &input)
			if err != nil {
				return nil, err
			}

			if C == 0 {
				output[output[i+1]] = input
			} else {
				output[i+1] = input
			}
			i += 2
		case 4:
			// Write output

			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}
			fmt.Printf("%d\n", p1)
			i += 2

		case 5:
			// Jump if true
			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}

			if p1 != 0 {
				p2 := output[i+2]
				if B == 0 {
					p2 = output[output[i+2]]
				}
				i = p2
			} else {
				i += 3
			}
		case 6:
			// Jump if false
			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}

			if p1 == 0 {
				p2 := output[i+2]
				if B == 0 {
					p2 = output[output[i+2]]
				}
				i = p2
			} else {
				i += 3
			}

		case 7:
			// less than
			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}
			p2 := output[i+2]
			if B == 0 {
				p2 = output[output[i+2]]
			}
			p3 := output[i+3]
			// if A == 0 {
			// 	p3 = output[output[i+3]]
			// }
			if p1 < p2 {
				output[p3] = 1
			} else {
				output[p3] = 0
			}
			i += 4

		case 8:
			// equals
			p1 := output[i+1]
			if C == 0 {
				p1 = output[output[i+1]]
			}
			p2 := output[i+2]
			if B == 0 {
				p2 = output[output[i+2]]
			}
			p3 := output[i+3]
			// if A == 0 {
			// 	p3 = output[output[i+3]]
			// }
			if p1 == p2 {
				output[p3] = 1
				// fmt.Println("blubb")
			} else {
				output[p3] = 0
				// fmt.Println("oj")
			}
			// fmt.Println(output)
			i += 4

		case 99:
			finished = true
		default:
			return nil, fmt.Errorf("Bad opcode: %d", input[i])
		}
	}

	return output, nil
}
