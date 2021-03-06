package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	dir  Direction
	dist int
}

// Direction to move towards
type Direction string

const (
	left  Direction = "L"
	up    Direction = "U"
	down  Direction = "D"
	right Direction = "R"
)

func main() {

	fmt.Println("2019 #3")

	red, blue, err := readInput()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("red: %+v\n", red)
	fmt.Printf("blue: %+v\n", blue)

	rInst, err := parseWire(red)
	if err != nil {
		fmt.Print(err)
		return
	}
	bInst, err := parseWire(blue)
	if err != nil {
		fmt.Print(err)
		return
	}

	wMap, err := createMap(rInst, bInst)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("map: (%d,%d)\n", len(wMap), len(wMap[0]))

	walkMap(rInst, "r", wMap)
	walkMap(bInst, "b", wMap)

	findX(wMap)
}

func readInput() (red, blue []string, err error) {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, nil, err
	}

	both := strings.Split(string(file), "\n")
	red = strings.Split(both[0], ",")
	blue = strings.Split(both[1], ",")

	return
}

func parseWire(wire []string) (inst []instruction, err error) {

	inst = make([]instruction, len(wire))

	for idx, val := range wire {
		dir := Direction(val[0])
		dist, err := strconv.Atoi(val[1:])
		if err != nil {
			return nil, err
		}

		inst[idx] = instruction{
			dir:  dir,
			dist: dist,
		}
	}

	return
}

func createMap(redInst, blueInst []instruction) (wMap [][]string, err error) {

	rR := 0
	rU := 0
	for _, inst := range redInst {
		switch inst.dir {
		case right:
			rR += inst.dist
		case left:
			rR += inst.dist
		case up:
			rU += inst.dist
		case down:
			rU += inst.dist
		}
	}

	bR := 0
	bU := 0
	for _, inst := range blueInst {
		switch inst.dir {
		case right:
			bR += inst.dist
		case left:
			bR += inst.dist
		case up:
			bU += inst.dist
		case down:
			bU += inst.dist
		}
	}

	if bR > rR {
		rR = bR
	}

	if bU < rU {
		rU = bU
	}

	rR *= 2
	rU *= 2

	wMap = make([][]string, rR)
	for i := range wMap {
		wMap[i] = make([]string, rU)
	}

	return
}

type position struct {
	x int
	y int
}

func walkMap(wire []instruction, color string, wMap [][]string) {

	pos := position{
		x: len(wMap) / 2,
		y: len(wMap[0]) / 2,
	}

	step := 0

	for _, inst := range wire {
		switch inst.dir {
		case up:
			for i := 0; i < inst.dist; i++ {
				pos.y++
				step++
				mark(pos, color, step, wMap)
			}
		case down:
			for i := 0; i < inst.dist; i++ {
				pos.y--
				step++
				mark(pos, color, step, wMap)
			}
		case left:
			for i := 0; i < inst.dist; i++ {
				pos.x--
				step++
				mark(pos, color, step, wMap)
			}
		case right:
			for i := 0; i < inst.dist; i++ {
				pos.x++
				step++
				mark(pos, color, step, wMap)
			}
		}
	}
}

func mark(pos position, color string, step int, wMap [][]string) {

	// fmt.Println(pos)
	if wMap[pos.x][pos.y] == "" {
		wMap[pos.x][pos.y] = fmt.Sprintf("%s%d", color, step)
	} else if wMap[pos.x][pos.y][:1] == "x" {
		fmt.Println(wMap[pos.x][pos.y])
	} else if wMap[pos.x][pos.y][:1] != color {
		d, err := strconv.Atoi(wMap[pos.x][pos.y][1:])
		fmt.Printf("red: %d, blue: %d, sum: %d\n", d, step, d+step)
		if err != nil {
			fmt.Println(err)
			return
		}
		d += step
		wMap[pos.x][pos.y] = fmt.Sprintf("x%d", d)
		fmt.Println(wMap[pos.x][pos.y])
	}
}

func findX(wMap [][]string) {
	// center := position{
	// 	x: len(wMap) / 2,
	// 	y: len(wMap[0]) / 2,
	// }
	sum := 0
	sDist := len(wMap) * 2
	for _, x := range wMap {
		for _, y := range x {
			if y != "" && y[:1] == "x" {
				sum++
				dist, err := strconv.Atoi(y[1:])
				fmt.Println(y)
				if err != nil {
					fmt.Println(err)
					return
				}
				if dist < sDist {
					sDist = dist
				}
			}
		}
	}
	fmt.Printf("Sum crossings: %d\n", sum)
	fmt.Printf("Shortest: %d\n", sDist)
}
