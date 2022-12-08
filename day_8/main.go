package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day_8.txt
var input string

func height(b byte) int {
	n, _ := strconv.Atoi(string(b))
	return n
}

func up(lines []string, x, y int) []int {
	l := []int{}
	for yy := y - 1; yy >= 0; yy-- {
		l = append(l, height(lines[yy][x]))
	}
	return l
}

func right(lines []string, x, y int) []int {
	l := []int{}
	for xx := x + 1; xx < len(lines[0]); xx++ {
		l = append(l, height(lines[y][xx]))
	}
	return l
}

func down(lines []string, x, y int) []int {
	l := []int{}
	for yy := y + 1; yy < len(lines); yy++ {
		l = append(l, height(lines[yy][x]))
	}
	return l
}

func left(lines []string, x, y int) []int {
	l := []int{}
	for xx := x - 1; xx >= 0; xx-- {
		l = append(l, height(lines[y][xx]))
	}
	return l
}

func view_dist(n int, hs []int) int {
	vd := 0
	for _, h := range hs {
		if h >= n {
			vd += 1
			break
		}
		vd += 1
	}
	return vd
}

func visible(n int, hs []int) bool {
	vis := true
	for _, h := range hs {
		if h >= n {
			vis = false
			break
		}
	}
	return vis
}

func max(l []int) int {
	m := l[0]
	for _, v := range l {
		if v > m {
			m = v
		}
	}
	return m
}

func part1(lines []string) int {
	v := 2*(len(lines[0])-2) + 2*len(lines)

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			n := height(lines[y][x])

			u := up(lines, x, y)
			uv := visible(n, u)
			r := right(lines, x, y)
			rv := visible(n, r)
			d := down(lines, x, y)
			dv := visible(n, d)
			l := left(lines, x, y)
			lv := visible(n, l)
			if uv || rv || dv || lv {
				v += 1
			}
		}
	}

	return v
}

func part2(lines []string) int {
	ss := []int{}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			n := height(lines[y][x])

			u := up(lines, x, y)
			ud := view_dist(n, u)
			r := right(lines, x, y)
			rd := view_dist(n, r)
			d := down(lines, x, y)
			dd := view_dist(n, d)
			l := left(lines, x, y)
			ld := view_dist(n, l)

			ss = append(ss, ud*rd*dd*ld)
		}
	}

	return max(ss)
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}
