package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"strconv"
	"strings"

	"github.com/sanity-io/litter"
)

func main() {

	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	// Part 1 + 2
	horizontal := 0
	depth := 0
	for _, i := range inputs {
		lines := strings.Split(i, " ")
		if len(lines) != 2 {
			fmt.Println("meh")
			return
		}
		command, valueStr := lines[0], lines[1]

		val, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("bad")
			return
		}

		switch command {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		default:
			fmt.Println("bad command")
			return
		}
	}

	fmt.Printf("Part 1 horizontal: %d\tdepth: %d\tmult: %d\n", horizontal, depth, horizontal*depth)

	// Part 2
	horizontal = 0
	depth = 0
	aim := 0
	for _, i := range inputs {
		lines := strings.Split(i, " ")
		if len(lines) != 2 {
			fmt.Println("meh")
			return
		}
		command, valueStr := lines[0], lines[1]

		val, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Println("bad")
			return
		}

		switch command {
		case "forward":
			horizontal += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		default:
			fmt.Println("bad command")
			return
		}
	}

	fmt.Printf("Part 2 horizontal: %d\tdepth: %d\tmult: %d\n", horizontal, depth, horizontal*depth)

	type Command struct {
		Action string
		Dir    int
	}

	c := []Command{}

	// fmt.Println(reflect.TypeOf(&c).String())

	fmt.Printf("Ptr: %p\n", &c)

	err = input.ReflectAll(inputs[:2], " ", &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Ptr: %p\n", &c)
	litter.Dump(c)
}
