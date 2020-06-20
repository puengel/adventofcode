package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2019 #1")

	fuel := p1()

	fmt.Printf("All Fuel: %d\n", fuel)
}

func p1() int {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return 0
	}

	// fmt.Printf("%s", file)

	a := (14 / 3) - 2
	fmt.Println(a)

	lines := strings.Split(string(file), "\n")

	sum := 0

	for _, line := range lines {

		// fmt.Printf("%s\n", line)

		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Print("Noooo")
			return 0
		}

		fuel := (mass / 3) - 2

		allfuel := p2(fuel)

		sum += allfuel
	}

	fmt.Printf("Sum: %d\n", sum)

	return sum
}

func p2(fuel int) int {

	if fuel < 9 {
		return fuel
	}

	add := (fuel / 3) - 2

	return fuel + p2(add)
}
