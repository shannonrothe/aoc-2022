package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day_3.txt
var input string

const al = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func split(rs string) (string, string) {
	half := len(rs) / 2
	a := rs[:half]
	b := rs[half:]
	return a, b
}

func part1(rs []string) int {
	s := 0

	for _, r := range rs {
		a, b := split(r)
		c := new(rune)

		for _, i := range a {
			for _, j := range b {
				if i == j {
					*c = i
					break
				}
			}
		}

		p := strings.Index(al, string(*c)) + 1
		s += p
	}
	return s
}

func common(d map[string]int) string {
	c := new(string)
	for l, n := range d {
		if n == 6 { // 1 + 2 + 3
			*c = l
			break
		}
	}
	return *c
}

func indexOf(a []string, s string) int {
	for i, r := range a {
		if r == s {
			return i
		}
	}
	return -1
}

func part2(rs []string) int {
	s := 0
	for i := 0; i < len(rs); i += 3 {
		m := map[string]int{}
		sn := []string{}
		for j := i; j < i+3; j++ {
			l := rs[j]
			for _, r := range l {
				if indexOf(sn, string(r)) == -1 {
					m[string(r)] += (j - i) + 1
					sn = append(sn, string(r))
				}
			}
			sn = []string{}
		}
		l := common(m)
		s += strings.Index(al, l) + 1
	}
	return s
}

func main() {
	rs := strings.Split(input, "\n")

	fmt.Printf("part 1: %d\n", part1(rs))
	fmt.Printf("part 2: %d\n", part2(rs))
}
