package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	// file, _ := os.Open("../inputs/test.txt")
	file, _ := os.Open("../inputs/day14.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	columns, numRocks, indexCount := scannerToColumns(scanner)

	numRows := len(columns[0].col)
	part1Total := numRows*numRocks - indexCount
	fmt.Println("part 1:", part1Total)

	target := 1000000000

	firstCyclePoint, cycleLength, cache := part2Cycles(columns)
	endRocksIndex := (target-firstCyclePoint)%cycleLength + firstCyclePoint
	var endRocks string
	for k, v := range cache {
		if v == endRocksIndex {
			endRocks = k
		}
	}
	part2Total := numRows*numRocks - countIndicies(endRocks)

	fmt.Println("part 2:", part2Total)
}

func scannerToColumns(scanner *bufio.Scanner) ([]*Column, int, int) {
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
	return columns, numRocks, indexCount
}

type Column struct {
	col   []byte
	nextI int
}

func printColumns(columns []*Column) {
	fmt.Println(columnsToString(columns))
}

func columnsToString(columns []*Column) string {
	var newString bytes.Buffer
	for i := 0; i < len(columns[0].col); i++ {
		for j := 0; j < len(columns); j++ {
			newString.WriteByte(columns[j].col[i])
		}
		newString.WriteByte('\n')
	}
	return newString.String()
}

func part2Cycles(columns []*Column) (int, int, map[string]int) {
	cache := make(map[string]int, 10000)

	// west
	columns = rotate(columns)
	columns = tilt(columns)

	// south
	columns = rotate(columns)
	columns = tilt(columns)

	// east
	columns = rotate(columns)
	columns = tilt(columns)

	numRotations := 1000000
	for i := 2; i <= numRotations; i++ {
		columns = cycle(columns)

		tempColumns := rotate(columns)
		str := columnsToString(tempColumns)
		if j, ok := cache[str]; ok {
			return j, i - j, cache
		} else {
			cache[str] = i
		}
	}

	fmt.Println("Problem")
	os.Exit(1)
	return -1, -1, cache
}

func cycle(columns []*Column) []*Column {
	//north
	columns = rotate(columns)
	columns = tilt(columns)

	// west
	columns = rotate(columns)
	columns = tilt(columns)

	// south
	columns = rotate(columns)
	columns = tilt(columns)

	// east
	columns = rotate(columns)
	columns = tilt(columns)

	return columns
}

func rotate(columns []*Column) []*Column {
	columns = transpose(columns)
	slices.Reverse(columns)
	return columns
}

func transpose(columns []*Column) []*Column {
	transposed := make([]*Column, 0, len(columns[0].col))
	for i := range columns[0].col {
		var newCol []byte
		for _, c := range columns {
			newCol = append(newCol, c.col[i])
		}
		transposed = append(transposed, &Column{newCol, -1})
	}
	return transposed
}

func tilt(columns []*Column) []*Column {
	r := strings.NewReader(columnsToString(columns))
	scanner := bufio.NewScanner(r)
	s, _, _ := scannerToColumns(scanner)
	return s
}

func countIndicies(rockString string) (count int) {
	for i, line := range strings.Split(rockString, "\n") {
		for j := 0; j < len(line); j++ {
			if line[j] == 'O' {
				count += i
			}
		}
		/*
			if len(line) > 0 {
				fmt.Println(line, count)
			}
		*/
	}
	return
}
