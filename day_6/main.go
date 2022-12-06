package main

import (
	_ "embed"
	"fmt"
)

//go:embed day_6.txt
var input string

func setOf(chars string) string {
	s := map[rune]bool{}
	for _, c := range chars {
		_, ok := s[c]
		if !ok {
			s[c] = true
		}
	}
	keys := make([]rune, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return string(keys)
}

func part1(packet string) int {
	for i := 0; i < len(packet); i++ {
		p := packet[i : i+4]
		if len(p) == len(setOf(p)) {
			return i + 4
		}
	}
	panic("didn't find start of message marker")
}

func part2(packet string) int {
	for i := 0; i < len(packet); i++ {
		p := packet[i : i+14]
		if len(p) == len(setOf(p)) {
			return i + 14
		}
	}
	panic("didn't find start of message marker")
}

func main() {
	fmt.Printf("part 1: %d\n", part1(input))
	fmt.Printf("part 2: %d\n", part2(input))
}
