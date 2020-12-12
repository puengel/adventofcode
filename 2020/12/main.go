package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strconv"
)

type ship struct {
	NS int
	EW int
	Facing
	WP Waypoint
}

type Waypoint struct {
	NS int
	EW int
}

type Facing string

const (
	N Facing = "N"
	E Facing = "E"
	S Facing = "S"
	W Facing = "W"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := NewShip()

	for _, action := range inputs {
		err = s.DoAction(action)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Part 1: %d\n", s.ManhattanDist())

	s = NewShip()

	for _, action := range inputs {
		err = s.DoWPAction(action)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Part 2: %d\n", s.ManhattanDist())
}

func (s *ship) ManhattanDist() int {
	x := s.EW
	y := s.NS
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func NewShip() *ship {
	return &ship{
		NS:     0,
		EW:     0,
		Facing: E,
		WP: Waypoint{
			NS: 1,
			EW: 10,
		},
	}
}

func (s *ship) DoAction(action string) error {
	a := action[:1]
	v, err := strconv.Atoi(action[1:])
	if err != nil {
		return err
	}

	switch a {
	case "N":
		s.NS += v
	case "E":
		s.EW += v
	case "S":
		s.NS -= v
	case "W":
		s.EW -= v
	case "R":
		s.Turn(a, v/90)
	case "L":
		s.Turn(a, v/90)
	case "F":
		s.Forward(v)
	}

	return nil
}

func (s *ship) Forward(dist int) {
	switch s.Facing {
	case N:
		s.NS += dist
	case E:
		s.EW += dist
	case S:
		s.NS -= dist
	case W:
		s.EW -= dist
	}
}

func (s *ship) Turn(dir string, times int) {
	for times > 0 {
		if dir == "R" {
			s.TurnR()
		} else {
			s.TurnL()
		}
		times--
	}
}

func (s *ship) TurnR() {
	switch s.Facing {
	case N:
		s.Facing = E
	case E:
		s.Facing = S
	case S:
		s.Facing = W
	case W:
		s.Facing = N
	}
}

func (s *ship) TurnL() {
	switch s.Facing {
	case N:
		s.Facing = W
	case E:
		s.Facing = N
	case S:
		s.Facing = E
	case W:
		s.Facing = S
	}
}

func (s *ship) DoWPAction(action string) (err error) {
	a := action[:1]
	v, err := strconv.Atoi(action[1:])
	if err != nil {
		return err
	}

	switch a {
	case "N":
		s.WP.NS += v
	case "E":
		s.WP.EW += v
	case "S":
		s.WP.NS -= v
	case "W":
		s.WP.EW -= v
	case "R":
		s.WP.Rotate(a, v/90)
	case "L":
		s.WP.Rotate(a, v/90)
	case "F":
		s.FollowWP(v)
	}

	return
}

func (w *Waypoint) Rotate(dir string, times int) {
	for times > 0 {
		tmp := w.EW
		if dir == "R" {
			w.EW = w.NS
			w.NS = -tmp
		} else {
			w.EW = -w.NS
			w.NS = tmp
		}
		times--
	}
}

func (s *ship) FollowWP(times int) {
	s.EW += times * s.WP.EW
	s.NS += times * s.WP.NS
}
