package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	sRules := strings.Split(inputs[0], "\n")

	stripPrefix := regexp.MustCompile(`(\d+): `)
	checkNum := regexp.MustCompile(`\d`)

	rules := make(map[int]string)

	for i := range sRules {

		k, err := strconv.Atoi(strings.Split(sRules[i], ":")[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		// also pad
		rules[k] = strings.ReplaceAll(" "+strings.ReplaceAll(stripPrefix.ReplaceAllString(sRules[i], ""), "\"", "")+" ", " ", "  ")
		if strings.Contains(rules[k], "|") {
			rules[k] = "(" + rules[k] + ")"
		}

	}

	replacers := make(map[int]int)

	for k, v := range rules {
		if checkNum.FindString(v) == "" {
			replacers[k] = 1
		}
	}

	rule := rules[0]

	for checkNum.FindString(rules[0]) != "" {

		for num, used := range replacers {
			if used != 1 {
				continue
			}
			r := rules[num]
			find := fmt.Sprintf(" %d ", num)
			repl := fmt.Sprintf(" %s ", r)
			for k, v := range rules {
				if _, ok := replacers[k]; ok {
					continue
				}
				rules[k] = strings.ReplaceAll(v, find, repl)
				if checkNum.FindString(rules[k]) == "" {
					replacers[k] = 1
				}
			}

			replacers[num]++
		}

	}

	rule = strings.ReplaceAll(rules[0], " ", "")

	tester := regexp.MustCompile(rule)

	msgs := strings.Split(inputs[1], "\n")

	sum := 0

	for _, m := range msgs {
		if tester.FindString(m) == m {
			sum++
		}
	}

	fmt.Printf("Part 1: %d\n", sum)

	p2(inputs)
}

func p2(inputs []string) {

	rules := make(map[int]string)

	sRules := strings.Split(inputs[0], "\n")

	stripPrefix := regexp.MustCompile(`(\d+): `)
	checkNum := regexp.MustCompile(`\d`)

	for i := range sRules {

		k, err := strconv.Atoi(strings.Split(sRules[i], ":")[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		if k == 8 {
			// 8: 42 | 42 8
			rules[k] = "(  42  )+"
			continue
		}
		if k == 11 {
			// 11: 42 31 | 42 11 31
			rules[k] = "((  42  )(?R)?(  31  ))"
			continue
		}

		// also pad
		rules[k] = strings.ReplaceAll(" "+strings.ReplaceAll(stripPrefix.ReplaceAllString(sRules[i], ""), "\"", "")+" ", " ", "  ")
		if strings.Contains(rules[k], "|") {
			rules[k] = "(" + rules[k] + ")"
		}

	}

	replacers := make(map[int]int)

	for k, v := range rules {
		if checkNum.FindString(v) == "" {
			replacers[k] = 1
		}
	}

	rule := rules[0]

	for checkNum.FindString(rules[0]) != "" {

		for num, used := range replacers {
			if used != 1 {
				continue
			}
			r := rules[num]
			find := fmt.Sprintf(" %d ", num)
			repl := fmt.Sprintf(" %s ", r)
			for k, v := range rules {
				if _, ok := replacers[k]; ok {
					continue
				}
				rules[k] = strings.ReplaceAll(v, find, repl)
				if checkNum.FindString(rules[k]) == "" {
					replacers[k] = 1
				}
			}

			replacers[num]++
		}

	}

	rule = strings.ReplaceAll(rules[0], " ", "")

	fmt.Println(rule)

	// tester = regexp.MustCompilePOSIX(rule)
	// regexp.

	msgs := strings.Split(inputs[1], "\n")

	sum := 0

	for _, m := range msgs {
		// finds := tester.FindAllString(m, -1)
		// fmt.Println(finds)
		cmd := exec.Cmd{
			Path:   "./regex.pl",
			Args:   []string{"./regex.pl", rule, m},
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		if err := cmd.Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				sum += exitError.ExitCode()
			}
		}

		// for _, f := range finds {
		// 	fmt.Println(f)
		// 	if f == m {
		// 		sum++
		// 	}
		// }
	}

	fmt.Printf("Part 2: %d\n", sum)
}
