package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day14.txt")
	defer file.Close()

	part1Total := 0
	part2Total := 0

	scanner := bufio.NewScanner(file)

	columns := make([]*Column, 0)
	scanner.Scan()
	numRocks := 0
	indexCount := 0
	line := scanner.Text()
	for i := 0; i < len(line); i++ {
		c := line[i]
		nextI := 0
		if c == '#' {
			nextI = 1
		}
		if c == 'O' {
			nextI = 1
			numRocks++
		}
		columns = append(columns, &Column{[]byte{c}, nextI})
	}

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			c := line[i]
			column := columns[i]
			switch c {
			case '#':
				column.col = append(column.col, c)
				column.nextI = len(column.col)
			case 'O':
				column.col = append(column.col, '.')
				column.col[column.nextI] = c
				indexCount += column.nextI
				numRocks++
				column.nextI++
			case '.':
				column.col = append(column.col, '.')
			}
		}
	}

	numRows := len(columns[0].col)
	part1Total = numRows*numRocks - indexCount

	fmt.Println("part 1:", part1Total)
	fmt.Println("part 2:", part2Total)
}

type Column struct {
	col   []byte
	nextI int
}
