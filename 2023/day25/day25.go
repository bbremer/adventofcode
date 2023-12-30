package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day25.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	verticies := map[string]*Vertex{}
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ": ")
		splitLine2 := strings.Split(splitLine[1], " ")
		n := GetVertex(verticies, splitLine[0])
		for _, name := range splitLine2 {
			n.AddEdge(verticies, name, 1)
		}
	}

	originalSize := len(verticies)
	for {
		size, score := MinimumCutPhase(verticies)
		if score <= 3 {
			fmt.Println("Part 1:", size*(originalSize-size))
			break
		}
	}
}

func MinimumCutPhase(verticies map[string]*Vertex) (cliqueSize, score int) {
	a := GetAnyKey(verticies)
	lastAdded := verticies[a]
	q := PriorityQueue{}
	for _, v := range verticies {
		if v.name == a {
			continue
		}
		if s, ok := verticies[a].edges[v]; ok {
			q = append(q, &Node{v, s})
		} else {
			q = append(q, &Node{v, 0})
		}
	}
	heap.Init(&q)
	for len(q) > 1 {
		lastAdded = q.PopVertex()
	}

	last := q.PopVertex()
	delete(verticies, lastAdded.name)
	delete(verticies, last.name)
	newV := last.Combine(lastAdded)
	verticies[newV.name] = newV

	for _, s := range newV.edges {
		score += s
	}
	cliqueSize = newV.size

	return
}

func GetAnyKey(m map[string]*Vertex) string {
	for k := range m {
		return k
	}
	fmt.Println("GetAnyValue")
	os.Exit(1)
	return ""
}

func GetVertex(verticies map[string]*Vertex, name string) *Vertex {
	if n, ok := verticies[name]; ok {
		return n
	}
	verticies[name] = &Vertex{name, 1, map[*Vertex]int{}}
	return verticies[name]
}

type Vertex struct {
	name  string
	size  int
	edges map[*Vertex]int
}

func (v *Vertex) AddEdge(verticies map[string]*Vertex, name string, weight int) {
	v2 := GetVertex(verticies, name)
	v.edges[v2] = weight
	v2.edges[v] = weight
}

func (v *Vertex) Combine(v2 *Vertex) *Vertex {
	name := strings.Join([]string{v.name, v2.name}, ",")
	newV := &Vertex{name, v.size + v2.size, map[*Vertex]int{}}
	delete(v.edges, v2)
	delete(v2.edges, v)
	for v3, s := range v.edges {
		delete(v3.edges, v)
		newV.edges[v3] = s
		v3.edges[newV] = s
	}
	for v3, s := range v2.edges {
		delete(v3.edges, v2)
		if _, ok := newV.edges[v3]; ok {
			newV.edges[v3] += s
			v3.edges[newV] += s
		} else {
			newV.edges[v3] = s
			v3.edges[newV] = s
		}
	}
	return newV
}

type Node struct {
	vertex *Vertex
	score  int
}

type PriorityQueue []*Node

/*
func NewPriorityQueue() PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return pq
}
*/

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score > pq[j].score
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

func (pq *PriorityQueue) PopVertex() *Vertex {
	v := heap.Pop(pq).(*Node).vertex
	for _, n := range *pq {
		if s, ok := v.edges[n.vertex]; ok {
			n.score += s
		}
	}
	heap.Init(pq)
	return v
}

/*
func (pq *PriorityQueue) pushPos(p *Vertex, score int) {
	heap.Push(pq, &Node{p, score})
}

func (pq *PriorityQueue) popPos() (*Vertex, int) {
	n := heap.Pop(pq).(*Node)
	return n.vertex, n.score
}
*/
