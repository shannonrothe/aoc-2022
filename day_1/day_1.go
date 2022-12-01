package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day_1.txt
var input string

func part1(lines []string) int {
	c := []int{}
	cc := 0

	for _, l := range lines {
		if l == "" {
			c = append(c, cc)
			cc = 0
			continue
		}
		n, err := strconv.ParseInt(string(l), 10, 64)
		if err != nil {
			panic(err)
		}
		cc += int(n)
	}
	v, _ := max(c)
	return v
}

func part2(lines []string) int {
	c := []int{}
	cc := 0

	for _, l := range lines {
		if l == "" {
			c = append(c, cc)
			cc = 0
			continue
		}
		n, err := strconv.ParseInt(string(l), 10, 64)
		if err != nil {
			panic(err)
		}

		cc += int(n)
	}

	t := 0
	for l := 0; l < 3; l++ {
		v, i := max(c)
		t += v
		c = remove(c, i)
	}

	return t
}

func main() {
	lines := strings.Split(input, "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func max(c []int) (int, int) {
	m := c[0]
	mi := -1
	for i, v := range c {
		if m < v {
			m = v
			mi = i
		}
	}
	return m, mi
}
