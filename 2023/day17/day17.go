package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day17.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]int, 0)
	for scanner.Scan() {
		line := make([]int, 0)
		for _, c := range scanner.Bytes() {
			i, _ := strconv.Atoi(string(c))
			line = append(line, i)
		}
		lines = append(lines, line)
	}

	part1Total := run(lines, 1, 3)
	fmt.Println("part 1:", part1Total)

	part2Total := run(lines, 4, 10)
	fmt.Println("part 2:", part2Total)
}

func run(lines [][]int, minC, maxC int) int {
	moveF := makeMoveF(lines, minC, maxC)

	startP := Pos{0, 0, 0, direction{0, 0}}

	scores := make(map[Pos]int, len(lines)*len(lines[0])*maxC*4)
	for i, l := range lines {
		for j := range l {
			for _, d := range []direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				for l := 0; l < maxC; l++ {
					scores[Pos{i, j, l, d}] = 1000000000
				}
			}
		}
	}
	scores[startP] = 0

	pq := NewPriorityQueue()
	pq.pushPos(startP, 0)
	for pq.Len() > 0 {
		p, score := pq.popPos()

		if p.r == len(lines)-1 && p.c == len(lines[0])-1 {
			return score
		}

		for _, d := range []direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			if newP, newScore, ok := moveF(p, score, d); ok {
				if newScore < scores[newP] {
					scores[newP] = newScore
					pq.pushPos(newP, newScore)
				}
			}
		}
	}

	fmt.Println("Got through pq")
	os.Exit(1)
	return 0
}

type direction struct {
	r int
	c int
}

func (d1 direction) Eq(d2 direction) bool {
	return d1.r == d2.r && d1.c == d2.c
}

func (d1 direction) backwards(d2 direction) bool {
	return abs(d1.r-d2.r) >= 2 || abs(d1.c-d2.c) >= 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Pos struct {
	r           int
	c           int
	consecutive int
	d           direction
}

func (p Pos) valid(lines [][]int) bool {
	return 0 <= p.r && p.r < len(lines) && 0 <= p.c && p.c < len(lines[0])
}

func makeMoveF(lines [][]int, minC, maxC int) func(p Pos, s int, d direction) (Pos, int, bool) {
	return func(p Pos, s int, d direction) (Pos, int, bool) {
		var consecutive int
		var newP Pos
		if d.Eq(p.d) {
			consecutive = p.consecutive + 1
			newP = Pos{p.r + d.r, p.c + d.c, consecutive, d}
		} else {
			consecutive = minC - 1
			newP = Pos{p.r + minC*d.r, p.c + minC*d.c, consecutive, d}
		}

		if newP.valid(lines) && consecutive < maxC && !d.backwards(p.d) {
			r := p.r
			c := p.c
			newScore := s
			for r != newP.r || c != newP.c {
				r += d.r
				c += d.c
				newScore += lines[r][c]
			}
			return newP, newScore, true
		}
		return newP, 0, false
	}
}

type Node struct {
	p     Pos
	score int
}

type PriorityQueue []*Node

func NewPriorityQueue() PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return pq
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score <= pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) pushPos(p Pos, score int) {
	heap.Push(pq, &Node{p, score})
}

func (pq *PriorityQueue) popPos() (Pos, int) {
	n := heap.Pop(pq).(*Node)
	return n.p, n.score
}
