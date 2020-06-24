package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type powerLine struct {
	amps      []amplifier
	masterIn  chan int
	masterOut chan int
	halt      chan int
	program   []int
}

type amplifier interface {
	getInput() (in chan int)
	setOutput(out chan int)
	operate() error
	setup(int, []int)
	setHalt(chan int)
}

type amp struct {
	in      chan int
	out     chan int
	halt    chan int
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
		amps:      make([]amplifier, 0),
		halt:      make(chan int),
		masterOut: make(chan int),
		program:   program,
	}

	// create amp array
	for i := 0; i < 5; i++ {

		a := &amp{
			in:      make(chan int, 2),
			program: program,
		}

		line.amps = append(line.amps, a)

	}

	phases := []int{5, 6, 7, 8, 9}

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

func (l *powerLine) setup(shape []int) {
	for idx, amp := range l.amps {
		amp.setup(shape[idx], l.program)
		amp.setHalt(l.halt)
		if idx == 0 {
			l.masterIn = amp.getInput()
		}
		if idx == len(l.amps)-1 {
			amp.setOutput(l.masterOut)
		} else {
			amp.setOutput(l.amps[idx+1].getInput())
		}
	}
}

func (a *amp) setup(phase int, program []int) {

	// flush channels
	flushing := true
	for flushing {
		select {
		case <-a.halt:
		case <-a.in:
		case <-a.out:
		default:
			flushing = false
		}
	}

	// a.phase = phase
	a.in <- phase

	// make a copy
	p := make([]int, len(program))
	copy(p, program)
	a.program = p
}

func (l *powerLine) run() int {

	for _, amp := range l.amps {
		go amp.operate()
	}

	// start
	l.masterIn <- 0

	last := 0
	halts := 0

	done := false

	for !done {

		select {
		case out := <-l.masterOut:
			last = out
			l.masterIn <- out
		case h := <-l.halt:
			halts += h
			if halts >= 5 {
				done = true
			}
		default:

		}

	}

	return last
}

func (a *amp) operate() (err error) {
	program := make([]int, len(a.program))
	copy(program, a.program)
	// input := []int{a.phase, a.input}
	// fmt.Println(input)
	inputPtr := 0

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
				program[program[i+1]] = <-a.in
			} else {
				program[i+1] = <-a.in
			}
			i += 2
		case 4:
			// Write output

			p1 := program[i+1]
			if C == 0 {
				p1 = program[program[i+1]]
			}
			// fmt.Printf("%d\n", p1)
			a.out <- p1
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
			a.halt <- 1
		default:
			return fmt.Errorf("Bad opcode: %d", program[i])
		}
	}

	// fmt.Println(output)
	// a.output = output

	return nil
}

func (a *amp) getInput() chan int {
	return a.in
}

func (a *amp) setOutput(out chan int) {
	a.out = out
}

func (a *amp) setHalt(halt chan int) {
	a.halt = halt
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
