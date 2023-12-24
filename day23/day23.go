package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day23.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) (max int) {
	startR, startC := 0, 1
	endR, endC := len(lines)-1, len(lines[0])-2
	if lines[startR][startC] == '#' {
		return 0
	}

	stack := []Node{NewNode(startR, startC, len(lines))}
	var n Node

	for {
		if len(stack) <= 0 {
			break
		}
		stack, n = stack[:len(stack)-1], stack[len(stack)-1]
		if n.r == endR && n.c == endC {
			if n.d > max {
				max = n.d
			}
			continue
		}
		a := Part1Switch(lines[n.r][n.c])
		for _, p := range a {
			r, c := n.r+p[0], n.c+p[1]
			if Valid(r, c, lines) && !n.s.Visited(r, c) {
				s, d := n.s.Visit(r, c), n.d+1
				stack = append(stack, Node{r, c, s, d})
			}
		}
	}

	return
}

func Part1Switch(b byte) [][2]int {
	var a [][2]int
	switch b {
	case '>':
		a = [][2]int{{0, 1}}
	case '<':
		a = [][2]int{{0, -1}}
	case 'v':
		a = [][2]int{{1, 0}}
	case '^':
		a = [][2]int{{-1, 0}}
	default:
		a = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	}
	return a
}

func Valid(r, c int, lines []string) bool {
	return 0 <= r && r < len(lines) && 0 <= c && c < len(lines[0]) && lines[r][c] != '#'
}

type Node struct {
	r int
	c int
	s MapState
	d int
}

func NewNode(r, c, numRows int) Node {
	s := NewMapState(numRows).Visit(r, c)
	return Node{r, c, s, 0}
}

type MapState struct {
	rows []Row
}

func NewMapState(numRows int) MapState {
	s := MapState{make([]Row, numRows)}
	for i := range s.rows {
		s.rows[i] = Row{[3]uint64{0, 0, 0}}
	}
	return s
}

func (s MapState) Visit(r, c int) MapState {
	newMapState := MapState{make([]Row, len(s.rows))}
	copy(newMapState.rows, s.rows)
	newMapState.rows[r] = s.rows[r].Visit(c)
	return newMapState
}

func (s MapState) Visited(r, c int) bool {
	return s.rows[r].Visited(c)
}

type Row struct {
	visited [3]uint64
}

func (r Row) Visit(c int) Row {
	i := c / 64
	newR := Row{[3]uint64{r.visited[0], r.visited[1], r.visited[2]}}
	newR.visited[i] = r.visited[i] | 1<<(c%64)
	return newR
}

func (r Row) Visited(c int) bool {
	return r.visited[c/64]&(1<<(c%64)) != 0
}

func Part2(lines []string) (max int) {
	intersections := []Intersection{
		NewIntersection(0, 1, lines),
		NewIntersection(len(lines)-1, len(lines[0])-2, lines)}
	for r, line := range lines[1 : len(lines)-1] {
		for c, b := range line[1 : len(line)-1] {
			if b == '#' {
				continue
			}
			if IsIntersection(r+1, c+1, lines) {
				intersections = append(intersections, NewIntersection(r+1, c+1, lines))
			}
		}
	}
	intersectionMap := map[[2]int]int{}
	for j, i := range intersections {
		intersectionMap[[2]int{i.r, i.c}] = j
	}
	stack := []Node2{{intersections[intersectionMap[[2]int{0, 1}]], 1, 0}}
	var n Node2
	for len(stack) > 0 {
		stack, n = stack[:len(stack)-1], stack[len(stack)-1]
		if n.i.r == len(lines)-1 && n.i.c == len(lines[0])-2 {
			if n.d > max {
				max = n.d
			}
			continue
		}
		for _, e := range n.i.edges {
			newIIndex := intersectionMap[e.dest]
			c := uint64(1) << newIIndex
			if n.s&c == 0 {
				stack = append(stack, Node2{intersections[newIIndex], n.s | c, n.d + e.dist})
			}
		}
	}
	return
}

type Node2 struct {
	i Intersection
	s uint64
	d int
}

func IsIntersection(r, c int, lines []string) bool {
	if (r == 0 && c == 1) || (r == len(lines)-1 && c == len(lines[0])-2) {
		return true
	}
	count := 0
	for _, i := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		if lines[r+i[0]][c+i[1]] != '#' {
			count++
		}
	}
	return count > 2
}

type Intersection struct {
	r     int
	c     int
	edges []Edge
}

func NewIntersection(iR, iC int, lines []string) Intersection {
	stack := []Node{}
	ms := NewMapState(len(lines)).Visit(iR, iC)
	for _, j := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		r, c := iR+j[0], iC+j[1]
		if Valid(r, c, lines) {
			stack = append(stack, Node{r, c, ms.Visit(r, c), 1})
		}
	}

	edges := []Edge{}
	var n Node
	for len(stack) > 0 {
		stack, n = stack[:len(stack)-1], stack[len(stack)-1]
		if IsIntersection(n.r, n.c, lines) {
			edges = append(edges, Edge{[2]int{n.r, n.c}, n.d})
			continue
		}
		for _, p := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			r, c := n.r+p[0], n.c+p[1]
			if Valid(r, c, lines) && !n.s.Visited(r, c) {
				s, d := n.s.Visit(r, c), n.d+1
				stack = append(stack, Node{r, c, s, d})
			}
		}
	}
	return Intersection{iR, iC, edges}
}

type Edge struct {
	dest [2]int
	dist int
}
