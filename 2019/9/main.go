package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type icc struct {
	in        chan int
	fromStdIn bool
	out       chan int
	halt      chan int
	program   []int
}

func main() {
	fmt.Println("2019 #9")

	program, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	in := make(chan int, 5)
	out := make(chan int, 5)
	halt := make(chan int, 5)

	comp := icc{
		in:        in,
		out:       out,
		halt:      halt,
		fromStdIn: true,
		program:   program,
	}

	err = comp.operate()
	if err != nil {
		fmt.Println(err)
	}

	// stopped := false
	// for !stopped {
	// 	select {
	// 	case val := <-out:
	// 		fmt.Println(val)
	// 	case <-halt:
	// 		stopped = true
	// 	default:
	// 		// fmt.Println("Nothing")
	// 	}
	// }
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

func (a *icc) operate() (err error) {
	// program := make([]int, len(a.program))
	// copy(program, a.program)

	program := a.program
	// input := []int{a.phase, a.input}
	// fmt.Println(input)
	inputPtr := 0
	relativeBase := 0

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
			p1 := a.getParam(i+1, C, relativeBase)
			p2 := a.getParam(i+2, B, relativeBase)
			a.writeTo(i+3, A, relativeBase, p1+p2)

			i += 4
		case 2:
			// multiply
			p1 := a.getParam(i+1, C, relativeBase)
			p2 := a.getParam(i+2, B, relativeBase)
			a.writeTo(i+3, A, relativeBase, p1*p2)

			i += 4
		case 3:
			// Store input

			// check error
			if inputPtr > 1 {
				fmt.Println("Too many inputs asked")
				return errors.New("Program asked for too many inputs")
			}

			if a.fromStdIn {
				var val int
				fmt.Print("Input: ")
				_, err := fmt.Scanf("%d", &val)
				if err != nil {
					return err
				}
				a.in <- val
			}

			a.writeTo(i+1, C, relativeBase, <-a.in)

			i += 2
		case 4:
			// Write output

			p1 := a.getParam(i+1, C, relativeBase)
			// fmt.Printf("%d\n", p1)
			a.out <- p1
			i += 2

		case 5:
			// Jump if true
			p1 := a.getParam(i+1, C, relativeBase)

			if p1 != 0 {
				p2 := a.getParam(i+2, B, relativeBase)
				i = p2
			} else {
				i += 3
			}
		case 6:
			// Jump if false
			p1 := a.getParam(i+1, C, relativeBase)

			if p1 == 0 {
				p2 := a.getParam(i+2, B, relativeBase)
				i = p2
			} else {
				i += 3
			}

		case 7:
			// less than
			p1 := a.getParam(i+1, C, relativeBase)
			p2 := a.getParam(i+2, B, relativeBase)
			p3 := program[i+3]
			// if A == 0 {
			// 	p3 = output[output[i+3]]
			// }
			a.verifySize(p3)
			if p1 < p2 {
				program[p3] = 1
			} else {
				program[p3] = 0
			}
			i += 4

		case 8:
			// equals
			p1 := a.getParam(i+1, C, relativeBase)
			p2 := a.getParam(i+2, B, relativeBase)
			p3 := program[i+3]
			// if A == 0 {
			// 	p3 = output[output[i+3]]
			// }
			a.verifySize(p3)
			if p1 == p2 {
				program[p3] = 1
				// fmt.Println("blubb")
			} else {
				program[p3] = 0
				// fmt.Println("oj")
			}
			// fmt.Println(output)
			i += 4

		case 9:
			p1 := a.getParam(i+1, C, relativeBase)
			relativeBase += p1

			i += 2

		case 99:
			finished = true
			a.halt <- 1
		default:
			fmt.Printf("Bad opcode: %d\n", program[i])
			finished = true
			a.halt <- -1
		}
	}

	// fmt.Println(output)
	// a.output = output

	return nil
}

func (a *icc) writeTo(ptr, mode, base, value int) {
	switch mode {
	case 0:
		ptr = a.program[ptr]
	case 1:
	case 2:
		ptr = a.program[ptr] + base
	default:
		fmt.Printf("Bad mode: %d\n", mode)
	}
	a.verifySize(ptr)
	a.program[ptr] = value
}

func (a *icc) getParam(ptr, mode, base int) (p int) {
	a.verifySize(ptr)
	program := a.program

	switch mode {
	case 0:
		p = program[ptr]
	case 1:
		p = program[program[ptr]]
	case 2:
		p = program[base+program[ptr]]
	default:
		fmt.Printf("Bad mode: %d\n", mode)
	}

	return
}

func (a *icc) verifySize(ptr int) {
	if ptr >= len(a.program) {
		tmp := make([]int, ptr+1)
		copy(tmp, a.program)
		a.program = tmp
	}
}
