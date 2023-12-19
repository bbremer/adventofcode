package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day19.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Total := 0
	fLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		fLines = append(fLines, line)
	}

	f := makeF(fLines)

	for scanner.Scan() {
		part := NewPart(scanner.Text())
		if f(part) {
			part1Total += part.total()
		}
	}

	part2Total := part2(fLines)

	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)
}

type Part struct {
	x int
	m int
	a int
	s int
}

func (p Part) get(c byte) int {
	switch c {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}
	fmt.Println("problem")
	os.Exit(1)
	return 0
}

func (p Part) total() int {
	return p.x + p.m + p.a + p.s
}

func NewPart(line string) Part {
	line = line[1 : len(line)-1]
	splitLine := strings.Split(line, ",")
	if splitLine[0][:2] != "x=" || splitLine[1][:2] != "m=" || splitLine[2][:2] != "a=" || splitLine[3][:2] != "s=" {
		fmt.Println("invalid line:", line)
		os.Exit(1)
	}
	x, _ := strconv.Atoi(splitLine[0][2:])
	m, _ := strconv.Atoi(splitLine[1][2:])
	a, _ := strconv.Atoi(splitLine[2][2:])
	s, _ := strconv.Atoi(splitLine[3][2:])
	return Part{x, m, a, s}
}

func makeF(lines []string) func(Part) bool {
	fs := make(map[string]func(Part) string, 0)
	for _, line := range lines {
		index := strings.IndexByte(line, '{')
		name := line[:index]
		f := opsToF(line[index+1 : len(line)-1])
		fs[name] = f
	}
	return func(p Part) bool {
		s := "in"
		for {
			if s == "A" {
				return true
			}
			if s == "R" {
				return false
			}
			s = fs[s](p)
		}
	}
}

func opsToF(ops string) func(Part) string {
	splitOps := strings.Split(ops, ",")
	fs := make([]func(Part) (string, bool), 0)
	for _, op := range splitOps {
		fs = append(fs, opToF(op))
	}
	return func(p Part) string {
		for _, f := range fs {
			if s, ok := f(p); ok {
				return s
			}
		}
		fmt.Println("problem")
		os.Exit(1)
		return ""
	}
}

func opToF(op string) func(Part) (string, bool) {
	colonIndex := strings.IndexByte(op, ':')
	if colonIndex == -1 {
		return func(Part) (string, bool) { return op, true }
	}
	attr := op[0]
	compF := parseComp(op[1])
	value, _ := strconv.Atoi(op[2:colonIndex])
	next := op[colonIndex+1:]
	return func(p Part) (string, bool) {
		return next, compF(p.get(attr), value)
	}
}

func parseComp(c byte) func(int, int) bool {
	if c == '>' {
		return func(x, y int) bool { return x > y }
	}
	if c == '<' {
		return func(x, y int) bool { return x < y }
	}
	fmt.Println("problem in parseComp")
	os.Exit(1)
	return func(int, int) bool { return false }
}

func part2(lines []string) int {
	fLines := make(map[string]string, 0)
	for _, line := range lines {
		index := strings.IndexByte(line, '{')
		name := line[:index]
		fLines[name] = line[index+1 : len(line)-1]
	}
	xRange := [2]int{1, 4001}
	mRange := [2]int{1, 4001}
	aRange := [2]int{1, 4001}
	sRange := [2]int{1, 4001}
	return recursiveParse("in", fLines, xRange, mRange, aRange, sRange)
}

func recursiveParse(fName string, fLines map[string]string, xRange, mRange, aRange, sRange [2]int) (sum int) {
	if fName == "A" {
		return rangeLen(xRange) * rangeLen(mRange) * rangeLen(aRange) * rangeLen(sRange)
	}
	if fName == "R" {
		return 0
	}

	splitOps := strings.Split(fLines[fName], ",")
	for _, op := range splitOps {
		colonIndex := strings.IndexByte(op, ':')
		if colonIndex == -1 {
			sum += recursiveParse(op, fLines, xRange, mRange, aRange, sRange)
			return
		}
		newFName := op[colonIndex+1:]
		attr := op[0]
		comp := op[1]
		value, _ := strconv.Atoi(op[2:colonIndex])
		var r [2]int
		switch attr {
		case 'x':
			r, xRange = splitRange(comp, value, xRange)
			sum += recursiveParse(newFName, fLines, r, mRange, aRange, sRange)
		case 'm':
			r, mRange = splitRange(comp, value, mRange)
			sum += recursiveParse(newFName, fLines, xRange, r, aRange, sRange)
		case 'a':
			r, aRange = splitRange(comp, value, aRange)
			sum += recursiveParse(newFName, fLines, xRange, mRange, r, sRange)
		case 's':
			r, sRange = splitRange(comp, value, sRange)
			sum += recursiveParse(newFName, fLines, xRange, mRange, aRange, r)
		default:
			fmt.Println("problem in recursiveParse")
			os.Exit(1)
		}
	}
	return
}

func rangeLen(r [2]int) int {
	return r[1] - r[0]
}

func splitRange(c byte, v int, r [2]int) ([2]int, [2]int) {
	if v < r[0] || r[1] <= v || c != '>' && c != '<' {
		fmt.Println("problem in splitRange")
		os.Exit(1)
	}
	if c == '>' {
		return [2]int{v + 1, r[1]}, [2]int{r[0], v + 1}
	}
	return [2]int{r[0], v}, [2]int{v, r[1]}
}

func makeRange() []int {
	x := make([]int, 4000)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}
