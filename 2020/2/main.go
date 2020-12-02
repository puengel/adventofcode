package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Password struct {
	Lowest   int
	Highest  int
	Value    string
	Password string
}

func main() {

	lines, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var Passwords []Password
	var pass Password

	for _, line := range lines {
		sub := strings.Split(line, " ")
		borders := strings.Split(sub[0], "-")
		pass.Lowest, err = strconv.Atoi(borders[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		pass.Highest, err = strconv.Atoi(borders[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		pass.Value = strings.TrimSuffix(sub[1], ":")
		pass.Password = sub[2]

		Passwords = append(Passwords, pass)
	}

	sumValid := len(Passwords)

	for _, pw := range Passwords {

		num := strings.Count(pw.Password, pw.Value)
		if num < pw.Lowest || num > pw.Highest {
			// fmt.Printf("Got %d\t%s\tLow: %d\tHigh:%d\tPW: %s\n", num, pw.Value, pw.Lowest, pw.Highest, pw.Password)
			sumValid--
		}
	}

	fmt.Printf("Part 1: Sum valid PWs: %d\n", sumValid)

	sumValid = 0

	for _, pw := range Passwords {
		first := string(pw.Password[pw.Lowest-1])
		second := string(pw.Password[pw.Highest-1])
		if (first == pw.Value && second != pw.Value) || (first != pw.Value && second == pw.Value) {
			sumValid++
		}
	}

	fmt.Printf("Part 2: Sum valid PWs: %d\n", sumValid)

}
