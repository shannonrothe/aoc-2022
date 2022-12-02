package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day_2.txt
var input string

func part1(lines []string) int {
	ma := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	w := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	d := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	t := 0
	for _, line := range lines {
		p := strings.Split(line, " ")
		o := p[0]
		m := p[1]

		t += ma[m]
		if w[o] == m {
			t += 6
		} else if d[o] == m {
			t += 3
		}
	}

	return t
}

func part2(lines []string) int {
	l := map[string]int{
		"A": 3,
		"B": 1,
		"C": 2,
	}
	d := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	w := map[string]int{
		"A": 2,
		"B": 3,
		"C": 1,
	}

	t := 0
	for _, line := range lines {
		p := strings.Split(line, " ")
		o := p[0]
		m := p[1]

		switch m {
		case "X": // lose
			t += 0 + l[o]
		case "Y": // draw
			t += 3 + d[o]
		case "Z": // win
			t += 6 + w[o]
		}
	}

	return t
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}
