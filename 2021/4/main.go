package main

import (
	"adventofcode/2021/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	Number int
	Marked bool
}

type Board struct {
	Win   bool
	Field [][]Cell
}

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n\n")
	if err != nil {
		return
	}

	values := []int{}
	strVals := strings.Split(inputs[0], ",")

	for _, val := range strVals {
		iVal, err := strconv.Atoi(val)
		if err != nil {
			return
		}

		values = append(values, iVal)
	}

	// init boards
	boards := []*Board{}
	for _, boardStr := range inputs[1:] {
		boards = append(boards, NewBoard(boardStr))
	}

	// run values
	var currentVal int
	var currentBoard int
	firstBoard := -1
	firstVal := 0
	lastBoard := -1
	lastVal := 0
	for _, val := range values {
		currentVal = val
		// found := false

		for idx, board := range boards {
			currentBoard = idx
			if board.Win {
				continue
			}
			if board.check(val) {
				if firstBoard == -1 {
					firstBoard = currentBoard
					firstVal = currentVal
				} else {
					lastBoard = currentBoard
					lastVal = currentVal
				}
				// found = true
			}
		}
		// if found {
		// 	break
		// }
	}

	sumUnmarked := boards[firstBoard].sumUnmarked()

	fmt.Printf("Part 1: firstBoard: %v firstVal: %d, sumUnmarked: %d\n", boards[firstBoard], firstVal, sumUnmarked)
	fmt.Printf("Part 1: %d\n", firstVal*sumUnmarked)
	fmt.Printf("Part 2: lastBoard: %v lastVal: %d\n", boards[lastBoard], lastVal)
	fmt.Printf("Part 2: %d\n", lastVal*boards[lastBoard].sumUnmarked())
}

func NewBoard(board string) *Board {

	b := Board{
		Field: make([][]Cell, 5),
	}

	rows := strings.Split(board, "\n")
	for i, row := range rows {
		b.Field[i] = make([]Cell, 5)
		row = strings.ReplaceAll(strings.ReplaceAll(strings.Trim(row, " "), "  ", ","), " ", ",")
		// fmt.Println(row)
		cells := strings.Split(row, ",")

		for j, cell := range cells {
			val, err := strconv.Atoi(cell)
			if err != nil {
				fmt.Println("fuck", err, row)
				return &b
			}
			b.Field[i][j] = Cell{
				Number: val,
				Marked: false,
			}
		}
	}

	// fmt.Println(b)

	return &b
}

func (b *Board) check(number int) (bingo bool) {
	if b.Win {
		return true
	}

	// mark
	for i := range b.Field {
		for j := range b.Field[i] {
			if b.Field[i][j].Number == number {
				b.Field[i][j].Marked = true
			}
		}
	}

	// check bingo
	bingo = b.hasBingo()
	if bingo {
		b.Win = true
	}
	return bingo
}

func (b *Board) hasBingo() bool {

	for i := range b.Field {
		// vertical
		hasVertical := true
		for j := range b.Field[i] {
			if !b.Field[i][j].Marked {
				hasVertical = false
				break
			}
		}

		if hasVertical {
			b.Win = true
			return true
		}

		// horizontal
		hasHorizontal := true
		for j := range b.Field[i] {
			if !b.Field[j][i].Marked {
				hasHorizontal = false
				break
			}
		}

		if hasHorizontal {
			b.Win = true
			return true
		}
	}

	// diagonal
	hasDiagonal := true
	for i := range b.Field {
		if !b.Field[i][i].Marked {
			hasDiagonal = false
			break
		}
	}
	if hasDiagonal {
		b.Win = true
		return true
	}
	hasDiagonal = true
	for i := range b.Field {
		if !b.Field[i][len(b.Field)-1-i].Marked {
			hasDiagonal = false
			break
		}
	}
	if hasDiagonal {
		b.Win = true
		return true
	}

	return false
}

func (b *Board) sumUnmarked() (sum int) {
	for i := range b.Field {
		for j := range b.Field[i] {
			if !b.Field[i][j].Marked {
				sum += b.Field[i][j].Number
			}
		}
	}
	return
}
