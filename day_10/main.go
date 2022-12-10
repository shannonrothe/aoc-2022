package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day_10.txt
var input string

func parseLine(line string) (string, int) {
	s := strings.Split(line, " ")
	ins := s[0]
	if ins == "noop" {
		return ins, 0
	}
	n, _ := strconv.Atoi(s[1])
	return ins, n
}

func part1(lines []string) int {
	s := 0
	x := 1
	c := 0
	m := map[int]int{}

	for _, l := range lines {
		ins, n := parseLine(l)
		if ins != "noop" {
			m[c+2] = n
			c++
		}
		c += 1
	}

	for i := 0; i < c; i++ {
		if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
			s += i * x
		}

		if v, ok := m[i]; ok {
			x += v
			delete(m, i)
		}
	}

	return s
}

func part2(lines []string) string {
	s := 0
	x := 1
	c := 0
	m := map[int]int{}
	crt := make([]string, 6)

	for _, l := range lines {
		ins, n := parseLine(l)
		if ins != "noop" {
			m[c+2] = n
			c++
		}
		c += 1
	}

	l := x
	h := x + 2
	p := 0
	for i := 0; i < c; i++ {
		if p == 40 {
			p = 0
		}

		r := i / 40
		if p >= l && p <= h {
			crt[r] += "#"
		} else {
			crt[r] += "."
		}

		if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
			s += i * x
		}

		if v, ok := m[i]; ok {
			x += v
			delete(m, i)
		}

		p += 1
		l = x
		h = x + 2
	}

	return strings.Join(crt, "\n")
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func main() {
	i := parse(input)
	fmt.Printf("part 1: %d\n", part1(i))
	fmt.Printf("part 2:\n%s\n", part2(i))
}
