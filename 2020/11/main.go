package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {
	inputs, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	var origGrid [][]string
	for _, row := range inputs {
		var nRow []string

		for _, col := range row {
			nRow = append(nRow, string(col))
		}
		origGrid = append(origGrid, nRow)
	}

	grid := duplicateGrid(origGrid)

	var rounds int
	changed := true

	for changed {

		rounds++
		changed = false

		newGrid := duplicateGrid(grid)

		for y, row := range grid {
			for x := range row {
				if changeSeat(x, y, grid, &newGrid) {
					changed = true

				}
			}
		}

		grid = newGrid

		// printGrid(grid)
	}

	fmt.Printf("Part 1: %d\n", occupied(grid))

	grid = duplicateGrid(origGrid)
	rounds = 0
	changed = true

	for changed {

		rounds++
		changed = false

		newGrid := duplicateGrid(grid)

		for y, row := range grid {
			for x := range row {
				if changeSeat2(x, y, grid, &newGrid) {
					changed = true
				}
			}
		}

		grid = newGrid

		// printGrid(grid)
	}

	fmt.Printf("Part 2: %d\n", occupied(grid))
}

func occupied(grid [][]string) (seats int) {
	for _, row := range grid {
		for _, val := range row {
			if val == "#" {
				seats++
			}
		}
	}

	return
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Print("\n")
	}
}

func duplicateGrid(grid [][]string) [][]string {
	duplicate := make([][]string, len(grid))
	for i := range grid {
		duplicate[i] = make([]string, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	return duplicate
}

func changeSeat(x, y int, grid [][]string, newGrid *[][]string) (changed bool) {

	if grid[y][x] == "." {
		return false
	}

	var neighbours int

	// check neighbours
	if (x > 0) && (y > 0) && (grid[y-1][x-1] == "#") {
		neighbours++
	}
	if (y > 0) && (grid[y-1][x] == "#") {
		neighbours++
	}
	if (x < len(grid[0])-1) && (y > 0) && (grid[y-1][x+1] == "#") {
		neighbours++
	}
	if (x > 0) && (grid[y][x-1] == "#") {
		neighbours++
	}
	if (x < len(grid[0])-1) && (grid[y][x+1] == "#") {
		neighbours++
	}
	if (x > 0) && (y < len(grid)-1) && (grid[y+1][x-1] == "#") {
		neighbours++
	}
	if (y < len(grid)-1) && (grid[y+1][x] == "#") {
		neighbours++
	}
	if (x < len(grid[0])-1) && (y < len(grid)-1) && (grid[y+1][x+1] == "#") {
		neighbours++
	}

	// Change?
	if grid[y][x] == "L" && neighbours == 0 {
		(*newGrid)[y][x] = "#"
		changed = true
	}
	if grid[y][x] == "#" && neighbours >= 4 {
		(*newGrid)[y][x] = "L"
		changed = true
	}

	return
}

func checkDir(x, y, dx, dy int, grid [][]string) (occupied bool) {
	// border
	if (x == 0 && dx < 0) ||
		(y == 0 && dy < 0) ||
		(x == len(grid[0])-1 && dx > 0) ||
		(y == len(grid)-1 && dy > 0) {
		return false
	}

	nextSeat := grid[y+dy][x+dx]
	if nextSeat == "." {
		return checkDir(x+dx, y+dy, dx, dy, grid)
	}
	if nextSeat == "#" {
		occupied = true
	}
	return
}

func changeSeat2(x, y int, grid [][]string, newGrid *[][]string) (changed bool) {

	if grid[y][x] == "." {
		return false
	}

	var neighbours int

	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			if checkDir(x, y, dx, dy, grid) {
				neighbours++
			}
		}

	}

	// Change?
	if grid[y][x] == "L" && neighbours == 0 {
		(*newGrid)[y][x] = "#"
		changed = true
	}
	if grid[y][x] == "#" && neighbours >= 5 {
		(*newGrid)[y][x] = "L"
		changed = true
	}

	return
}
