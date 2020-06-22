package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type dude struct {
	name   string
	around string
	orbits int
}

var (
	dudes map[string]dude
)

func main() {
	fmt.Println("2019 #6")

	err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(dudes)

	sum := calcOrbits()

	fmt.Printf("All orbits: %d\n", sum)

	a1 := ancestors(dudes["YOU"])
	a2 := ancestors(dudes["SAN"])

	fmt.Printf("%+v\n%+v\v", a1, a2)

	p := path(a1, a2)

	fmt.Printf("Path YOU-SAN: %d\n", p)

}

func readInput() error {

	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}

	splits := strings.Split(string(file), "\n")

	dudes = make(map[string]dude, len(splits))
	// make com dude
	dudes["COM"] = dude{
		name:   "COM",
		around: "",
		orbits: 0,
	}

	for _, val := range splits {
		info := strings.Split(val, ")")
		if len(info) != 2 {
			return errors.New("This should be 2")
		}

		dudes[info[1]] = dude{
			name:   info[1],
			around: info[0],
			orbits: 0,
		}
	}

	return nil
}

func calcOrbits() int {

	sum := 0

	for _, val := range dudes {
		orbits := calcOrbit(val)
		sum += orbits
	}

	return sum

}

func calcOrbit(d dude) int {
	if d.around == "" || d.orbits != 0 {
		return d.orbits
	}

	return calcOrbit(dudes[d.around]) + 1
}

func ancestors(d dude) []string {

	l := make([]string, 0)

	a := dudes[d.around]
	for a.name != "COM" {
		l = append(l, a.around)
		a = dudes[a.around]
	}

	return l

}

func path(a1, a2 []string) int {

	for {
		if a1[len(a1)-2] != a2[len(a2)-2] {
			break
		}
		a1, a2 = a1[:len(a1)-1], a2[:len(a2)-1]
	}

	return len(a1) + len(a2)
}
