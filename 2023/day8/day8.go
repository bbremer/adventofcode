package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("../inputs/day8.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	nodes := buildNodes(scanner)
	part1Total := part1(instructions, nodes, "AAA", allZ)
	fmt.Println("part 1:", part1Total)

	startNodes := makePart2Starts(nodes)
	var total uint64 = 1
	for _, x := range startNodes {
		y := part1(instructions, nodes, x, oneZ)
		total *= uint64(y / len(instructions))
	}
	fmt.Println("part 2:", total*uint64(len(instructions)))
}

func buildNodes(scanner *bufio.Scanner) (m map[string][2]string) {
	m = make(map[string][2]string, 0)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " = ")
		name := splitLine[0]
		children := strings.Split(splitLine[1][1:len(splitLine[1])-1], ", ")
		m[name] = [2]string{children[0], children[1]}
	}
	return
}

func part1(instructions string, nodes map[string][2]string, node string, endFunc func(string) bool) (count int) {
	for {
		for _, c := range instructions {
			count++
			i := 0
			if c == 'R' {
				i = 1
			}
			node = nodes[node][i]
			if endFunc(node) {
				return
			}
		}
	}
}

func oneZ(node string) bool {
	return node[2] == 'Z'
}

func allZ(node string) bool {
	return node == "ZZZ"
}

func makePart2Starts(nodes map[string][2]string) []string {
	newStarts := make([]string, 0)
	for k := range nodes {
		if k[2] == 'A' {
			newStarts = append(newStarts, k)
		}
	}
	return newStarts
}
