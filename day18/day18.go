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
	file, _ := os.Open("../inputs/day18.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Total := 0
	var part2Total int64
	var x1, y1 int
	var a1, b1 int64
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])
		x2 := x1
		y2 := y1
		switch direction {
		case "R":
			x2 += distance
		case "L":
			x2 -= distance
		case "U":
			y2 -= distance
		case "D":
			y2 += distance
		}
		part1Total += x1*y2 - y1*x2 + distance
		x1 = x2
		y1 = y2

		l := len(splitLine[2])
		direction2 := splitLine[2][l-2]
		distance2, _ := strconv.ParseInt(splitLine[2][2:l-2], 16, 0)
		a2 := a1
		b2 := b1
		switch direction2 {
		case '0':
			a2 += distance2
		case '2':
			a2 -= distance2
		case '3':
			b2 -= distance2
		case '1':
			b2 += distance2
		}
		part2Total += a1*b2 - b1*a2 + distance2
		a1 = a2
		b1 = b2
	}
	part1Total /= 2
	part2Total /= 2
	fmt.Println("Part 1:", part1Total+1)
	fmt.Println("Part 1:", part2Total+1)
}
