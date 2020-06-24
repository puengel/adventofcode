package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type imgLayer struct {
	layer [][]int
	n0    int
	n1    int
	n2    int
}

func main() {
	fmt.Println("2019 #8")

	readInput()

	layers := createLayers()

	var best imgLayer
	lowest := height * width

	for _, l := range layers {
		l.checklayer()

		if l.n0 < lowest {
			best = l
			lowest = l.n0
		}
	}

	printLayer(best.layer)
	fmt.Printf("Checksum: %d\n", best.checksum())

	pic := initPic(2)

	for _, l := range layers {
		pic.applyLayer(l)
	}

	pic.printPic()
}

const (
	width  = 25
	height = 6
)

var (
	numbers []int
)

func readInput() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// file = string(f)

	s := string(f)
	numbers = make([]int, 0)

	for _, val := range s {
		sval := string(val)
		nval, err := strconv.Atoi(sval)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, nval)

	}

}

func createLayers() []imgLayer {

	ptr := 0

	layers := make([]imgLayer, 0)

	for ptr < len(numbers) {
		layer := make([][]int, height)

		for y := 0; y < height; y++ {
			layer[y] = make([]int, width)
			for x := 0; x < width; x++ {
				layer[y][x] = numbers[ptr]
				ptr++
			}
		}
		// printLayer(layer)
		// fmt.Println("")

		layers = append(layers, imgLayer{
			layer: layer,
		})
	}

	return layers
}

func printLayer(layer [][]int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(layer[y][x])
		}
		fmt.Print("\n")
	}
}

func (l *imgLayer) checklayer() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch l.layer[y][x] {
			case 0:
				l.n0++
			case 1:
				l.n1++
			case 2:
				l.n2++
			default:
				fmt.Println("huh?")
			}
		}
	}
}

func (l *imgLayer) checksum() int {
	return l.n1 * l.n2
}

func initPic(defaultValue int) imgLayer {
	pic := imgLayer{
		layer: make([][]int, height),
	}

	for y := 0; y < height; y++ {
		pic.layer[y] = make([]int, width)
		for x := 0; x < width; x++ {
			pic.layer[y][x] = defaultValue
		}
	}

	return pic
}

func (l *imgLayer) applyLayer(apply imgLayer) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if l.layer[y][x] == 2 {
				l.layer[y][x] = apply.layer[y][x]
			}
		}
	}
}

func (l *imgLayer) printPic() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch l.layer[y][x] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("X")
			case 2:
				fmt.Print("O")
			}
		}
		fmt.Print("\n")
	}
}
