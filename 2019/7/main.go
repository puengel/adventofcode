package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type powerLine struct {
	amps    []amplifier
	program []int
}

type amplifier interface {
	setInput(in int)
	getOutput() (out int)
	operate() error
	setup(int, []int)
}

type amp struct {
	phase   int
	input   int
	output  int
	program []int
}

func main() {
	fmt.Println("2019 #7")

	program, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	line := powerLine{
		amps:    make([]amplifier, 0),
		program: program,
	}

	// create amp array
	for i := 0; i < 5; i++ {

		a := &amp{
			input:   0,
			output:  0,
			phase:   0,
			program: program,
		}

		line.amps = append(line.amps, a)

	}

	phases := []int{0, 1, 2, 3, 4}

	perms := permutations(phases)

	max := 0

	for _, perm := range perms {

		fmt.Printf("Perm: %+v\n", perm)

		line.setup(perm)
		result := line.run()

		fmt.Printf("Result: %d\n", result)

		if result > max {
			max = result
		}
	}

	fmt.Printf("Max: %d\n", max)
}

func (l *powerLine) setup(shape []int) {
	for idx, amp := range l.amps {
		amp.setup(shape[idx], l.program)
	}
}

func (l *powerLine) run() int {

	passValue := 0
	for _, amp := range l.amps {
		amp.setInput(passValue)
		err := amp.operate()
		if err != nil {
			fmt.Println("FCK")
			passValue = 0
			break
		}
		passValue = amp.getOutput()
		fmt.Println(passValue)
	}

	return passValue
}

func readInput() ([]int, error) {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}

	// fmt.Println(file)

	splits := strings.Split(string(file), ",")
	result := make([]int, len(splits))
	for idx, val := range splits {
		result[idx], err = strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (a *amp) setup(phase int, program []int) {
	a.phase = phase

	// make a copy
	p := make([]int, len(program))
	copy(p, program)
	a.program = p
}

func (a *amp) operate() (err error) {
	program := make([]int, len(a.program))
	copy(program, a.program)
	input := []int{a.phase, a.input}
	fmt.Println(input)
	inputPtr := 0
	output := 0

	// fmt.Print(output)
	finished := false
	for i := 0; !finished; {

		opcode := program[i]
		DE := opcode % 100
		C := (opcode / 100) % 10
		B := (opcode / 1000) % 10
		A := (opcode / 10000) % 10
		// fmt.Println(opcode)

		switch DE {
		case 1:
			// add
			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}
			p2 := program[i+2]
			if B == 0 {
				p2 = program[program[i+2]]
			}
			if A == 0 {
				program[program[i+3]] = p1 + p2
			} else {
				program[i+3] = p1 + p2
			}
			i += 4
		case 2:
			// multiply
			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}
			p2 := program[i+2]
			if B == 0 {
				p2 = program[program[i+2]]
			}
			if A == 0 {
				program[program[i+3]] = p1 * p2
			} else {
				program[i+3] = p1 * p2
			}
			i += 4
		case 3:
			// Store input

			// check error
			if inputPtr > 1 {
				fmt.Println("Too many inputs asked")
				return errors.New("Program asked for too many inputs")
			}

			if C == 0 {
				program[program[i+1]] = input[inputPtr]
			} else {
				program[i+1] = input[inputPtr]
			}
			inputPtr++
			i += 2
		case 4:
			// Write output

			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}
			// fmt.Printf("%d\n", p1)
			output = p1
			i += 2

		case 5:
			// Jump if true
			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}

			if p1 != 0 {
				p2 := program[i+2]
				if B == 0 {
					p2 = program[program[i+2]]
				}
				i = p2
			} else {
				i += 3
			}
		case 6:
			// Jump if false
			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}

			if p1 == 0 {
				p2 := program[i+2]
				if B == 0 {
					p2 = program[program[i+2]]
				}
				i = p2
			} else {
				i += 3
			}

		case 7:
			// less than
			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}
			p2 := program[i+2]
			if B == 0 {
				p2 = program[program[i+2]]
			}
			p3 := program[i+3]
			// if A == 0 {
			// 	p3 = output[output[i+3]]
			// }
			if p1 < p2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			i += 4

		case 8:
			// equals
			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}
			p2 := program[i+2]
			if B == 0 {
				p2 = program[program[i+2]]
			}
			p3 := program[i+3]
			// if A == 0 {
			// 	p3 = output[output[i+3]]
			// }
			if p1 == p2 {
				program[p3] = 1
				// fmt.Println("blubb")
			} else {
				program[p3] = 0
				// fmt.Println("oj")
			}
			// fmt.Println(output)
			i += 4

		case 99:
			finished = true
		default:
			return fmt.Errorf("Bad opcode: %d", program[i])
		}
	}

	// fmt.Println(output)
	a.output = output

	return nil
}

func (a *amp) setInput(in int) {
	a.input = in
}

func (a *amp) getOutput() (out int) {
	out = a.output
	return
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
