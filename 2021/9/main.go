package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"sort"
	"strconv"
)

var heatmap [][]int

type Point struct {
	Height int
	X      int
	Y      int
}

func (p *Point) ID() string {
	return fmt.Sprintf("%d:%d", p.X, p.Y)
}

func (p *Point) FindValley() *Valley {
	v := &Valley{}
	v.Points = make(map[string]Point)
	v.Points[p.ID()] = *p
	p.FindNeighbours(v)
	return v
}

func toID(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func (p *Point) FindNeighbours(v *Valley) bool {
	newPoints := []Point{}
	// north
	if p.Y > 0 {
		if newP, ok := v.Points[toID(p.X, p.Y-1)]; !ok {

			newP = Point{
				X:      p.X,
				Y:      p.Y - 1,
				Height: heatmap[p.Y-1][p.X],
			}
			if newP.Height < 9 { //&& newP.Height > p.Height {
				v.Points[newP.ID()] = newP
				newPoints = append(newPoints, newP)
			}
		}
	}
	// south
	if p.Y < len(heatmap)-1 {
		if newP, ok := v.Points[toID(p.X, p.Y+1)]; !ok {

			newP = Point{
				X:      p.X,
				Y:      p.Y + 1,
				Height: heatmap[p.Y+1][p.X],
			}
			if newP.Height < 9 { //&& newP.Height > p.Height {
				v.Points[newP.ID()] = newP
				newPoints = append(newPoints, newP)
			}
		}
	}
	// east
	if p.X > 0 {
		if newP, ok := v.Points[toID(p.X-1, p.Y)]; !ok {

			newP = Point{
				X:      p.X - 1,
				Y:      p.Y,
				Height: heatmap[p.Y][p.X-1],
			}
			if newP.Height < 9 { //&& newP.Height > p.Height {
				v.Points[newP.ID()] = newP
				newPoints = append(newPoints, newP)
			}
		}
	}
	// west
	if p.X < len(heatmap[0])-1 {
		if newP, ok := v.Points[toID(p.X+1, p.Y)]; !ok {

			newP = Point{
				X:      p.X + 1,
				Y:      p.Y,
				Height: heatmap[p.Y][p.X+1],
			}
			if newP.Height < 9 { //&& newP.Height > p.Height {
				v.Points[newP.ID()] = newP
				newPoints = append(newPoints, newP)
			}
		}
	}

	for _, newPoint := range newPoints {
		if !newPoint.FindNeighbours(v) {
			fmt.Println("shit")
		}
	}

	return true
}

type Valley struct {
	Points map[string]Point
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		return
	}

	heatmap = make([][]int, len(inputs))
	for idx, line := range inputs {
		heatmap[idx] = make([]int, len(line))
		for x := 0; x < len(line); x++ {
			val, err := strconv.Atoi(string(line[x]))
			if err != nil {
				fmt.Print(err)
				return
			}

			heatmap[idx][x] = val
		}
	}

	minima := []Point{}
	riskLevel := 0
	for y := 0; y < len(heatmap); y++ {
		for x := 0; x < len(heatmap[y]); x++ {
			inMin := true
			curr := heatmap[y][x]

			// if y > 0 && x > 0 && heatmap[y-1][x-1] <= curr {
			// 	inMin = false
			// }
			if y > 0 && heatmap[y-1][x] <= curr {
				inMin = false
			}
			// if y > 0 && x < len(heatmap[y])-1 && heatmap[y-1][x+1] <= curr {
			// 	inMin = false
			// }
			if x > 0 && heatmap[y][x-1] <= curr {
				inMin = false
			}
			if x < len(heatmap[y])-1 && heatmap[y][x+1] <= curr {
				inMin = false
			}
			// if y < len(heatmap)-1 && x > 0 && heatmap[y+1][x-1] <= curr {
			// 	inMin = false
			// }
			if y < len(heatmap)-1 && heatmap[y+1][x] <= curr {
				inMin = false
			}
			// if y < len(heatmap)-1 && x < len(heatmap[y])-1 && heatmap[y+1][x+1] <= curr {
			// 	inMin = false
			// }

			if inMin {
				// fmt.Println("min", curr, y, x)
				minima = append(minima, Point{
					X:      x,
					Y:      y,
					Height: curr,
				})
				riskLevel += curr + 1
			}
		}
	}

	fmt.Printf("Part 1: risklevel %d\n", riskLevel)

	valleys := []*Valley{}
	for _, point := range minima {
		valleys = append(valleys, point.FindValley())
	}

	sort.Slice(valleys, func(i, j int) bool {
		return len(valleys[i].Points) > len(valleys[j].Points)
	})
	if len(valleys) < 3 {
		fmt.Println("not enough valleys")
		return
	}
	one := len(valleys[0].Points)
	two := len(valleys[1].Points)
	three := len(valleys[2].Points)

	fmt.Printf("Part 2: product of largest basins %d\n", one*two*three)
}
