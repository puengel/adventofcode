package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Color    string
	Contains map[string]int
}

func main() {
	inputs, err := input.ReadStrings("input.txt", ".\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var rules []*Rule

	for _, in := range inputs {
		rules = append(rules, NewRule(in))
	}

	result := make(map[string]int)

	rule := FindRule("shiny gold", rules)
	if rule == nil {
		fmt.Println("that's bad")
		return
	}
	rule.FindParents(rules, &result)

	fmt.Printf("Part 1: %d\n", len(result))

	p2 := rule.NumChildren(rules)

	fmt.Printf("Part 2: %d\n", p2)

}

func NewRule(in string) *Rule {
	splits := strings.Split(in, " bags contain ")
	containRules := strings.Split(splits[1], ", ")
	contains := make(map[string]int)

	for _, co := range containRules {
		if co == "no other bags" || co == "no other bags." {
			break
		}
		ca := strings.SplitN(co, " ", 2)
		num, err := strconv.Atoi(ca[0])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		strip1 := strings.TrimSuffix(ca[1], " bags")
		color := strings.TrimSuffix(strip1, " bag")
		contains[color] = num
	}

	return &Rule{
		Color:    splits[0],
		Contains: contains,
	}
}

func FindRule(color string, rules []*Rule) *Rule {

	for _, rule := range rules {
		if rule.Color == color {
			return rule
		}
	}

	return nil
}

func (r *Rule) CanHold(color string) bool {
	_, ok := r.Contains[color]
	return ok
}

func (r *Rule) FindParents(rules []*Rule, result *map[string]int) {

	for _, rule := range rules {
		if rule.CanHold(r.Color) {
			(*result)[rule.Color] = 1
			rule.FindParents(rules, result)
		}
	}

	return
}

func (r *Rule) NumChildren(rules []*Rule) (count int) {
	if len(r.Contains) == 0 {
		return
	}

	for color, amount := range r.Contains {
		bag := FindRule(color, rules)
		if bag == nil {
			fmt.Println("oh maaan")
			return
		}
		count += (bag.NumChildren(rules) * amount) + amount
	}

	return
}
