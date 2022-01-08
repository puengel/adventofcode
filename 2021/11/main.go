package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"strconv"
)

type Octopus struct {
	ID      string
	Energy  int
	Flashed bool

	Neighbours map[string]*Octopus
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	all := make(map[string]*Octopus)

	fmt.Println(inputs[0])
	for y := 0; y < len(inputs); y++ {
		for x := 0; x < len(inputs[y]); x++ {
			e, err := strconv.Atoi(string(inputs[y][x]))
			if err != nil {
				fmt.Println("meh")
				return
			}
			o := Octopus{
				ID:         fmt.Sprintf("%d:%d", y, x),
				Energy:     e,
				Neighbours: make(map[string]*Octopus),
			}
			if neigh, ok := all[fmt.Sprintf("%d:%d", y-1, x-1)]; ok {
				o.Neighbours[neigh.ID] = neigh
			}
			if neigh, ok := all[fmt.Sprintf("%d:%d", y-1, x)]; ok {
				o.Neighbours[neigh.ID] = neigh
			}
			if neigh, ok := all[fmt.Sprintf("%d:%d", y-1, x+1)]; ok {
				o.Neighbours[neigh.ID] = neigh
			}
			if neigh, ok := all[fmt.Sprintf("%d:%d", y, x-1)]; ok {
				o.Neighbours[neigh.ID] = neigh
			}
			all[o.ID] = &o
		}
	}

	for _, o := range all {
		for _, neigh := range o.Neighbours {
			neigh.Neighbours[o.ID] = o
		}
	}

	flashes := 0
	allFlashStep := 0
	allFlashed := false
	for step := 0; !allFlashed; step++ {
		stepFlashes := 0

		// increase all
		for _, o := range all {
			o.Energy++
		}

		// flash
		for _, o := range all {
			o.flash()
		}

		// reset
		for _, o := range all {
			if o.reset() {
				stepFlashes++
			}
		}

		if step < 100 {
			flashes += stepFlashes
		}

		if stepFlashes == len(all) {
			allFlashStep = step + 1
			allFlashed = true
		}

	}

	fmt.Printf("Part 1: Flashes %d\n", flashes)
	fmt.Printf("Part 2: all flashed at step %d\n", allFlashStep)
}

func (o *Octopus) flash() {
	if o.Energy <= 9 {
		return
	}

	if o.Flashed {
		return
	}

	o.Flashed = true
	for _, neigh := range o.Neighbours {
		neigh.Energy++
		neigh.flash()
	}
}

func (o *Octopus) reset() (hasFlashed bool) {
	hasFlashed = o.Flashed
	if hasFlashed {
		o.Energy = 0
	}
	o.Flashed = false
	return
}
