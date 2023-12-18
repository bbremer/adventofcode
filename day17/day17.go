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

	p := Pos{0, 0, 0, direction{0, 0}, 0}
	part1Total := run(lines, p, move)
	fmt.Println("part 1:", part1Total)
}

func run(lines [][]int, startP Pos, moveF func(Pos, direction, [][]int) (Pos, bool)) int {
	scores := map[Pos]int{}
	pq := make(PriorityQueue, 0)
	pq.pushPos(startP, 0)
	for pq.Len() > 0 {
		p, score := pq.popPos()

		if !p.valid(lines) {
			continue
		}

		if p.r == len(lines)-1 && p.c == len(lines[0])-1 {
			return score
		}

		if _, ok := scores[p]; ok {
			continue
		}
		scores[p] = score

		for _, d := range []direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			if newP, ok := moveF(p, d, lines); ok {
				newScore := score + lines[newP.r][newP.c]
				pq.pushPos(newP, newScore)
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
	return abs(d1.r-d2.r) == 2 || abs(d1.c-d2.c) == 2
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
	distance    int
}

func (p Pos) valid(lines [][]int) bool {
	return 0 <= p.r && p.r < len(lines) && 0 <= p.c && p.c < len(lines[0])
}

func move(p Pos, d direction, lines [][]int) (Pos, bool) {
	consecutive := 0
	if d.Eq(p.d) {
		consecutive = p.consecutive + 1
	}
	newP := Pos{p.r + d.r, p.c + d.c, consecutive, d, 1}
	isValid := newP.valid(lines) && consecutive < 3 && !d.backwards(p.d)
	return newP, isValid
}

/*
func move2(p, Pos, s int, d direction, lines [][]int) (Pos, bool) {
	var consecutive int
	if d.Eq(p.d) {
		consecutive = p.consecutive + 1
		newP := Pos{p.r + d.r, p.c + d.c, consecutive, d, 1}
		return newP, newP.valid(lines) && consecutive < 10 && !d.backwards(p.d)
	}
	consecutive = 4
	newP := Pos{p.r + 4*d.r, p.c + 4*d.c, consecutive, d, 4}
	return newP, newP.valid(lines) && consecutive < 10 && !d.backwards(p.d)
}
*/

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
