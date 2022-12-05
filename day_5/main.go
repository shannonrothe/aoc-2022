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

func (c *crate) join(o *crate) {
	c.stack = append(c.stack, o.stack...)
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
	o.join(&s)
}

func (c *crate) reverse() {
	for i, j := 0, len(c.stack)-1; i < j; i, j = i+1, j-1 {
		c.stack[i], c.stack[j] = c.stack[j], c.stack[i]
	}
}

type command struct {
	n, f, t int
}

func parseMovement(str string) (int, int, int) {
	line := strings.Split(str, " ")
	n, _ := strconv.Atoi(line[1])
	f, _ := strconv.Atoi(line[3])
	t, _ := strconv.Atoi(line[5])
	return n, f - 1, t - 1
}

func parseInput(crates []string, cmd []string) ([9]crate, []command) {
	crs := [9]crate{}
	cmds := []command{}

	for _, line := range crates {
		for j := 0; j < len(line); j += 3 {
			if j != 0 {
				j += 1
			}
			if line[j] == byte('[') {
				crs[j/4].push(line[j+1])
			}
		}
	}

	for _, c := range crs {
		c.reverse()
	}

	for _, cmd := range cmd {
		n, f, t := parseMovement(cmd)
		cmds = append(cmds, command{n, f, t})
	}

	return crs, cmds
}

func top(crs [9]crate) string {
	tops := []string{}
	for _, c := range crs {
		tops = append(tops, string(c.pop()))
	}
	return strings.Join(tops, "")
}

func part1(crs [9]crate, cmds []command) string {
	for _, cmd := range cmds {
		crs[cmd.f].moveIndividualToN(&crs[cmd.t], cmd.n)
	}
	return top(crs)
}

func part2(crs [9]crate, cmds []command) string {
	for _, cmd := range cmds {
		crs[cmd.f].moveToN(&crs[cmd.t], cmd.n)
	}
	return top(crs)
}

func main() {
	parts := strings.Split(input, "\n\n")
	crates, commands := parseInput(strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n"))
	fmt.Printf("part 1: %s\n", part1(crates, commands))

	crates, commands = parseInput(strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n"))
	fmt.Printf("part 2: %s\n", part2(crates, commands))
}
