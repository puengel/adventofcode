package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
)

func main() {
	in, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(in)

	// grows at most 6 into each dir => 12 for each dimension
	// place first layed in center of [15][22][22] cube
	// making extra space for edges which will never have to change

	rounds := 6
	lenY := len(in)
	lenX := len(in[0])
	maxY := lenY + 2 + 2*rounds
	maxX := lenX + 2 + 2*rounds
	maxZ := 3 + 2*rounds
	maxW := maxZ

	cube := make([][][]bool, maxZ)
	for z := range cube {
		cube[z] = make([][]bool, maxY)
		for y := range cube[z] {
			cube[z][y] = make([]bool, maxX)
			for x := range cube[z][y] {
				if z == len(cube)/2 && (x >= rounds+1 && x < rounds+1+lenX) && (y >= rounds+1 && y < rounds+1+lenY) {
					// fmt.Printf("got: %v\n", in[y-(rounds+1)][x-(rounds+1)])
					if in[y-(rounds+1)][x-(rounds+1)] == '#' {
						cube[z][y][x] = true
					}
				}

			}
		}
	}
	// fmt.Println(cube[7])
	// fmt.Println(countActive(cube))

	for round := 0; round < rounds; round++ {
		fmt.Println("Round: ", round+1)

		res := copyDim3(cube)

		for z := range cube {
			for y := range cube[z] {
				for x := range cube[z][y] {
					// can do nothing if border
					if z <= 0 || z >= len(cube)-1 || y <= 0 || y >= lenY+1+2*rounds || x <= 0 || x >= lenX+1+2*rounds {
						continue
					}

					activeN := 0
					for dz := -1; dz < 2; dz++ {
						for dy := -1; dy < 2; dy++ {
							for dx := -1; dx < 2; dx++ {
								if dx == 0 && dy == 0 && dz == 0 {
									// self
									continue
								}
								if cube[z+dz][y+dy][x+dx] {
									activeN++
								}
							}
						}
					}

					if (cube[z][y][x] && (activeN == 2 || activeN == 3)) ||
						(!cube[z][y][x] && activeN == 3) {
						res[z][y][x] = true
					} else {
						res[z][y][x] = false
					}
				}
			}
		}

		// Set all at once
		cube = res

		fmt.Println(countActive(cube))
	}

	fmt.Printf("Part 1: %d\n", countActive(cube))

	hcube := make([][][][]bool, maxW)
	for w := range hcube {
		hcube[w] = make([][][]bool, maxZ)
		for z := range hcube[w] {
			hcube[w][z] = make([][]bool, maxY)
			for y := range hcube[w][z] {
				hcube[w][z][y] = make([]bool, maxX)
				for x := range hcube[w][z][y] {
					if w == len(hcube)/2 && z == len(hcube[0])/2 && (x >= rounds+1 && x < rounds+1+lenX) && (y >= rounds+1 && y < rounds+1+lenY) {
						// fmt.Printf("got: %v\n", in[y-(rounds+1)][x-(rounds+1)])
						if in[y-(rounds+1)][x-(rounds+1)] == '#' {
							hcube[w][z][y][x] = true
						}
					}
				}

			}
		}
	}

	for round := 0; round < rounds; round++ {
		fmt.Println("Round: ", round+1)

		res := copyDim4(hcube)

		for w := range hcube {
			for z := range hcube[w] {
				for y := range hcube[w][z] {
					for x := range hcube[w][z][y] {
						// can do nothing if border
						if w <= 0 || w >= len(hcube)-1 || z <= 0 || z >= len(hcube[0])-1 || y <= 0 || y >= lenY+1+2*rounds || x <= 0 || x >= lenX+1+2*rounds {
							continue
						}

						activeN := 0
						for dw := -1; dw < 2; dw++ {
							for dz := -1; dz < 2; dz++ {
								for dy := -1; dy < 2; dy++ {
									for dx := -1; dx < 2; dx++ {
										if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
											// self
											continue
										}
										if hcube[w+dw][z+dz][y+dy][x+dx] {
											activeN++
										}
									}
								}
							}
						}

						if (hcube[w][z][y][x] && (activeN == 2 || activeN == 3)) ||
							(!hcube[w][z][y][x] && activeN == 3) {
							res[w][z][y][x] = true
						} else {
							res[w][z][y][x] = false
						}
					}
				}
			}
		}

		// Set all at once
		hcube = res

		fmt.Println(countActive4(hcube))
	}

	fmt.Printf("Part 2: %d\n", countActive4(hcube))
}

func copyDim3(dim [][][]bool) (cp [][][]bool) {
	cp = make([][][]bool, len(dim))
	for z := range dim {
		cp[z] = make([][]bool, len(dim[z]))
		for y := range dim[z] {
			cp[z][y] = make([]bool, len(dim[z][y]))
			copy(cp[z][y], dim[z][y])
		}
	}
	return
}

func copyDim4(dim [][][][]bool) (cp [][][][]bool) {
	cp = make([][][][]bool, len(dim))
	for w := range dim {
		cp[w] = make([][][]bool, len(dim[w]))
		for z := range dim[w] {
			cp[w][z] = make([][]bool, len(dim[w][z]))
			for y := range dim[w][z] {
				cp[w][z][y] = make([]bool, len(dim[w][z][y]))
				copy(cp[w][z][y], dim[w][z][y])
			}
		}
	}
	return
}

func countActive(dim [][][]bool) int {
	var active int
	for z := range dim {
		for y := range dim[z] {
			for x := range dim[z][y] {
				if dim[z][y][x] {
					active++
				}

			}
		}
	}
	return active
}

func countActive4(dim [][][][]bool) int {
	var active int
	for w := range dim {
		for z := range dim[w] {
			for y := range dim[w][z] {
				for x := range dim[w][z][y] {
					if dim[w][z][y][x] {
						active++
					}

				}
			}
		}
	}
	return active
}
