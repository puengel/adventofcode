package main

import (
	"adventofcode/2020/pkg/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type comp struct {
	mask string
	mem  []int
}

type comp2 struct {
	mask string
	mem  map[int]int
}

func main() {
	instructions, err := input.ReadStrings("input.txt", "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	c := comp{
		mask: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		mem:  make([]int, 100),
	}

	for _, instruction := range instructions {
		c.handleInst(instruction)
	}

	fmt.Printf("Part 1: %d\n", c.sumMem())

	c2 := comp2{
		mask: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		mem:  make(map[int]int),
	}

	for _, instruction := range instructions {
		c2.handleInstV2(instruction)
	}

	fmt.Printf("Part 2: %d\n", c2.sumMem())
}

func (c *comp) sumMem() int {
	var sum int
	for _, val := range c.mem {
		sum += val
	}

	return sum
}

func (c *comp2) sumMem() int {
	var sum int
	for _, val := range c.mem {
		sum += val
	}

	return sum
}

func (c *comp2) handleInstV2(inst string) {
	splits := strings.Split(inst, " = ")
	if splits[0] == "mask" {
		c.mask = splits[1]
		return
	}

	index, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(splits[0], "mem["), "]"))
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := strconv.Atoi(splits[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	c.writeMemV2(val, index)
}

func (c *comp) handleInst(inst string) {
	splits := strings.Split(inst, " = ")
	if splits[0] == "mask" {
		c.mask = splits[1]
		return
	}

	index, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(splits[0], "mem["), "]"))
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err := strconv.Atoi(splits[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	c.writeMem(index, val)
}

func (c *comp2) writeMemV2(val, ind int) {
	indices := c.applyMaskV2(ind)

	for _, index := range indices {
		c.mem[index] = val
	}

}

func (c *comp) writeMem(index, val int) {
	if index >= len(c.mem) {
		tmp := make([]int, index+1)
		copy(tmp, c.mem)
		c.mem = tmp
	}

	c.mem[index] = c.applyMask(val)
}

func (c *comp2) applyMaskV2(val int) []int {
	r := strings.Count(c.mask, "X")
	possibilities := getPossibilities(r)

	arr := make([]int, len(possibilities))

	for arrIdx, poss := range possibilities {

		sVal := []rune(fmt.Sprintf("%036b", val))

		cpos := 0
		for i, v := range c.mask {
			if v == 49 {
				sVal[i] = v
			} else if v == 88 {
				if poss[cpos] == 1 {
					sVal[i] = 49
				} else {
					sVal[i] = 48
				}
				cpos++
			}
		}

		result, err := strconv.ParseInt(string(sVal), 2, 64)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		arr[arrIdx] = int(result)
	}

	return arr
}

func getPossibilities(num int) [][]int {
	rounds := int(math.Pow(2, float64(num)))

	result := make([][]int, rounds)

	for i := 0; i < rounds; i++ {
		bf := make([]int, num)
		btoflag(bf, i)
		result[i] = bf
	}

	return result
}

func btoflag(bf []int, n int) {
	for i := range bf {
		if n&(1<<i) != 0 {
			bf[i] = 1
		}
	}
}

func (c *comp) applyMask(val int) int {
	sVal := []rune(fmt.Sprintf("%036b", val))

	for i, v := range c.mask {
		if v == 48 || v == 49 {
			sVal[i] = v
		}
	}

	result, err := strconv.ParseInt(string(sVal), 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(result)
}
