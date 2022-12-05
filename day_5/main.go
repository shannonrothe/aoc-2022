package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day_5.txt
var input string

type crate struct {
	stack []byte
}

func (c *crate) push(b byte) {
	c.stack = append(c.stack, b)
}

func (c *crate) pop() byte {
	if len(c.stack) > 0 {
		b := c.stack[len(c.stack)-1]
		c.stack = c.stack[:len(c.stack)-1]
		return b
	}
	return 0
}

func (c *crate) moveIndividualToN(o *crate, n int) {
	for i := 0; i < n; i++ {
		b := c.pop()
		o.push(b)
	}
}

func (c *crate) moveToN(o *crate, n int) {
	s := crate{
		stack: []byte{},
	}
	for i := 0; i < n; i++ {
		s.stack = append(s.stack, c.pop())
	}
	s.reverse()
	for _, i := range s.stack {
		o.push(i)
	}
}

func (c *crate) reverse() {
	for i, j := 0, len(c.stack)-1; i < j; i, j = i+1, j-1 {
		c.stack[i], c.stack[j] = c.stack[j], c.stack[i]
	}
}

type command struct {
	n, f, t int
}

func parseInput(lines []string) ([9]crate, []command) {
	crates := [9]crate{}
	commands := []command{}

	for i := 0; i < 8; i++ {
		line := lines[i]

		for j := 0; j < len(line); j += 3 {
			if j != 0 {
				j += 1
			}

			if string(line[j]) == "[" {
				crates[j/4].push(line[j+1])
			}
		}
	}

	for _, c := range crates {
		c.reverse()
	}

	for i := 10; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		n, _ := strconv.Atoi(line[1])
		f, _ := strconv.Atoi(line[3])
		t, _ := strconv.Atoi(line[5])
		commands = append(commands, command{
			n: n,
			f: f - 1,
			t: t - 1,
		})
	}

	return crates, commands
}

func part1(lines []string) string {
	crates, commands := parseInput(lines)
	for _, cmd := range commands {
		crates[cmd.f].moveIndividualToN(&crates[cmd.t], cmd.n)
	}

	tops := []string{}
	for _, c := range crates {
		tops = append(tops, string(c.pop()))
	}
	return strings.Join(tops, "")
}

func part2(lines []string) string {
	crates, commands := parseInput(lines)
	for _, cmd := range commands {
		crates[cmd.f].moveToN(&crates[cmd.t], cmd.n)
	}

	tops := []string{}
	for _, c := range crates {
		tops = append(tops, string(c.pop()))
	}
	return strings.Join(tops, "")
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Printf("part 1: %s\n", part1(lines))
	fmt.Printf("part 2: %s\n", part2(lines))
}
