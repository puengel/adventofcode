package main

import (
	"adventofcode/2020/pkg/input"
	"container/list"
	"fmt"
)

func main() {
	nums, err := input.ReadInts("input.txt", ",")
	if err != nil {
		fmt.Println(err)
		return
	}

	l := list.New()
	for _, v := range nums {
		l.PushBack(v)
	}

	// Have 3 already
	for i := len(nums) + 1; i <= 2020; i++ {
		last := l.Back()

		steps := 0
		seen := false

		for e := last.Prev(); e != nil; e = e.Prev() {
			steps++
			if e.Value == last.Value {
				seen = true
				break
			}
		}

		if seen {
			l.PushBack(steps)
		} else {
			l.PushBack(0)
		}

	}

	fmt.Printf("Part 1: %v\n", l.Back().Value)

	m := make(map[int]int)
	seen := make(map[int]bool)

	last := nums[0]
	for i, v := range nums {
		m[v] = i + 1
		seen[last] = true
		last = v
	}

	for i := len(nums) + 1; i <= 30000000; i++ {
		check := last
		if !seen[check] {
			seen[check] = true
			check = 0
		} else {
			a, ok := m[check]
			if !ok {
				fmt.Println("should not get here")
				return
			}

			check = i - a - 1
		}
		m[last] = i - 1
		last = check
	}

	fmt.Printf("Part 2: %d\n", last)
}
