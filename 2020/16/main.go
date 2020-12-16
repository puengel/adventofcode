package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type condition struct {
	idx    int
	range1 condRange
	range2 condRange
}

type condRange struct {
	upper int
	lower int
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	reg := regexp.MustCompile(`.+:\s(.*)`)

	conditionsStr := strings.Split(inputs[0], "\n")
	conditions := make([]condition, len(conditionsStr))

	for i, cond := range conditionsStr {
		res := reg.ReplaceAllString(cond, "${1}")
		ranges := strings.Split(res, " or ")
		r1 := strings.Split(ranges[0], "-")
		lo1, err := strconv.Atoi(r1[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		up1, err := strconv.Atoi(r1[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		r2 := strings.Split(ranges[1], "-")
		lo2, err := strconv.Atoi(r2[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		up2, err := strconv.Atoi(r2[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		c := condition{
			idx: -1,
			range1: condRange{
				lower: lo1,
				upper: up1,
			},
			range2: condRange{
				lower: lo2,
				upper: up2,
			},
		}
		conditions[i] = c
		// fmt.Println(c)
	}

	errRate := 0

	var validTickets []string

	tickets := strings.Split(inputs[2], "\n")[1:]
	for _, ticket := range tickets {
		fields := strings.Split(ticket, ",")

		allFieldsOK := true
		for _, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				return
			}

			var matchesCondition bool
			for _, c := range conditions {
				if (val >= c.range1.lower && val <= c.range1.upper) || (val >= c.range2.lower && val <= c.range2.upper) {
					matchesCondition = true
					break
				}
			}

			if !matchesCondition {
				allFieldsOK = false
				errRate += val
			}

		}

		if allFieldsOK {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Printf("Part 1: %d\n", errRate)

	found := make(map[int]bool)
	for len(found) != len(conditions) {

		for cIdx, c := range conditions {
			if c.idx != -1 {
				continue
			}
			var possible []int

			for i := 0; i < len(conditions); i++ {
				if found[i] {
					continue
				}

				matchesCondition := true

				for _, ticket := range validTickets {
					fields := strings.Split(ticket, ",")

					val, err := strconv.Atoi(fields[i])
					if err != nil {
						fmt.Println(err)
						return
					}
					if !((val >= c.range1.lower && val <= c.range1.upper) || (val >= c.range2.lower && val <= c.range2.upper)) {
						matchesCondition = false
						break
					}
				}

				if matchesCondition {
					possible = append(possible, i)
				}
			}

			if len(possible) == 1 {
				conditions[cIdx].idx = possible[0]
				found[possible[0]] = true
				break
			}

		}
	}

	myTicket := strings.Split(strings.Split(inputs[1], "\n")[1], ",")
	res := 1
	for i := 0; i < 6; i++ {
		val, err := strconv.Atoi(myTicket[conditions[i].idx])
		if err != nil {
			fmt.Println(err)
			return
		}
		res *= val
	}

	fmt.Printf("Part 2: %d\n", res)
}
