package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day_7.txt
var input string

type file struct {
	name string
	size int
}

type directory struct {
	name   string
	parent *directory
	dirs   []*directory
	files  []*file
}

func (d *directory) root() *directory {
	c := d
	for c.parent != nil {
		c = c.parent
	}
	return c
}

func (d *directory) find(name string) *directory {
	for _, d2 := range d.dirs {
		if d2.name == name {
			return d2
		} else {
			d2.find(name)
		}
	}
	return d.parent
}

func (d *directory) dirSize() int {
	s := 0
	for _, f := range d.files {
		s += f.size
	}
	for _, d2 := range d.dirs {
		s += d2.dirSize()
	}
	return s
}

func (d *directory) sum() int {
	s := 0
	sz := d.dirSize()
	if sz <= 100000 {
		s += sz
	}
	for _, d2 := range d.dirs {
		s += d2.sum()
	}
	return s
}

func (d *directory) findSmallest(sn int) []int {
	dl := []int{}
	sz := d.dirSize()
	if sz >= sn {
		dl = append(dl, sz)
	}
	for _, d2 := range d.dirs {
		dl = append(dl, d2.findSmallest(sn)...)
	}
	return dl
}

func newDir(name string, parent *directory) *directory {
	return &directory{
		name:   name,
		parent: parent,
		dirs:   []*directory{},
		files:  []*file{},
	}
}

func newFs(name string, lines []string) *directory {
	d := newDir(name, nil)

	for _, l := range lines {
		switch l[0] {
		case '$':
			if l[2] == 'c' { // cd
				s := strings.Split(l, " ")
				name := s[2]

				switch name {
				case "/":
					continue
				default:
					d = d.find(name)
				}
			}
		default:
			s := strings.Split(l, " ")
			if l[0] == 'd' { // dir
				d.dirs = append(d.dirs, newDir(s[1], d))
			} else { // file
				size, _ := strconv.Atoi(s[0])
				name := s[1]

				d.files = append(d.files, &file{
					name: name,
					size: size,
				})
			}
		}
	}
	return d
}

func part1(lines []string) int {
	return newFs("/", lines).root().sum()
}

func part2(lines []string) int {
	d := newFs("/", lines).root()
	sn := 30000000 - (70000000 - d.dirSize())
	dl := d.findSmallest(sn)
	sort.Ints(dl)
	return dl[0]
}

func main() {
	lines := strings.Split(input, "\n")
	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}
