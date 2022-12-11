package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day_11.txt
var input string

func parseStartingItems(l string) []int {
	s := strings.Split(l, ": ")
	s = strings.Split(s[1], ",")
	si := []int{}
	for _, v := range s {
		n, _ := strconv.Atoi(strings.Trim(v, " "))
		si = append(si, n)
	}
	return si
}

func parseExpr(l string) expr {
	s := strings.Split(l, ": ")
	s = strings.Split(s[1], "= ")
	s = strings.Split(s[1], " ")
	lhs := s[0]
	op := s[1]
	rhs := s[2]
	return expr{
		lhs: lhs,
		op:  op,
		rhs: rhs,
	}
}

func parseTest(l []string) test {
	s := strings.Split(l[0], "by ")
	cond, _ := strconv.Atoi(s[1])

	t := strings.Split(l[1], "monkey ")
	then, _ := strconv.Atoi(strings.Trim(t[1], " "))

	o := strings.Split(l[2], "monkey ")
	otherwise, _ := strconv.Atoi(strings.Trim(o[1], " "))

	return test{
		cond:      cond,
		then:      then,
		otherwise: otherwise,
	}
}

type expr struct {
	lhs string
	op  string
	rhs string
}

type test struct {
	cond      int
	then      int
	otherwise int
}

type monkey struct {
	name string
	si   []int
	op   expr
	t    test
	ins  int
}

func eval(lhs, op, rhs string, w int) int {
	switch op {
	case "*":
		l, _ := strconv.Atoi(lhs)
		r, _ := strconv.Atoi(rhs)
		if lhs == "old" && rhs == "old" {
			return w * w
		} else if lhs == "old" && rhs != "old" {
			return w * r
		} else {
			return l * r
		}
	case "+":
		l, _ := strconv.Atoi(lhs)
		r, _ := strconv.Atoi(rhs)
		if lhs == "old" && rhs == "old" {
			return w + w
		} else if lhs == "old" && rhs != "old" {
			return w + r
		} else {
			return l + r
		}
	default:
		return 0
	}
}

func (m *monkey) operate(it int) int {
	return eval(m.op.lhs, m.op.op, m.op.rhs, it)
}

func (m *monkey) test(it int) int {
	if it%m.t.cond == 0 {
		return m.t.then
	} else {
		return m.t.otherwise
	}
}

func (m *monkey) inspect(mks []*monkey) {
	for len(m.si) > 0 {
		m.ins++

		it := m.si[0]
		it = m.operate(it) / 3
		to := m.test(it)
		mks[to].si = append(mks[to].si, it)
		m.si = m.si[1:]
	}
}

func (m *monkey) inspect2(mks []*monkey, mod int) {
	for len(m.si) > 0 {
		m.ins++

		it := m.si[0]
		it = m.operate(it) % mod
		to := m.test(it)

		mks[to].si = append(mks[to].si, it)
		m.si = m.si[1:]
	}
}

func parse(input string) []*monkey {
	sections := strings.Split(input, "\n\n")
	monkeys := []*monkey{}

	for i, s := range sections {
		l := strings.Split(s, "\n")
		si := parseStartingItems(l[1])
		op := parseExpr(l[2])
		test := parseTest(l[3:6])

		monkeys = append(monkeys, &monkey{
			name: fmt.Sprintf("Monkey %d", i),
			si:   si,
			op:   op,
			t:    test,
			ins:  0,
		})
	}

	return monkeys
}

func part1() int {
	monkeys := parse(input)

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.inspect(monkeys)
		}
	}

	ins := []int{}
	for _, m := range monkeys {
		ins = append(ins, m.ins)
	}
	sort.Ints(ins)
	m := ins[len(ins)-2:]
	return m[0] * m[1]
}

func part2() int {
	monkeys := parse(input)

	mod := 1
	for _, mk := range monkeys {
		mod = mod * mk.t.cond
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.inspect2(monkeys, mod)
		}
	}

	ins := []int{}
	for _, m := range monkeys {
		ins = append(ins, m.ins)
	}
	sort.Ints(ins)
	t := ins[len(ins)-2:]
	return t[0] * t[1]
}

func main() {
	fmt.Printf("part 1: %d\n", part1())
	fmt.Printf("part 2: %d\n", part2())
}
