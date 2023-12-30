package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day16.txt")
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	part1Total := run(lines, 0, -1, NewRight)
	fmt.Println("part 1:", part1Total)

	part2Total := 0
	for i := range lines {
		if r := run(lines, i, -1, NewRight); r > part2Total {
			part2Total = r
		}
		if l := run(lines, i, len(lines[0]), NewLeft); l > part2Total {
			part2Total = l
		}
	}
	for j := range lines[0] {
		if d := run(lines, -1, j, NewDown); d > part2Total {
			part2Total = d
		}
		if u := run(lines, len(lines), j, NewUp); u > part2Total {
			part2Total = u
		}
	}
	fmt.Println("part 2:", part2Total)
}

func run(lines []string, startR, startC int, startF func(int, int, [][]*Node) (Mover, bool)) (count int) {
	nodes := make([][]*Node, 0)
	for _, line := range lines {
		nodeLine := make([]*Node, 0)
		for j := 0; j < len(line); j++ {
			nodeLine = append(nodeLine, &Node{line[j], [4]bool{false}})
		}
		nodes = append(nodes, nodeLine)
	}

	startM, _ := startF(startR, startC, nodes)
	movers := []Mover{startM}
	for len(movers) > 0 {
		newMovers := make([]Mover, 0)
		for _, m := range movers {
			newMs := m.Move(nodes)
			newMovers = append(newMovers, newMs...)
		}
		movers = newMovers
	}

	for _, l := range nodes {
		for _, n := range l {
			for _, d := range n.directions {
				if d {
					count++
					break
				}
			}
		}
	}
	return
}

type Node struct {
	char       byte
	directions [4]bool
}

func ValidNode(r, c int, nodes [][]*Node, directionInt int) bool {
	if 0 <= r && r < len(nodes) && 0 <= c && c < len(nodes[0]) && !nodes[r][c].directions[directionInt] {
		nodes[r][c].directions[directionInt] = true
		return true
	}
	return false
}

type Mover interface {
	Move(nodes [][]*Node) []Mover
}

type Up struct {
	r int
	c int
}

func NewUp(r, c int, nodes [][]*Node) (Mover, bool) {
	return Up{r - 1, c}, ValidNode(r-1, c, nodes, 0)
}

func (m Up) Move(nodes [][]*Node) []Mover {
	fs := make([]func(int, int, [][]*Node) (Mover, bool), 0)
	switch nodes[m.r][m.c].char {
	case '.', '|':
		fs = append(fs, NewUp)
	case '\\':
		fs = append(fs, NewLeft)
	case '/':
		fs = append(fs, NewRight)
	case '-':
		fs = append(fs, NewLeft, NewRight)
	default:
		fmt.Println("Problem in up")
		os.Exit(1)
	}
	movers := make([]Mover, 0)
	for _, f := range fs {
		if n, ok := f(m.r, m.c, nodes); ok {
			movers = append(movers, n)
		}
	}
	return movers
}

type Down struct {
	r int
	c int
}

func NewDown(r, c int, nodes [][]*Node) (Mover, bool) {
	return Down{r + 1, c}, ValidNode(r+1, c, nodes, 1)
}

func (m Down) Move(nodes [][]*Node) []Mover {
	fs := make([]func(int, int, [][]*Node) (Mover, bool), 0)
	switch nodes[m.r][m.c].char {
	case '.', '|':
		fs = append(fs, NewDown)
	case '\\':
		fs = append(fs, NewRight)
	case '/':
		fs = append(fs, NewLeft)
	case '-':
		fs = append(fs, NewLeft, NewRight)
	default:
		fmt.Println("Problem in down")
		os.Exit(1)
	}

	movers := make([]Mover, 0)
	for _, f := range fs {
		if n, ok := f(m.r, m.c, nodes); ok {
			movers = append(movers, n)
		}
	}
	return movers
}

type Left struct {
	r int
	c int
}

func NewLeft(r, c int, nodes [][]*Node) (Mover, bool) {
	return Left{r, c - 1}, ValidNode(r, c-1, nodes, 2)
}

func (m Left) Move(nodes [][]*Node) []Mover {
	fs := make([]func(int, int, [][]*Node) (Mover, bool), 0)
	switch nodes[m.r][m.c].char {
	case '.', '-':
		fs = append(fs, NewLeft)
	case '\\':
		fs = append(fs, NewUp)
	case '/':
		fs = append(fs, NewDown)
	case '|':
		fs = append(fs, NewUp, NewDown)
	default:
		fmt.Println("Problem in left")
		os.Exit(1)
	}
	movers := make([]Mover, 0)
	for _, f := range fs {
		if n, ok := f(m.r, m.c, nodes); ok {
			movers = append(movers, n)
		}
	}
	return movers
}

type Right struct {
	r int
	c int
}

func NewRight(r, c int, nodes [][]*Node) (Mover, bool) {
	return Right{r, c + 1}, ValidNode(r, c+1, nodes, 3)
}

func (m Right) Move(nodes [][]*Node) []Mover {
	fs := make([]func(int, int, [][]*Node) (Mover, bool), 0)
	switch nodes[m.r][m.c].char {
	case '.', '-':
		fs = append(fs, NewRight)
	case '\\':
		fs = append(fs, NewDown)
	case '/':
		fs = append(fs, NewUp)
	case '|':
		fs = append(fs, NewUp, NewDown)
	default:
		fmt.Println("Problem in right")
		os.Exit(1)
	}
	movers := make([]Mover, 0)
	for _, f := range fs {
		if n, ok := f(m.r, m.c, nodes); ok {
			movers = append(movers, n)
		}
	}
	return movers
}
