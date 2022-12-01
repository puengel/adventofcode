package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	elfs, err := ReadElfs("input", "\n")
	if err != nil {
		return
	}

	most := 0
	for _, elf := range elfs {
		if elf > most {
			most = elf
		}
	}

	fmt.Printf("Part 1: %d\n", most)

	sort.Ints(elfs)

	lastthree := elfs[len(elfs)-3:]
	sum := 0
	for _, elf := range lastthree {
		sum += elf
	}
	fmt.Printf("Part 2: %d\n", sum)
}

// ReadElfs - return int array from given input file and separator
func ReadElfs(path, separator string) (elfs []int, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	strVals := strings.Split(string(file), separator)

	elf := 0
	for _, val := range strVals {
		if len(val) == 0 {
			elfs = append(elfs, elf)
			elf = 0
			continue
		}

		iVal, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		elf += iVal
	}

	return
}
