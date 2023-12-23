package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day21.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]byte, 0)
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		lines = append(lines, lineCopy)
	}

	part1Lines := copyLines(lines)
	fmt.Println("Part 1:", run(part1Lines, 64, 65))

	oddLines := copyLines(lines)
	run(oddLines, 10001, 65)

	/*
		evenLines := copyLines(lines)
		evenTotal := run(evenLines, 10000, 65)
	*/

	lim := 65
	o1 := 0
	e1 := 0
	oddSum := 0
	evenSum := 0
	for r, line := range oddLines {
		for c, b := range line {
			// if b != '.' && b != '#' {
			if b == 'S' || b == 'O' || b == 'E' {
				if (r+c)%2 == 1 {
					oddSum++
					if r+c < lim || r-c < -lim || c-r < -lim || r+c > lim*3 {
						oddLines[r][c] = 'x'
						o1++
					}
				} else {
					evenSum++
					if r+c <= lim || r-c <= -lim || c-r <= -lim || r+c >= lim*3 {
						oddLines[r][c] = 'x'
						e1++
					}
				}
			}
		}
	}

	n := 202300
	// n := 2
	numO := int(math.Pow(float64(n+1), 2))
	numE := int(math.Pow(float64(n), 2))
	// fmt.Println(oddSum, evenSum, o1, e1)
	part2Total := numO*oddSum + numE*evenSum + n*e1 - (n+1)*o1

	/*
		lines[65][65] = '.'
		newLines := make([][]byte, 0, n*2+1)
		for i := 0; i < n*2+1; i++ {
			for _, line := range lines {
				newLine := make([]byte, 0, n*2+1)
				for j := 0; j < n*2+1; j++ {
					for _, c := range line {
						newLine = append(newLine, c)
					}
				}
				newLines = append(newLines, newLine)
			}
		}
		newLines[n*131+65][n*131+65] = 'S'
		fmt.Println("Part 2 real:", run(newLines, n*131+65, n*131+65))
	*/
	/*
		for _, line := range oddLines {
			fmt.Println(string(line))
		}
	*/
	/*
		newSum := 0
		for _, line := range newLines {
			// fmt.Println(string(line))
			for _, b := range line {
				if b == 'O' {
					newSum++
				}
			}
		}
		fmt.Println(newSum)
	*/

	fmt.Println("Part 2:", part2Total)
}

func copyLines(lines [][]byte) [][]byte {
	newLines := make([][]byte, len(lines))
	for i, line := range lines {
		newLines[i] = make([]byte, len(line))
		copy(newLines[i], line)
	}
	return newLines
}

func run(lines [][]byte, numSteps, startI int) (count int) {
	parity := numSteps % 2
	q := NewQueue(&Pos{startI, startI})
	for !q.Empty() {
		p, c := q.Pop()
		if c > numSteps {
			continue
		}
		if (p.r+p.c)%2 == parity {
			lines[p.r][p.c] = 'O'
			count++
		} else {
			lines[p.r][p.c] = 'E'
		}
		for _, d := range [4]Pos{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			if p2 := p.Add(&d); p2.IsValid(lines) {
				q.Push(p2, c+1)
				lines[p2.r][p2.c] = 'x'
				/*
					if parity != c%2 {
						count++
					}
				*/
			}
		}
	}

	return count
}

type Pos struct {
	r int
	c int
}

func (p *Pos) Add(d *Pos) *Pos {
	return &Pos{p.r + d.r, p.c + d.c}
}

func (p *Pos) IsValid(lines [][]byte) bool {
	return p.r >= 0 && p.r < len(lines) && p.c >= 0 && p.c < len(lines[0]) && lines[p.r][p.c] == '.'
}

type Node struct {
	p *Pos
	n *Node
	c int
}

type Queue struct {
	h *Node
	t *Node
}

func NewQueue(p *Pos) Queue {
	n := &Node{p, nil, 0}
	return Queue{n, n}
}

func (q Queue) Empty() bool {
	return q.h == nil
}

func (q *Queue) Pop() (*Pos, int) {
	p := q.h.p
	c := q.h.c
	q.h = q.h.n
	if q.h == nil {
		q.t = nil
	}
	return p, c
}

func (q *Queue) Push(p *Pos, c int) {
	n := &Node{p, nil, c}
	if q.t == nil {
		q.t = n
		q.h = n
		return
	}
	q.t.n = n
	q.t = n
}
