package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed day_9.txt
var input string

type step struct {
	dir string
	amt int
}

type knot struct {
	pts  []point
	curr point
}

func newKnot() knot {
	return knot{
		pts:  []point{},
		curr: newPoint(),
	}
}

func (k *knot) has(p point) bool {
	h := false
	for _, pt := range k.pts {
		h = pt.x == p.x && pt.y == p.y
		if h {
			break
		}
	}
	return h
}

func (k *knot) record() {
	if !k.has(k.curr) {
		k.pts = append(k.pts, k.curr)
	}
}

type point struct {
	x, y int
}

func newPoint() point {
	return point{
		x: 0,
		y: 0,
	}
}

func (p *point) realign(o *point, dir string) {
	switch dir {
	case "U":
		p.x = o.x
		p.y = o.y + 1
	case "R":
		p.x = o.x - 1
		p.y = o.y
	case "D":
		p.x = o.x
		p.y = o.y - 1
	case "L":
		p.x = o.x + 1
		p.y = o.y
	}
}

func (p *point) move(dir string) {
	switch dir {
	case "U":
		p.y--
	case "R":
		p.x++
	case "D":
		p.y++
	case "L":
		p.x--
	}
}

func (p *point) dist(o *point) int {
	dx := o.x - p.x
	dy := o.y - p.y
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func (p *point) diag(o *point) bool {
	return (p.x == o.x-1 && p.y == o.y-1) || (p.x == o.x+1 && p.y == o.y-1) || (p.x == o.x+1 && p.y == o.y+1) || (p.x == o.x-1 && p.y == o.y+1)
}

func (p *point) sub(o *point) point {
	return point{
		x: p.x - o.x,
		y: p.y - o.y,
	}
}

func (k *knot) move(dir string) {
	if dir == "U" {
		k.curr.y--
	}
	if dir == "R" {
		k.curr.x++
	}
	if dir == "D" {
		k.curr.y++
	}
	if dir == "L" {
		k.curr.x--
	}
}

func part1(s []step) int {
	h := newPoint()
	t := newKnot()
	for _, st := range s {
		for i := 0; i < st.amt; i++ {
			b := h
			h.move(st.dir)

			d := h.dist(&t.curr)
			if d > 1 {
				if t.curr.diag(&h) {
					t.curr.realign(&h, st.dir)
				} else {
					t.curr = b
					t.record()
				}
			} else {
				t.record()
			}
		}
	}
	return len(t.pts)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *point) moveTail(p2 point) {
	d := p.sub(&p2)

	if abs(d.x) <= 1 && abs(d.y) <= 1 {
		return
	}

	if d.y == 0 {
		if d.x > 0 {
			p.x--
		} else {
			p.x++
		}
	} else if d.x == 0 {
		if d.y > 0 {
			p.y--
		} else {
			p.y++
		}
	} else {
		if d.x > 0 && d.y > 0 {
			p.x--
			p.y--
		} else if d.x > 0 && d.y < 0 {
			p.x--
			p.y++
		} else if d.x < 0 && d.y > 0 {
			p.x++
			p.y--
		} else if d.x < 0 && d.y < 0 {
			p.x++
			p.y++
		}
	}
}

type trail map[point]struct{}

func (t trail) update(tail point) {
	if _, ok := t[tail]; !ok {
		t[tail] = struct{}{}
	}
}

func part2(s []step) int {
	t := []knot{}
	for i := 0; i < 10; i++ {
		k := newKnot()
		k.record() // record s
		t = append(t, k)
	}
	tr := trail{point{x: 0, y: 0}: struct{}{}}
	for _, st := range s {
		for step := 0; step < st.amt; step++ {
			t[0].move(st.dir)
			for i := range t[1:] {
				t[i+1].curr.moveTail(t[i].curr)
			}
			tr.update(t[9].curr)
		}
	}
	return len(tr)
}

func parse(input string) []step {
	lines := strings.Split(input, "\n")
	s := []step{}
	for _, l := range lines {
		sp := strings.Split(l, " ")
		n, _ := strconv.Atoi(sp[1])
		s = append(s, step{
			dir: sp[0],
			amt: n,
		})
	}
	return s
}

func main() {
	s := parse(input)
	fmt.Printf("part 1: %d\n", part1(s))
	fmt.Printf("part 2: %d\n", part2(s))
}
