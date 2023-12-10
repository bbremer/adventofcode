package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("../inputs/day10.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make([][]*Node, 0)
	for y := 0; scanner.Scan(); y++ {
		graph = append(graph, parse(scanner.Text(), y))
	}

	part1Total := maxDepth(graph)

	part2Total := 0
	for _, line := range graph {
		parity := false
		for _, n := range line {
			if n.Visited {
				if n.symbol == '|' || n.symbol == 'J' || n.symbol == 'L' {
					parity = !parity
				}
			} else if parity {
				part2Total++
			}
		}
	}

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

func parse(line string, y int) []*Node {
	nodes := make([]*Node, 0)
	for x := 0; x < len(line); x++ {
		nodes = append(nodes, NewNode(line[x], x, y))
	}
	return nodes
}

var startLocation [2]int

type Node struct {
	symbol    byte
	Neighbors [][2]int
	Visited   bool
}

func NewNode(symbol byte, x, y int) (n *Node) {
	neighbors := make([][2]int, 0)

	directions := symbolToDirections(symbol)
	if directions[0] {
		neighbors = append(neighbors, [2]int{x, y - 1})
	}
	if directions[1] {
		neighbors = append(neighbors, [2]int{x, y + 1})
	}
	if directions[2] {
		neighbors = append(neighbors, [2]int{x - 1, y})
	}
	if directions[3] {
		neighbors = append(neighbors, [2]int{x + 1, y})
	}

	n = &Node{symbol, neighbors, false}

	if symbol == 'S' {
		startLocation = [2]int{x, y}
		n.Visited = true
	}

	return
}

func symbolToDirections(symbol byte) [4]bool {
	var up, down, left, right bool
	if symbol == 'S' {
		return [4]bool{up, down, left, right}
	}
	switch symbol {
	case '|':
		up = true
		down = true
	case '-':
		left = true
		right = true
	case 'L':
		up = true
		right = true
	case 'J':
		up = true
		left = true
	case '7':
		left = true
		down = true
	case 'F':
		right = true
		down = true
	}
	return [4]bool{up, down, left, right}
}

func maxDepth(g [][]*Node) (depth int) {
	queue := NewQueue(g)
	for !queue.IsEmpty() {
		n, d := queue.Pop()
		if d > depth {
			depth = d
		}
		for _, newLoc := range n.Neighbors {
			x := newLoc[0]
			y := newLoc[1]
			newN := g[y][x]
			if !newN.Visited {
				newN.Visited = true
				queue.Push(newN, d+1)
			}
		}
	}
	return
}

type QueueNode struct {
	N        *Node
	NextN    *QueueNode
	Distance int
}

type Queue struct {
	firstNode *QueueNode
	lastNode  *QueueNode
}

func NewQueue(g [][]*Node) *Queue {
	q := &Queue{}
	directions := [4]string{"up", "down", "left", "right"}
	for _, direction := range directions {
		if n, ok := isStartNeighbor(g, direction); ok {
			q.Push(n, 1)
			n.Visited = true
		}
	}
	return q
}

func isStartNeighbor(g [][]*Node, direction string) (*Node, bool) {
	var x, y, index int
	switch direction {
	case "up":
		x = startLocation[0]
		y = startLocation[1] - 1
		index = 1
	case "down":
		x = startLocation[0]
		y = startLocation[1] + 1
		index = 0
	case "left":
		x = startLocation[0] - 1
		y = startLocation[1]
		index = 3
	case "right":
		x = startLocation[0] + 1
		y = startLocation[1]
		index = 2
	}
	if symbolToDirections(g[y][x].symbol)[index] {
		return g[y][x], true
	}

	return nil, false
}

func indexValid(g [][]*Node, x, y int) bool {
	if y < 0 || len(g) <= y {
		return false
	}
	if x < 0 || len(g[y]) <= x {
		return false
	}
	return true
}

func (q *Queue) IsEmpty() bool {
	return q.firstNode == nil
}

func (q *Queue) Push(n *Node, d int) {
	qn := &QueueNode{n, nil, d}
	if q.IsEmpty() {
		q.firstNode = qn
		q.lastNode = qn
		return
	}
	q.lastNode.NextN = qn
	q.lastNode = qn
}

func (q *Queue) Pop() (n *Node, d int) {
	n = q.firstNode.N
	d = q.firstNode.Distance
	if nextN := q.firstNode.NextN; nextN != nil {
		q.firstNode = nextN
	} else {
		q.firstNode = nil
		q.lastNode = nil
	}
	return n, d
}
