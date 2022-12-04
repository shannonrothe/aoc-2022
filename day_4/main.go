package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day_4.txt
var input string

type pair struct {
	l int
	h int
}

func newPair(l int, h int) pair {
	return pair{
		l,
		h,
	}
}

func (p *pair) contains(o *pair) bool {
	return p.l <= o.l && p.h >= o.h
}

func (p *pair) overlaps(o *pair) bool {
	for i := p.l; i <= p.h; i++ {
		for j := o.l; j <= o.h; j++ {
			if i == j {
				return true
			}
		}
	}
	return false
}

func bounds(pair string) (int, int) {
	s := strings.Split(pair, "-")
	n1, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		panic(err)
	}
	p1 := int(n1)
	n2, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		panic(err)
	}
	p2 := int(n2)
	return p1, p2
}

func pairs(line string) (string, string) {
	s := strings.Split(line, ",")
	return s[0], s[1]
}

func part1(lines []string) int {
	p := 0

	for _, line := range lines {
		p1, p2 := pairs(line)
		fl, fh := bounds(p1)
		p1p := newPair(fl, fh)
		sl, sh := bounds(p2)
		p2p := newPair(sl, sh)

		if p1p.contains(&p2p) || p2p.contains(&p1p) {
			p += 1
		}
	}

	return p
}

func part2(lines []string) int {
	o := 0

	for _, line := range lines {
		p1, p2 := pairs(line)
		fl, fh := bounds(p1)
		p1p := newPair(fl, fh)
		sl, sh := bounds(p2)
		p2p := newPair(sl, sh)

		if p1p.overlaps(&p2p) || p2p.overlaps(&p1p) {
			o += 1
		}
	}

	return o
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}
