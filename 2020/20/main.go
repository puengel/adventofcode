package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	id  int
	val []string
	border
	neightbours []*tile
}

type border struct {
	top    string
	left   string
	right  string
	bottom string
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var tiles []*tile

	for _, in := range inputs {
		tile := NewTile(strings.Split(in, "\n"))
		tiles = append(tiles, tile)
		// fmt.Println(tile.id, tile.right)
	}

	p1 := 1

	var startCorner *tile

	for idx, t := range tiles {
		fmt.Println(t.id)

		for _, possN := range tiles[idx:] {

			for _, n := range t.neightbours {
				// already neighbours
				if n == possN {
					break
				}
			}

			fits := false
			for i := 0; i < 16; i++ {
				// Flip
				if i == 4 || i == 12 {
					possN.FlipHorizontal()
				}
				if i == 8 {
					possN.FlipVertical()
				}
				if possN.top == Reverse(t.bottom) {
					fits = true
				}
				if possN.right == Reverse(t.left) {
					fits = true
				}
				if possN.bottom == Reverse(t.top) {
					fits = true
				}
				if possN.left == Reverse(t.right) {
					fits = true
				}
				if fits {
					break
				}
				possN.rotateR()
			}
			if fits {
				t.neightbours = append(t.neightbours, possN)
				possN.neightbours = append(possN.neightbours, t)
			}
		}

		fmt.Print(t.id, len(t.neightbours))
		fmt.Print(" [ ")
		for _, n := range t.neightbours {
			fmt.Printf(" %d ", n.id)
		}
		fmt.Println(" ]")

		if len(t.neightbours) == 2 {
			p1 *= t.id
			startCorner = t
		}
	}

	fmt.Printf("Part 1: %d\n", p1)

	startCorner.FullImage()
}

func (t *tile) FullImage() *tile {

	// var val []string

	// make it the top right tile
	for r := 0; r > 4; r++ {
		var right bool
		var bottom bool
		for _, n := range t.neightbours {
			if t.right == Reverse(n.left) {
				right = true
			}
			if t.bottom == Reverse(n.top) {
				bottom = true
			}
		}

		if right && bottom {
			fmt.Println("yay")
			break
		}
		t.rotateR()
	}

	return nil
}

func (t *tile) rotateR() {
	tmp := t.right
	t.right = t.top
	t.top = t.left
	t.left = t.bottom
	t.bottom = tmp
}

func (t *tile) FixVal() {
	for i := 0; i < 16; i++ {
		if t.val[0] == t.top {
			return
		}

	}
}

func RotateValR(val []string) []string {
	nval := make([][]rune, len(val))
	for i := 0; i < len(val); i++ {
		nval[i] = make([]rune, len(val[i]))
		for j := 0; j < len(val[0]); j++ {
			nval[i][j] = rune(val[len(val[0])-j-1][i])
		}
	}

	for i, row := range nval {
		val[i] = string(row)
	}

	return val
}

func NewTile(input []string) *tile {
	id, err := strconv.Atoi(strings.Split(strings.ReplaceAll(input[0], ":", ""), " ")[1])
	if err != nil {
		fmt.Println(err)
		return nil
	}

	top := input[1]
	bottom := Reverse(input[len(input)-1])

	var left string
	var right string

	for _, row := range input[1:] {
		left = string(row[0]) + left
		right += string(row[len(row)-1])
	}

	return &tile{
		id:  id,
		val: input[1:],
		border: border{
			top:    top,
			right:  right,
			bottom: bottom,
			left:   left,
		},
	}
}

func (t *tile) FlipHorizontal() {
	t.right = Reverse(t.right)
	t.left = Reverse(t.left)
}

func (t *tile) FlipVertical() {
	t.top = Reverse(t.top)
	t.bottom = Reverse(t.bottom)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
