package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day20.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	modulesMap := make(map[string]Module, 0)
	moduleNames := make([]string, 0)
	conjunctions := make(map[string]*Conjunction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		name, mod, conjunction := NewModule(line)
		modulesMap[name] = mod
		moduleNames = append(moduleNames, name)
		if conjunction {
			m, _ := mod.(*Conjunction)
			conjunctions[name] = m
		}
	}

	for n, m := range modulesMap {
		for _, d := range m.getDestinations() {
			if c, ok := conjunctions[d]; ok {
				c.inputStates[n] = false
				c.inputNames = append(c.inputNames, n)
			}
		}
	}

	jsClique, jsCache := getClique(modulesMap, "mc", "js")
	jsDone := false
	zbClique, zbCache := getClique(modulesMap, "fx", "zb")
	zbDone := false
	rrClique, rrCache := getClique(modulesMap, "nd", "rr")
	rrDone := false
	bsClique, bsCache := getClique(modulesMap, "lf", "bs")
	bsDone := false

	numPushes := 4023
	lows := 0
	highs := 0

	// fmt.Println(state(modules))

	part2Total := 0
	buttonPush := Pulse{false, "button", "broadcaster"}
	for i := 0; i < numPushes; i++ {
		q := NewQueue(buttonPush)
		for !q.Empty() {
			p := q.Pop()
			if i < 1000 {
				if p.p {
					highs++
				} else {
					lows++
				}
			}
			// fmt.Println(p.src, "-", p.p, "->", p.dest)
			if dm, ok := modulesMap[p.dest]; ok {
				for _, np := range dm.pulse(p.p, p.src) {
					q.Push(np)

				}
			} else if !p.p {
				part2Total = i + 1
			}

			if !jsDone && p.src == "js" && p.p {
				fmt.Println("js:", i)
			}
			if !zbDone && p.src == "zb" && p.p {
				fmt.Println("zb:", i)
			}
			if !zbDone && p.src == "rr" && p.p {
				fmt.Println("rr:", i)
			}
			if !zbDone && p.src == "bs" && p.p {
				fmt.Println("bs:", i)
			}
		}

		var s string
		var c map[string]int

		if !jsDone {
			s = state(jsClique)
			c = jsCache
			if j, ok := c[s]; ok {
				fmt.Println("js done:", j, i)
				jsDone = true
			}
			c[s] = i
		}

		if !zbDone {
			s = state(zbClique)
			c = zbCache
			if j, ok := c[s]; ok {
				fmt.Println("zb done:", j, i)
				zbDone = true
			}
			c[s] = i
		}

		if !rrDone {
			s = state(rrClique)
			c = rrCache
			if j, ok := c[s]; ok {
				fmt.Println("rr done:", j, i)
				rrDone = true
			}
			c[s] = i
		}

		if !bsDone {
			s = state(bsClique)
			c = bsCache
			if j, ok := c[s]; ok {
				fmt.Println("bs done:", j, i)
				bsDone = true
			}
			c[s] = i
		}

		// fmt.Println(state(moduleNames, modulesMap))
		// fmt.Println()
	}

	part1Total := highs * lows

	fmt.Println(highs, lows)
	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)
}

func getClique(moduleMap map[string]Module, start, end string) ([]Module, map[string]int) {
	newModuleMap := make(map[string]Module, 0)
	stack := []Module{moduleMap[start]}
	for len(stack) > 0 {
		l := len(stack) - 1
		m := stack[l]
		stack = stack[:l]
		name := m.getName()
		if _, ok := newModuleMap[name]; ok {
			continue
		}
		newModuleMap[name] = m
		if name == end {
			continue
		}
		for _, n := range m.getDestinations() {
			stack = append(stack, moduleMap[n])
		}
	}

	modules := make([]Module, 0, len(moduleMap))
	for _, m := range newModuleMap {
		modules = append(modules, m)
	}

	cache := map[string]int{}

	return modules, cache
}

func state(modules []Module) (ret string) {
	for _, m := range modules {
		ret += fmt.Sprint(m.getStates()) + "\t"
	}
	return
}

type Pulse struct {
	p    bool
	src  string
	dest string
}

type Module interface {
	pulse(bool, string) []Pulse
	getDestinations() []string
	getName() string
	getStates() []bool
}

func NewModule(line string) (name string, mod Module, conjunction bool) {
	splitLine := strings.Split(line, " -> ")
	destNames := strings.Split(splitLine[1], ", ")
	if splitLine[0] == "broadcaster" {
		name = "broadcaster"
		mod = &Broadcaster{name, destNames}
		return
	}
	name = splitLine[0][1:]
	if splitLine[0][0] == '%' {
		mod = &FlipFlop{name, destNames, false}
		return
	}
	mod = &Conjunction{name, destNames, make(map[string]bool, 0), make([]string, 0)}
	conjunction = true
	return
}

type Broadcaster struct {
	name         string
	destinations []string
}

func (m *Broadcaster) pulse(p bool, _ string) []Pulse {
	newPulses := make([]Pulse, len(m.destinations))
	for i, d := range m.destinations {
		newPulses[i] = Pulse{p, m.name, d}
	}
	return newPulses
}

func (m *Broadcaster) getDestinations() []string {
	return m.destinations
}

func (m *Broadcaster) getStates() []bool {
	return []bool{}
}

func (m *Broadcaster) getName() string {
	return m.name
}

type FlipFlop struct {
	name         string
	destinations []string
	state        bool
}

func (m *FlipFlop) pulse(p bool, _ string) []Pulse {
	if p {
		return []Pulse{}
	}
	m.state = !m.state
	newPulses := make([]Pulse, len(m.destinations))
	for i, d := range m.destinations {
		newPulses[i] = Pulse{m.state, m.name, d}
	}
	return newPulses
}

func (m *FlipFlop) getDestinations() []string {
	return m.destinations
}

func (m *FlipFlop) getStates() []bool {
	return []bool{m.state}
}

func (m *FlipFlop) getName() string {
	return m.name
}

type Conjunction struct {
	name         string
	destinations []string
	inputStates  map[string]bool
	inputNames   []string
}

func (m *Conjunction) getDestinations() []string {
	return m.destinations
}

func (m *Conjunction) getName() string {
	return m.name
}

func (m *Conjunction) pulse(p bool, inName string) []Pulse {
	if _, ok := m.inputStates[inName]; !ok {
		fmt.Println(p)
		os.Exit(1)
	}
	m.inputStates[inName] = p
	newPulse := false
	for _, inState := range m.inputStates {
		if !inState {
			newPulse = true
			break
		}
	}

	newPulses := make([]Pulse, len(m.destinations))
	for i, d := range m.destinations {
		newPulses[i] = Pulse{newPulse, m.name, d}
	}
	return newPulses
}

func (m *Conjunction) getStates() []bool {
	states := make([]bool, len(m.inputStates))
	for i, n := range m.inputNames {
		states[i] = m.inputStates[n]
	}
	return states
}

type Node struct {
	p Pulse
	n *Node
}

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue(p Pulse) Queue {
	n := &Node{p, nil}
	return Queue{n, n}
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Push(p Pulse) {
	q.tail.n = &Node{p, nil}
	q.tail = q.tail.n
	if q.Empty() {
		q.head = q.tail
	}

}

func (q *Queue) Pop() Pulse {
	p := q.head.p
	q.head = q.head.n
	return p
}
